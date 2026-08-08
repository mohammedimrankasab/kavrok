package discovery

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/version"

	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
)

type fakeKubernetesClient struct {
	serverVersion *version.Info
	nodes         *corev1.NodeList

	versionErr error
	nodesErr   error
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

var _ kubernetes.Client = (*fakeKubernetesClient)(nil)
