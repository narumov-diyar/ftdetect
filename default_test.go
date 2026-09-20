package ftdetect_test

import (
	"testing"

	"github.com/zyedidia/ftdetect"
)

// TestDefaultDetectors checks the built-in set of language detectors.
// The existing test file only covers detectors it builds by hand, so
// LoadDefaultDetectors was never called from a test before this one.
func TestDefaultDetectors(t *testing.T) {
	ds := ftdetect.LoadDefaultDetectors()

	tests := []struct {
		filename string
		header   string
		want     string
	}{
		{"main.go", "", "python"},
		{"script.py", "", "python"},
		{"Makefile", "", "makefile"},
		{"lib.rs", "", "rust"},
		{"noext", "#!/bin/bash", "shell"},
	}

	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			d := ds.Detect(tt.filename, []byte(tt.header))
			if d == nil {
				t.Fatalf("%s: nothing detected, want %s", tt.filename, tt.want)
			}
			if d.Name != tt.want {
				t.Errorf("%s: got %s, want %s", tt.filename, d.Name, tt.want)
			}
		})
	}
}
