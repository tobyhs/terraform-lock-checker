package check

import (
	"io/fs"
	"path/filepath"
	"regexp"

	"github.com/hashicorp/hcl/v2/hclsimple"

	"github.com/tobyhs/terraform-lock-checker/pkg/parsing"
)

// CheckOptions contains options running .terraform.lock.hcl checks
type CheckOptions struct {
	// Number of platforms/checksums to expect
	ExpectedCount uint

	// Glob patterns of files to exclude from check
	ExclusionPatterns []string
}

// Check checks the current directory for .terraform.lock.hcl files and returns
// paths to files that don't have the expected number of provider checksums.
func Check(checkOptions *CheckOptions) ([]string, error) {
	results := []string{}
	var terraformLock parsing.TerraformLock

	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if d.Name() != ".terraform.lock.hcl" {
			return nil
		}
		if anyPatternsMatch(checkOptions.ExclusionPatterns, path) {
			return nil
		}
		if err := hclsimple.DecodeFile(path, nil, &terraformLock); err != nil {
			return err
		}
		if !checkTerraformLock(&terraformLock, checkOptions.ExpectedCount) {
			results = append(results, path)
		}
		return nil
	})

	return results, err
}

func anyPatternsMatch(patterns []string, path string) bool {
	for _, pattern := range patterns {
		matched, _ := filepath.Match(pattern, path)
		if matched {
			return true
		}
	}
	return false
}

var h1Regexp = regexp.MustCompile("^h1:")

func checkTerraformLock(lock *parsing.TerraformLock, expectedCount uint) bool {
	for _, provider := range lock.Providers {
		var actualCount uint = 0
		for _, hash := range provider.Hashes {
			if h1Regexp.MatchString(hash) {
				actualCount += 1
			}
		}
		if actualCount != expectedCount {
			return false
		}
	}
	return true
}
