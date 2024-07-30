package check

// CheckOptions contains options running .terraform.lock.hcl checks
type CheckOptions struct {
	// Number of platforms/checksums to expect
	ExpectedCount uint

	// Glob patterns of files to exclude from check
	ExclusionPatterns []string
}
