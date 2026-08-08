// Package discovery provides Kubernetes environment discovery for Kavrok.
package discovery

// ClusterInfo contains information about the Kubernetes cluster.
type ClusterInfo struct {
	Version  string
	Platform string
	Commit   string
}

// NodeSummary contains the information Kavrok needs about a node.
type NodeSummary struct {
	Name              string
	Ready             bool
	Roles             []string
	KubernetesVersion string
	OS                string
	Architecture      string
}

// NodeSummaryList contains discovered Kubernetes nodes.
type NodeSummaryList struct {
	Nodes []NodeSummary
}

// ClusterSummary contains the high-level health state of a Kubernetes cluster.
type ClusterSummary struct {
	KubernetesVersion string
	NodeCount         int
	ReadyNodes        int
	NotReadyNodes     int
	Healthy           bool
}
