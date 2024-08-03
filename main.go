package main

import (
	"os"

	"github.com/tobyhs/terraform-lock-checker/pkg/check"
)

func main() {
	os.Exit(check.Run(os.Args[1:], os.Stdout, os.Stderr))
}
