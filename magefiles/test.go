//go:build mage

package main

import (
	"github.com/magefile/mage/sh"
)

func Test() error {
	return sh.RunV("go", "test", "-race", "-cover", "-coverprofile", "cp.out", "./...")
}
