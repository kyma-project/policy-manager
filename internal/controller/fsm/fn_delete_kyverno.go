package fsm

import (
	"context"
	ctrl "sigs.k8s.io/controller-runtime"
)

// to save the KymaPolicy status at the beginning of the reconciliation
func sFnDeleteKyverno(ctx context.Context, m *fsm, s *systemState) (stateFn, *ctrl.Result, error) {
	m.log.Info("sFnDeleteKyverno function called")
	return requeue()
}
