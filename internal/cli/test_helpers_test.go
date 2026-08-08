package cli

import (
	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
	"k8s.io/apimachinery/pkg/version"
)

type fakeKubernetesClient struct {
	serverVersion *version.Info
	err           error
}

func (f *fakeKubernetesClient) ServerVersion() (*version.Info, error) {
	return f.serverVersion, f.err
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
