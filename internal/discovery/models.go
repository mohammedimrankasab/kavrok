// Package discovery provides Kubernetes environment discovery for Kavrok.
package discovery

// ClusterInfo contains information about the Kubernetes cluster.
type ClusterInfo struct {
	Version  string
	Platform string
	Commit   string
}
