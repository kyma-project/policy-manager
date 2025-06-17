package fsm

import (
	"context"
	log "log/slog"
	"reflect"
	"runtime"

	operatorv1alpha1 "github.com/kyma-project/policy-manager/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type stateFn func(context.Context, *fsm, *systemState) (stateFn, *ctrl.Result, error)

func (f stateFn) String() string {
	return f.name()
}

func (f stateFn) name() string {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return name
}

type Fsm interface {
	Run(ctx context.Context, v operatorv1alpha1.KymaPolicyConfig) (ctrl.Result, error)
}
type fsm struct {
	fn  stateFn
	log log.Logger
	client.Client
}

func NewFsm(logger log.Logger, client client.Client) Fsm {
	return &fsm{
		fn:     sFnTakeSnapshot,
		log:    logger,
		Client: client,
	}
}

func (m *fsm) Run(ctx context.Context, v operatorv1alpha1.KymaPolicyConfig) (ctrl.Result, error) {
	state := systemState{instance: v}
	var err error
	var result *ctrl.Result
loop:
	for {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			break loop
		default:
			stateFnName := m.fn.name()
			m.fn, result, err = m.fn(ctx, m, &state)
			newStateFnName := m.fn.name()
			m.log.Debug("switching state",
				"result", result,
				"err", err,
				"from", stateFnName,
				"to", newStateFnName,
			)
			if m.fn == nil || err != nil {
				break loop
			}
		}
	}

	if result != nil {
		return *result, err
	}

	return ctrl.Result{}, err
}
