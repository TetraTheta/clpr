//go:build windows

package main

import (
	"log"
	"os"
	"path/filepath"
)

func getLCCDir() string {
	LAD := os.Getenv("LocalAppData")
	if LAD == "" {
		// This should not happen
		log.Fatal("%LocalAppData% is not found.")
	}
	return filepath.Join(LAD, "clpr")
}
