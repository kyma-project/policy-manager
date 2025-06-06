package controller

import (
	"context"

	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var (
	// document - generic + update works by default; create + delete is muted;
	KyvernoResourceEventHandler = handler.Funcs{
		CreateFunc: func(
			_ context.Context,
			_ event.TypedCreateEvent[client.Object],
			_ workqueue.TypedRateLimitingInterface[reconcile.Request],
		) {
			// mute event
		},

		DeleteFunc: func(
			_ context.Context,
			_ event.TypedDeleteEvent[client.Object],
			_ workqueue.TypedRateLimitingInterface[reconcile.Request],
		) {
			// mute event
		},
	}
)
