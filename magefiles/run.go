//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Run(cmd string) error {
	return sh.Run("go", "run", "./cmd/"+cmd)
}
