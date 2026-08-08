// Package discovery provides Kubernetes environment discovery for Kavrok.
package discovery

// ClusterInfo contains information about the Kubernetes cluster.
type ClusterInfo struct {
	Name     string
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

// NamespaceSummary contains information about a Kubernetes namespace.
type NamespaceSummary struct {
	Name string
}

// PodSummary contains information about a Kubernetes pod.
type PodSummary struct {
	Name           string
	Namespace      string
	Phase          string
	Ready          bool
	PendingReason  string
	PendingMessage string
}

// WorkloadSummary contains high-level workload information.
type WorkloadSummary struct {
	PodCount     int
	RunningPods  int
	PendingPods  int
	FailedPods   int
	ReadyPods    int
	NotReadyPods int
	Pods         []PodSummary
}

// ClusterSnapshot contains the discovered state of a Kubernetes cluster.
type ClusterSnapshot struct {
	Cluster    ClusterInfo
	Nodes      NodeSummaryList
	Namespaces []NamespaceSummary
	Workloads  WorkloadSummary
	Health     ClusterSummary
}
