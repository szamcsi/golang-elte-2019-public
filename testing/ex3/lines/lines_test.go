package lines

import (
	"path/filepath"
	"testing"
)

const testDir = "../../testdata/"

func TestCount(t *testing.T) {
	for _, tc := range []struct {
		path  string
		count int
	}{
		{"10", 10},
		{"11", 11},
	} {
		// TODO: check with a not-existing file, if an error is returned from Count().
		if lc, _ := Count(filepath.Join(testDir, tc.path)); lc != tc.count {
			t.Errorf("Count(%q) = %d; want = %d", tc.path, lc, tc.count)
		}
	}
}
