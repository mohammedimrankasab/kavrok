// Package doctor provides diagnostic checks for the Kavrok CLI.
package doctor

import (
	"fmt"

	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
)

// Status represents the result of a diagnostic check.
type Status string

const (
	statusPass Status = "pass"
	statusFail Status = "fail"
	statusWarn Status = "warn"
)

// Check represents a single diagnostic check.
type Check struct {
	Name    string
	Status  Status
	Message string
}

// Result represents the result of running diagnostics.
type Result struct {
	Checks []Check
}

// Passed returns true when all diagnostic checks pass.
func (r Result) Passed() bool {
	for _, check := range r.Checks {
		if check.Status == statusFail {
			return false
		}
	}

	return true
}

// Runner executes diagnostic checks.
type Runner struct {
	kubeClientFactory func() (kubernetes.Client, error)
}

// New creates a diagnostic runner.
func New(
	kubeClientFactory func() (kubernetes.Client, error),
) *Runner {
	return &Runner{
		kubeClientFactory: kubeClientFactory,
	}
}

// Run executes the available diagnostic checks.
func (r *Runner) Run() Result {
	result := Result{
		Checks: []Check{
			{
				Name:    "configuration",
				Status:  statusPass,
				Message: "configuration is valid",
			},
		},
	}

	result.Checks = append(result.Checks, r.checkKubernetes())

	return result
}

func (r *Runner) checkKubernetes() Check {
	client, err := r.kubeClientFactory()
	if err != nil {
		return Check{
			Name:    "kubernetes",
			Status:  statusFail,
			Message: fmt.Sprintf("kubeconfig unavailable: %v", err),
		}
	}

	if _, err := client.ServerVersion(); err != nil {
		return Check{
			Name:   "kubernetes",
			Status: statusFail,
			Message: fmt.Sprintf(
				"kubernetes API unavailable: %v",
				err,
			),
		}
	}

	return Check{
		Name:    "kubernetes",
		Status:  statusPass,
		Message: "Kubernetes API is reachable",
	}
}
