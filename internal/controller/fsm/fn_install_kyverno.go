package fsm

import (
	"context"
	ctrl "sigs.k8s.io/controller-runtime"
)

// to save the KymaPolicyConfig status at the beginning of the reconciliation
func sFnInstallKyverno(ctx context.Context, m *fsm, s *systemState) (stateFn, *ctrl.Result, error) {
	m.log.Info("sFnInstallKyverno function called")
	return requeue()
}
