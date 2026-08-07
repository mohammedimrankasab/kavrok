package main

import (
	"log"
	"os"

	"github.com/mohammedimrankasab/kavrok/internal/app"
)

func main() {
	if err := app.New().Execute(); err != nil {
		log.Printf("application failed: %v", err)
		os.Exit(1)
	}
}
