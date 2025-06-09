package fsm

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
)

func sFnInitialize(ctx context.Context, m *fsm, s *systemState) (stateFn, *ctrl.Result, error) {
	m.log.Info("KymaPolicy FSM initialisation state")
	return stop()
}
