//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Lint() error {
	return sh.RunV("go", "run", "github.com/golangci/golangci-lint/cmd/golangci-lint", "run")
}
