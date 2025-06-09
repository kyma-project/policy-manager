package fsm

import (
	operatorv1alpha1 "github.com/kyma-project/policy-manager/api/v1alpha1"
	log_level "github.com/kyma-project/policy-manager/internal/log"
	ctrl "sigs.k8s.io/controller-runtime"
)

import (
	"context"
	"fmt"
	"reflect"
	"runtime"

	"github.com/go-logr/logr"
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
	log logr.Logger
	client.Client
}

func NewFsm(log logr.Logger, client client.Client) Fsm {
	return &fsm{
		fn:     sFnTakeSnapshot,
		log:    log,
		Client: client,
	}
}

/*
type Watch = func(src source.Source, eventhandler handler.EventHandler, predicates ...predicate.Predicate) error

type K8s struct {
	client.Client
	record.EventRecorder
	ShootClient client.Client
}*/

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
			m.log.V(log_level.TRACE).WithValues("result", result, "err", err, "mFnIsNill", m.fn == nil).Info(fmt.Sprintf("switching state from %s to %s", stateFnName, newStateFnName))
			if m.fn == nil || err != nil {
				break loop
			}
		}
	}

	m.log.V(log_level.DEBUG).
		WithValues("result", result).
		Info("Reconciliation done")

	if result != nil {
		return *result, err
	}

	return ctrl.Result{}, err
}
