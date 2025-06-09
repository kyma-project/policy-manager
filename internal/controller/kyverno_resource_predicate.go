package controller

import (
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

func managedByPolicyManagerPredicate() predicate.Predicate {
	p, err := predicate.LabelSelectorPredicate(*v1.SetAsLabelSelector(labels.Set(labelsManagedByPolicyManager)))
	if err != nil {
		panic(fmt.Errorf("failed to create label selector predicate for policy manager: %w", err))
	}
	return predicate.And(
		predicate.GenerationChangedPredicate{},
		p,
	)
}
