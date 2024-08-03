package check

import (
	"strings"
	"testing"
)

func TestRunErrorParseCliArgs(t *testing.T) {
	exitCode := Run(
		[]string{"-expected-count"},
		new(strings.Builder),
		new(strings.Builder),
	)
	if exitCode != 2 {
		t.Errorf("Incorrect exit code: %d", exitCode)
	}
}

func TestRunErrorCheck(t *testing.T) {
	var stdout, stderr strings.Builder
	chdir(t, "../..")
	exitCode := Run([]string{}, &stdout, &stderr)
	if exitCode != 2 {
		t.Errorf("Incorrect exit code: %d", exitCode)
	}

	if stdout.String() != "" {
		t.Errorf("Expected empty stdout, got: %s", stdout.String())
	}
	if !strings.Contains(stderr.String(), "The argument \"version\" is required") {
		t.Errorf("Incorrect stderr, got: %s", stderr.String())
	}
}

func TestRunFailingPaths(t *testing.T) {
	var stdout, stderr strings.Builder
	chdir(t, "../..")
	exitCode := Run(
		[]string{
			"-expected-count", "4",
			"-exclusion-patterns", "test/data/invalid/*",
		},
		&stdout,
		&stderr,
	)
	if exitCode != 1 {
		t.Errorf("Incorrect exit code: %d", exitCode)
	}

	if stderr.String() != "" {
		t.Errorf("Expected empty stderr, got: %s", stderr.String())
	}

	expectedStdout := strings.Join([]string{
		"The following files have an unexpected number of provider checksums:\n",
		"test/data/hash_1/.terraform.lock.hcl\n",
	}, "")
	if stdout.String() != expectedStdout {
		t.Errorf("Incorrect stdout, got: %s", stdout.String())
	}
}

func TestRunSuccess(t *testing.T) {
	var stdout, stderr strings.Builder
	chdir(t, "../../test/data/hash_1")
	exitCode := Run([]string{}, &stdout, &stderr)
	if exitCode != 0 {
		t.Errorf("Incorrect exit code: %d", exitCode)
	}

	if stdout.String() != "" {
		t.Errorf("Expected empty stdout, got: %s", stdout.String())
	}
	if stderr.String() != "" {
		t.Errorf("Expected empty stderr, got: %s", stderr.String())
	}
}
