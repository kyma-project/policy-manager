package fsm

import (
	"context"
	"reflect"
	ctrl "sigs.k8s.io/controller-runtime"
)

func sFnUpdateStatus(result *ctrl.Result, err error) stateFn {
	return func(ctx context.Context, m *fsm, s *systemState) (stateFn, *ctrl.Result, error) {

		// make sure there is a change in status
		if reflect.DeepEqual(s.instance.Status, s.status) {
			return nil, result, err
		}

		updateErr := m.Status().Update(ctx, &s.instance)

		if updateErr != nil {
			m.log.Error(updateErr, "unable to update instance status!")
			if err == nil {
				err = updateErr
			}
			return nil, nil, err
		}

		//next := sFnEmmitEventfunc(nil, result, err)
		//return next, nil, nil
		return stop()
	}
}
