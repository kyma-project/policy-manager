package fsm

import (
	operatorv1alpha1 "github.com/kyma-project/policy-manager/api/v1alpha1"
)

// the state of the system during the reconciliation process

type systemState struct {
	instance operatorv1alpha1.KymaPolicyConfig
	status   operatorv1alpha1.KymaPolicyConfigStatus
}

func (s *systemState) savePolicyStatus() {
	result := s.instance.Status.DeepCopy()
	if result == nil {
		result = &operatorv1alpha1.KymaPolicyConfigStatus{}
	}
	s.status = *result
}
