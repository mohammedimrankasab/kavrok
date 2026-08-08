package discovery

import (
	"fmt"

	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
)

// Discover retrieves information about the Kubernetes cluster.
func Discover(client kubernetes.Client) (ClusterInfo, error) {
	info, err := client.ServerVersion()
	if err != nil {
		return ClusterInfo{}, fmt.Errorf(
			"get Kubernetes server version: %w",
			err,
		)
	}
	if info == nil {
		return ClusterInfo{}, fmt.Errorf(
			"get Kubernetes server version: empty response",
		)
	}

	return ClusterInfo{
		Version:  info.GitVersion,
		Platform: info.Platform,
		Commit:   info.GitCommit,
	}, nil
}
