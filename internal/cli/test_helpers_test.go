// Package cli provides tests for the Kavrok command-line interface.
package cli

import (
	"context"

	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
)

type fakeKubernetesClient struct {
	serverVersion *version.Info
	err           error
	pods          *corev1.PodList
	namespaceList *corev1.NamespaceList
	nodes         *corev1.NodeList
	clusterName   string
}

func (f *fakeKubernetesClient) ServerVersion() (*version.Info, error) {
	return f.serverVersion, f.err
}

func (f *fakeKubernetesClient) ListNamespaces(
	_ context.Context,
) (*corev1.NamespaceList, error) {
	if f.err != nil {
		return nil, f.err
	}

	if f.namespaceList == nil {
		return &corev1.NamespaceList{}, nil
	}
	return f.namespaceList, nil
}
func (f *fakeKubernetesClient) ClusterName() string {
	return f.clusterName
}
func (f *fakeKubernetesClient) ListPods(
	_ context.Context,
) (*corev1.PodList, error) {
	if f.err != nil {
		return nil, f.err
	}
	if f.pods == nil {
		return &corev1.PodList{}, nil
	}

	return f.pods, nil
}

func (f *fakeKubernetesClient) ListNodes(
	_ context.Context,
) (*corev1.NodeList, error) {
	if f.err != nil {
		return nil, f.err
	}

	if f.nodes == nil {
		return &corev1.NodeList{}, nil
	}

	return f.nodes, nil
}

var _ kubernetes.Client = (*fakeKubernetesClient)(nil)

func newHealthyFakeKubernetesClient() *fakeKubernetesClient {
	return &fakeKubernetesClient{
		serverVersion: &version.Info{
			GitVersion: "v1.34.0",
		},
		clusterName: "test-cluster",
		nodes: &corev1.NodeList{
			Items: []corev1.Node{
				{
					Status: corev1.NodeStatus{
						Conditions: []corev1.NodeCondition{
							{
								Type:   corev1.NodeReady,
								Status: corev1.ConditionTrue,
							},
						},
					},
				},
			},
		},
		namespaceList: &corev1.NamespaceList{
			Items: []corev1.Namespace{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "default",
					},
				},
			},
		},
		pods: &corev1.PodList{
			Items: []corev1.Pod{
				{
					Status: corev1.PodStatus{
						Phase: corev1.PodRunning,
					},
				},
			},
		},
	}
}

func newTestDependencies() Dependencies {
	return Dependencies{
		KubernetesClientFactory: func() (kubernetes.Client, error) {
			return newHealthyFakeKubernetesClient(), nil
		},
	}
}
