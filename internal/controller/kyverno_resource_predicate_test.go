package controller

import (
	"testing"

	policyv1 "github.com/kyverno/kyverno/api/kyverno/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

func TestPredicate(t *testing.T) {
	instance := managedByPolicyManagerPredicate()

	type testCase struct {
		name     string
		evt      event.UpdateEvent
		expected bool
	}

	for _, tc := range []testCase{
		{
			name: "should handle event if generation changed and managed by policy manager",
			evt: event.UpdateEvent{
				ObjectOld: clusterPolicy(t,
					podPolicyWithNamespacedNameOpt("test", "me"),
					podPolicyMabagedByPolicyManagerOpt(),
					podPolicyWithGenerationOpt(1),
				),
				ObjectNew: clusterPolicy(t,
					podPolicyWithNamespacedNameOpt("test", "me"),
					podPolicyMabagedByPolicyManagerOpt(),
					podPolicyWithGenerationOpt(2),
				),
			},
			expected: true,
		},
		{
			name: "should not handle event if not managed by policy manager",
			evt: event.UpdateEvent{
				ObjectOld: clusterPolicy(t,
					podPolicyWithNamespacedNameOpt("test", "me"),
					podPolicyWithGenerationOpt(1),
				),
				ObjectNew: clusterPolicy(t,
					podPolicyWithNamespacedNameOpt("test", "me"),
					podPolicyWithGenerationOpt(2),
				),
			},
			expected: false,
		},
		{
			name: "should not handle event if resource version changed",
			evt: event.UpdateEvent{
				ObjectOld: clusterPolicy(t,
					podPolicyWithNamespacedNameOpt("test", "me"),
					podPolicyMabagedByPolicyManagerOpt(),
					podPolicyWithWithResourceVersionOpt("1"),
					podPolicyWithGenerationOpt(1),
				),
				ObjectNew: clusterPolicy(t,
					podPolicyWithNamespacedNameOpt("test", "me"),
					podPolicyMabagedByPolicyManagerOpt(),
					podPolicyWithWithResourceVersionOpt("2"),
					podPolicyWithGenerationOpt(1),
				),
			},
			expected: false,
		},
	} {
		actual := instance.Update(tc.evt)
		assert.Equal(t, tc.expected, actual)
	}
}

type podPolicyOpt func(*policyv1.ClusterPolicy) error

func clusterPolicy(t *testing.T, opts ...podPolicyOpt) *policyv1.ClusterPolicy {
	var clusterPolicy policyv1.ClusterPolicy
	for _, opt := range opts {
		if err := opt(&clusterPolicy); err != nil {
			require.NoError(t, err)
		}
	}
	return &clusterPolicy
}

func podPolicyWithNamespacedNameOpt(name, namespace string) podPolicyOpt {
	return func(cp *policyv1.ClusterPolicy) error {
		cp.Name = name
		cp.Namespace = namespace
		return nil
	}
}

func podPolicyMabagedByPolicyManagerOpt() podPolicyOpt {
	return func(cp *policyv1.ClusterPolicy) error {
		cp.Labels = labelsManagedByPolicyManager
		return nil
	}
}

func podPolicyWithGenerationOpt(generation int64) podPolicyOpt {
	return func(cp *policyv1.ClusterPolicy) error {
		cp.Generation = generation
		return nil
	}
}

func podPolicyWithWithResourceVersionOpt(resourceVersion string) podPolicyOpt {
	return func(cp *policyv1.ClusterPolicy) error {
		cp.ResourceVersion = resourceVersion
		return nil
	}
}
