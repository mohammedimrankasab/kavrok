package discovery

import (
	"context"
	"fmt"

	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
)

// DiscoverNamespaces retrieves Kubernetes namespaces.
func DiscoverNamespaces(
	ctx context.Context,
	client kubernetes.Client,
) ([]NamespaceSummary, error) {
	namespaces, err := client.ListNamespaces(ctx)
	if err != nil {
		return nil, fmt.Errorf(
			"list Kubernetes namespaces: %w",
			err,
		)
	}

	result := make([]NamespaceSummary, 0, len(namespaces.Items))

	for _, namespace := range namespaces.Items {
		result = append(result, NamespaceSummary{
			Name: namespace.Name,
		})
	}

	return result, nil
}
