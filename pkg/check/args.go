package check

import (
	"flag"
	"strings"
)

// ParseCliArgs parses the given command line arguments
func ParseCliArgs(arguments []string) (*CheckOptions, error) {
	flagSet := flag.NewFlagSet("terraform-lock-checker", flag.ContinueOnError)
	expectedCount := flagSet.Uint("expected-count", 1, "Number of platforms/checksums to expect")
	exclusionPatternsStr := flagSet.String("exclusion-patterns", "", "Comma separated glob patterns of files to exclude from checking")
	if err := flagSet.Parse(arguments); err != nil {
		return nil, err
	}

	exclusionPatterns := []string{}
	if *exclusionPatternsStr != "" {
		exclusionPatterns = strings.Split(*exclusionPatternsStr, ",")
	}

	return &CheckOptions{
		ExpectedCount:     *expectedCount,
		ExclusionPatterns: exclusionPatterns,
	}, nil
}
