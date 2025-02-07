//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Build(cmd string) error {
	return sh.Run("go", "build", "-o", "./bin/"+cmd, "./cmd/"+cmd)
}
