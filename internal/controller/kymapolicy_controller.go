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
	"github.com/kyma-project/policy-manager/internal/controller/fsm"
	"time"

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

// reconciler specific configuration
type PolicyModuleCfg struct {
	Finalizer string
	dryRun    bool
}

// KymaPolicyReconciler reconciles a KymaPolicy object
type KymaPolicyReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	PolicyModuleCfg
}

// +kubebuilder:rbac:groups=operator.kyma-project.io,resources=kymapolicies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=operator.kyma-project.io,resources=kymapolicies/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=operator.kyma-project.io,resources=kymapolicies/finalizers,verbs=update

// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.21.0/pkg/reconcile
func (r *KymaPolicyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx)

	var instance operatorv1alpha1.KymaPolicyConfig
	if err := r.Get(ctx, req.NamespacedName, &instance); err != nil {
		return ctrl.Result{
			RequeueAfter: time.Duration(5) * time.Second,
		}, client.IgnoreNotFound(err)
	}

	// TODO(user): your logic here

	stateFSM := fsm.NewFsm(log, r.Client)
	return stateFSM.Run(ctx, instance)

	//return ctrl.Result{}, nil
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
		For(&operatorv1alpha1.KymaPolicyConfig{}).
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
