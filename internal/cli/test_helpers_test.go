package cli

import (
	"context"

	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/version"
)

type fakeKubernetesClient struct {
	serverVersion *version.Info
	err           error
	pods          *corev1.PodList
	namespaceList *corev1.NamespaceList
	nodes         *corev1.NodeList
}

func (f *fakeKubernetesClient) ServerVersion() (*version.Info, error) {
	return f.serverVersion, f.err
}
func (f *fakeKubernetesClient) ListNodes(
	_ context.Context,
) (*corev1.NodeList, error) {
	if f.err != nil {
		return nil, f.err
	}

	return f.nodes, nil
}

func (f *fakeKubernetesClient) ListNamespaces(
	_ context.Context,
) (*corev1.NamespaceList, error) {
	if f.err != nil {
		return nil, f.err
	}
	return f.namespaceList, nil
}

func (f *fakeKubernetesClient) ListPods(
	_ context.Context,
) (*corev1.PodList, error) {
	if f.err != nil {
		return nil, f.err
	}
	return f.pods, nil
}

var _ kubernetes.Client = (*fakeKubernetesClient)(nil)

func newTestDependencies() Dependencies {
	return Dependencies{
		KubernetesClientFactory: func() (kubernetes.Client, error) {
			return &fakeKubernetesClient{
				serverVersion: &version.Info{
					GitVersion: "v1.34.0",
				},
			}, nil
		},
	}
}
