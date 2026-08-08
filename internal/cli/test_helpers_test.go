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
}

func (f *fakeKubernetesClient) ServerVersion() (*version.Info, error) {
	return f.serverVersion, f.err
}
func (f *fakeKubernetesClient) ListNodes(
	_ context.Context,
) (*corev1.NodeList, error) {
	return nil, nil
}

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
