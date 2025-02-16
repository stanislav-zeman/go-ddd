//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Compose() error {
	return sh.Run("docker", "compose", "-f", "compose.dev.yaml", "--profile", "all", "up")
}
