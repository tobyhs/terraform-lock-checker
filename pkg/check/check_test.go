package check

import (
	"os"
	"slices"
	"strings"
	"testing"
)

func chdir(t *testing.T, dir string) {
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(currentDir); err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
	})
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestCheckError(t *testing.T) {
	chdir(t, "../..")
	_, err := Check(&CheckOptions{ExpectedCount: 1, ExclusionPatterns: []string{}})
	if err == nil {
		t.Fatal("Expected non-nil error")
	}
	if !strings.Contains(err.Error(), "The argument \"version\" is required") {
		t.Fatalf("Incorrect error: %v", err)
	}
}

func TestCheckNonEmptyResults(t *testing.T) {
	chdir(t, "../..")
	paths, err := Check(&CheckOptions{
		ExpectedCount:     1,
		ExclusionPatterns: []string{"test/data/invalid/*"},
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	slices.Sort(paths)
	expectedPaths := []string{
		"test/data/.terraform.lock.hcl",
		"test/data/hash_4/.terraform.lock.hcl",
	}
	if !slices.Equal(paths, expectedPaths) {
		t.Errorf("Incorrect paths: %v", paths)
	}
}

func TestCheckEmptyResults(t *testing.T) {
	chdir(t, "../../test/data/hash_4")
	paths, err := Check(&CheckOptions{ExpectedCount: 4, ExclusionPatterns: []string{}})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(paths) != 0 {
		t.Errorf("Expected empty paths, got: %v", paths)
	}
}
