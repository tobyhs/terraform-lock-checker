package check

import (
	"reflect"
	"testing"
)

func TestParseCliArgsDefaults(t *testing.T) {
	checkOptions, err := ParseCliArgs([]string{})
	if err != nil {
		t.Fatalf("err is not nil: %v", err)
	}
	wanted := CheckOptions{ExpectedCount: 1, ExclusionPatterns: []string{}}
	if !reflect.DeepEqual(checkOptions, &wanted) {
		t.Errorf("Incorrect result: %v", checkOptions)
	}
}

func TestParseCliArgs(t *testing.T) {
	checkOptions, err := ParseCliArgs([]string{
		"-expected-count", "4",
		"-exclusion-patterns", "abc,def",
	})

	if err != nil {
		t.Fatalf("err is not nil: %v", err)
	}
	wanted := CheckOptions{
		ExpectedCount:     4,
		ExclusionPatterns: []string{"abc", "def"},
	}
	if !reflect.DeepEqual(checkOptions, &wanted) {
		t.Errorf("Incorrect result: %v", checkOptions)
	}
}

func TestParseCliArgsError(t *testing.T) {
	checkOptions, err := ParseCliArgs([]string{"-expected-count"})

	if checkOptions != nil {
		t.Errorf("checkOptions is not nil: %v", checkOptions)
	}
	if err.Error() != "flag needs an argument: -expected-count" {
		t.Errorf("error is incorrect: %s", err.Error())
	}
}
