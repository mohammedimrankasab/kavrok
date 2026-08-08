// Package kubernetes provides Kubernetes cluster connectivity
// and client abstractions for Kavrok.
package kubernetes

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Client represents the Kubernetes operations required by Kavrok.
type Client interface {
	ServerVersion() (*version.Info, error)
	ListNodes(ctx context.Context) (*corev1.NodeList, error)

	ListNamespaces(context.Context) (*corev1.NamespaceList, error)
	ListPods(context.Context) (*corev1.PodList, error)
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

// ListNodes returns all nodes in the Kubernetes cluster.
func (c *client) ListNodes(ctx context.Context) (*corev1.NodeList, error) {
	return c.clientset.CoreV1().
		Nodes().
		List(ctx, metav1.ListOptions{})
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

// ListNamespaces returns all namespaces in the Kubernetes cluster.
func (c *client) ListNamespaces(
	ctx context.Context,
) (*corev1.NamespaceList, error) {
	return c.clientset.CoreV1().
		Namespaces().
		List(ctx, metav1.ListOptions{})
}

// ListPods returns all pods across all namespaces.
func (c *client) ListPods(
	ctx context.Context,
) (*corev1.PodList, error) {
	return c.clientset.CoreV1().
		Pods("").
		List(ctx, metav1.ListOptions{})
}
