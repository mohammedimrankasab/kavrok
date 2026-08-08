// Package kubernetes provides Kubernetes cluster connectivity
// and client abstractions for Kavrok.
package kubernetes

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Client represents the Kubernetes operations required by Kavrok.
type Client interface {
	ServerVersion() (*version.Info, error)
}

// client implements Client using Kubernetes client-go.
type client struct {
	clientset kubernetes.Interface
}

// New creates a Kubernetes client using the current kubeconfig.
func New() (Client, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("create kubernetes client: %w", err)
	}

	return &client{
		clientset: clientset,
	}, nil
}

// ServerVersion returns the Kubernetes API server version.
func (c *client) ServerVersion() (*version.Info, error) {
	return c.clientset.Discovery().
		ServerVersion()
}

func loadConfig() (*rest.Config, error) {
	if configPath := os.Getenv("KUBECONFIG"); configPath != "" {
		return clientcmd.BuildConfigFromFlags("", configPath)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("determine home directory: %w", err)
	}

	configPath := filepath.Join(home, ".kube", "config")

	return clientcmd.BuildConfigFromFlags("", configPath)
}
