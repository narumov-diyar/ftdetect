package ftdetect_test

import (
	"testing"

	"github.com/zyedidia/ftdetect"
)

// TestLoadDetectorJson checks that a detector can be read from its JSON
// description, the same format as the files in detectors/.
func TestLoadDetectorJson(t *testing.T) {
	d, err := ftdetect.LoadDetectorJson([]byte(`{"name": "go", "exts": [".go"]}`))
	if err != nil {
		t.Fatal(err)
	}
	if d.Name != "go" || len(d.Exts) != 1 || d.Exts[0] != ".go" {
		t.Errorf("unexpected detector: %+v", d)
	}
}

// TestSetPriority checks that raising a language's priority makes it win
// when two languages claim the same file.
func TestSetPriority(t *testing.T) {
	c := &ftdetect.Detector{Exts: []string{".h"}, Header: ftdetect.MustRegex("."), Name: "c"}
	cpp := &ftdetect.Detector{Exts: []string{".h"}, Header: ftdetect.MustRegex("."), Name: "cpp"}

	ds := make(ftdetect.Detectors)
	ds.RegisterDetector(c)
	ds.RegisterDetector(cpp)

	if d := ds.Detect("x.h", []byte("int main")); d == nil || d.Name != "c" {
		t.Fatalf("with equal priority the first detector should win")
	}

	ds.SetPriority("cpp", 1)
	if d := ds.Detect("x.h", []byte("int main")); d == nil || d.Name != "cpp" {
		t.Errorf("after raising its priority, cpp should win")
	}
}
