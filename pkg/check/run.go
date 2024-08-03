package check

import (
	"fmt"
	"io"
)

// Run is an entrypoint function for the command line program
func Run(arguments []string, stdout io.Writer, stderr io.Writer) int {
	checkOptions, err := ParseCliArgs(arguments)
	if err != nil {
		return 2
	}
	failingPaths, err := Check(checkOptions)
	if err != nil {
		fmt.Fprintln(stderr, err.Error())
		return 2
	}

	if len(failingPaths) > 0 {
		fmt.Fprintln(stdout, "The following files have an unexpected number of provider checksums:")
		for _, path := range failingPaths {
			fmt.Fprintln(stdout, path)
		}
		return 1
	}
	return 0
}
