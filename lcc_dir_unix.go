//go:build !windows

package main

import (
	"log"
	"os"
	"path/filepath"
)

func getLCCDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(homeDir, ".config", "clpr")
}
