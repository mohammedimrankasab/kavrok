// Package main is the entry point for the Kavrok CLI.
package main

import (
	"os"

	"github.com/mohammedimrankasab/kavrok/internal/cli"
	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
	"github.com/mohammedimrankasab/kavrok/internal/logger"
	"go.uber.org/zap"
)

func main() {
	log, err := logger.New()
	if err != nil {
		os.Exit(1)
	}

	exitCode := run(log)

	_ = log.Sync()

	if exitCode != 0 {
		os.Exit(exitCode)
	}
}

func run(log *zap.Logger) int {
	err := cli.Execute(cli.Dependencies{
		KubernetesClientFactory: kubernetes.New,
	})
	if err != nil {
		log.Error("application failed", zap.Error(err))
		return 1
	}

	return 0
}
