// Package main is the entry point for the Kavrok CLI.
package main

import (
	"fmt"
	"os"

	"github.com/mohammedimrankasab/kavrok/internal/cli"
	"github.com/mohammedimrankasab/kavrok/internal/kubernetes"
	"github.com/mohammedimrankasab/kavrok/internal/logger"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	log, err := logger.New()
	if err != nil {
		return err
	}

	defer func() {
		_ = log.Sync()
	}()

	return cli.NewRootCommand(cli.Dependencies{
		KubernetesClientFactory: kubernetes.New,
	}).Execute()
}
