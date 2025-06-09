/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	operatorv1alpha1 "github.com/kyma-project/policy-manager/api/v1alpha1"
	policyv1 "github.com/kyverno/kyverno/api/kyverno/v1"
)

// KymaPolicyReconciler reconciles a KymaPolicy object
type KymaPolicyReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=operator.kyma-project.io,resources=kymapolicies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=operator.kyma-project.io,resources=kymapolicies/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=operator.kyma-project.io,resources=kymapolicies/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the KymaPolicy object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.21.0/pkg/reconcile
func (r *KymaPolicyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = logf.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

var (
	specChanged = predicate.GenerationChangedPredicate{}

	//nolint:gochecknoglobals
	labelsManagedByPolicyManager = map[string]string{
		"reconciler.kyma-project.io/managed-by": "policy-manager",
	}
)

// SetupWithManager sets up the controller with the Manager.
func (r *KymaPolicyReconciler) SetupWithManager(mgr ctrl.Manager) error {
	managedByPolicyManagerPredicate, err := predicate.LabelSelectorPredicate(*v1.SetAsLabelSelector(labels.Set(labelsManagedByPolicyManager)))
	if err != nil {
		return err
	}

	var handler kyvernoResourceEventHandler
	return ctrl.NewControllerManagedBy(mgr).
		For(&operatorv1alpha1.KymaPolicy{}).
		Watches(
			&policyv1.ClusterPolicy{},
			handler,
			builder.WithPredicates(
				predicate.And(
					specChanged,
					managedByPolicyManagerPredicate,
				),
			),
		).
		Named("kymapolicy").
		Complete(r)
}
