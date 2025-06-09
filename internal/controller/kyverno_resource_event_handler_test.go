package controller

import (
	"context"
	"testing"

	policyv1 "github.com/kyverno/kyverno/api/kyverno/v1"
	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllertest"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var podPolicy policyv1.ClusterPolicy
var instance kyvernoResourceEventHandler

func Test_Create(t *testing.T) {
	q := &controllertest.Queue{TypedInterface: workqueue.NewTyped[reconcile.Request]()}
	var ctx = context.Background()

	evt := event.CreateEvent{
		Object: &podPolicy,
	}
	instance.Create(ctx, evt, q)
	assert.Equal(t, 0, q.Len())
}

func Test_Delete(t *testing.T) {
	q := &controllertest.Queue{TypedInterface: workqueue.NewTyped[reconcile.Request]()}
	var ctx = context.Background()

	evt := event.DeleteEvent{
		Object: &podPolicy,
	}
	instance.Delete(ctx, evt, q)
	assert.Equal(t, 0, q.Len())
}

func Test_Update(t *testing.T) {
	q := &controllertest.Queue{TypedInterface: workqueue.NewTyped[reconcile.Request]()}
	var ctx = context.Background()

	evt := event.UpdateEvent{
		ObjectOld: &podPolicy,
		ObjectNew: &podPolicy,
	}
	instance.Update(ctx, evt, q)
	assert.Equal(t, 1, q.Len())
}

func Test_Generic(t *testing.T) {
	q := &controllertest.Queue{TypedInterface: workqueue.NewTyped[reconcile.Request]()}
	var ctx = context.Background()

	evt := event.GenericEvent{
		Object: &podPolicy,
	}
	instance.Generic(ctx, evt, q)
	assert.Equal(t, 1, q.Len())
}
