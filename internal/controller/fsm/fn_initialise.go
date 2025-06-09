package fsm

import (
	"context"
	v1alpha1 "github.com/kyma-project/policy-manager/api/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	ctrl "sigs.k8s.io/controller-runtime"
)

func sFnInitialize(ctx context.Context, m *fsm, s *systemState) (stateFn, *ctrl.Result, error) {
	m.log.Info("KymaPolicy FSM initialisation state")

	instanceIsBeingDeleted := !s.instance.GetDeletionTimestamp().IsZero()
	instanceHasFinalizer := controllerutil.ContainsFinalizer(&s.instance, v1alpha1.Finalizer)

	if !instanceIsBeingDeleted && !instanceHasFinalizer {
		return addFinalizerAndRequeue(ctx, m, s)
	}

	if instanceIsBeingDeleted {
		if isKyvernoInstalled() {
			return switchState(sFnDeleteKyverno)
		}

		if instanceHasFinalizer {
			return removeFinalizerAndStop(ctx, m, s) // resource cleanup completed
		}
	}

	return stop()
}

func addFinalizerAndRequeue(ctx context.Context, m *fsm, s *systemState) (stateFn, *ctrl.Result, error) {
	controllerutil.AddFinalizer(&s.instance, v1alpha1.Finalizer)

	err := m.Update(ctx, &s.instance)
	if err != nil {
		return updateStatusAndStopWithError(err)
	}
	return requeue()
}

func removeFinalizerAndStop(ctx context.Context, m *fsm, s *systemState) (stateFn, *ctrl.Result, error) {
	controllerutil.RemoveFinalizer(&s.instance, v1alpha1.Finalizer)
	err := m.Update(ctx, &s.instance)
	if err != nil {
		return updateStatusAndStopWithError(err)
	}

	m.log.Info("Kyverno resources are deleted")

	return stop()
}

func isKyvernoInstalled() bool {
	return true // This function should check if Kyverno is installed in the cluster
}
