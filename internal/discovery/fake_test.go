package discovery

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/version"

	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
)

type fakeKubernetesClient struct {
	serverVersion *version.Info
	clusterName   string
	nodes         *corev1.NodeList
	namespaces    *corev1.NamespaceList
	pods          *corev1.PodList

	versionErr    error
	nodesErr      error
	namespacesErr error
	podsErr       error
}

func (f *fakeKubernetesClient) ServerVersion() (*version.Info, error) {
	if f.versionErr != nil {
		return nil, f.versionErr
	}

	return f.serverVersion, nil
}

func (f *fakeKubernetesClient) ListNodes(
	_ context.Context,
) (*corev1.NodeList, error) {
	if f.nodesErr != nil {
		return nil, f.nodesErr
	}

	return f.nodes, nil
}

func (f *fakeKubernetesClient) ListNamespaces(
	_ context.Context,
) (*corev1.NamespaceList, error) {
	return f.namespaces, f.namespacesErr
}

func (f *fakeKubernetesClient) ListPods(
	_ context.Context,
) (*corev1.PodList, error) {
	return f.pods, f.podsErr
}

func (f *fakeKubernetesClient) ClusterName() string {
	return f.clusterName
}

var _ kubernetes.Client = (*fakeKubernetesClient)(nil)
