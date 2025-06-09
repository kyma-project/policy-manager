package controller

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllertest"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var instance kyvernoResourceEventHandler

func Test_Create(t *testing.T) {
	q := &controllertest.Queue{TypedInterface: workqueue.NewTyped[reconcile.Request]()}
	var ctx = context.Background()

	evt := event.CreateEvent{
		Object: clusterPolicy(t, podPolicyWithNamespacedNameOpt("test", "me")),
	}
	instance.Create(ctx, evt, q)
	assert.Equal(t, 0, q.Len())
}

func Test_Delete(t *testing.T) {
	q := &controllertest.Queue{TypedInterface: workqueue.NewTyped[reconcile.Request]()}
	var ctx = context.Background()

	evt := event.DeleteEvent{
		Object: clusterPolicy(t, podPolicyWithNamespacedNameOpt("test", "me")),
	}
	instance.Delete(ctx, evt, q)
	assert.Equal(t, 0, q.Len())
}

func Test_Update(t *testing.T) {
	q := &controllertest.Queue{TypedInterface: workqueue.NewTyped[reconcile.Request]()}
	var ctx = context.Background()

	evt := event.UpdateEvent{
		ObjectNew: clusterPolicy(t, podPolicyWithNamespacedNameOpt("test", "me")),
		ObjectOld: clusterPolicy(t, podPolicyWithNamespacedNameOpt("test", "me")),
	}
	instance.Update(ctx, evt, q)
	assert.Equal(t, 1, q.Len())
}

func Test_Generic(t *testing.T) {
	q := &controllertest.Queue{TypedInterface: workqueue.NewTyped[reconcile.Request]()}
	var ctx = context.Background()

	evt := event.GenericEvent{
		Object: clusterPolicy(t, podPolicyWithNamespacedNameOpt("test", "me")),
	}
	instance.Generic(ctx, evt, q)
	assert.Equal(t, 1, q.Len())
}
