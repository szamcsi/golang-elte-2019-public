package lines

import (
	// TODO: use "github.com/google/go-cmp/cmp"
	"path/filepath"
	"testing"
)

const testDir = "../../testdata/"

func TestLoad(t *testing.T) {
	for _, tc := range []struct {
		path    string
		lines   []string
		wantErr bool
	}{
		{"10", []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, false},
		{"11", []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}, false},
		{"12", nil, true},
	} {
		lines, err := load(filepath.Join(testDir, tc.path))
		if tc.wantErr != (err != nil) {
			t.Errorf("load(%q) err != nil is %v; want %v", tc.path, err != nil, tc.wantErr)
			continue
		}
		// START OMIT
		if true { // TODO: compare lines and tc.lines
			t.Errorf("load(%q) returned diff (-want +got):\n%s", tc.path,
				"") // TODO: print details of how the result differs from expected
		}
		// END OMIT
	}
}

func TestMinMaxCount(t *testing.T) {
	for _, tc := range []struct {
		lines []string
		mmc   *MinMaxCount
	}{
		{nil, nil},
		{[]string{"1"}, &MinMaxCount{1, 1, 1}},
		{[]string{"1", "1", "1"}, &MinMaxCount{1, 1, 3}},
		{[]string{"1", "22"}, &MinMaxCount{1, 2, 2}},
	} {
		if mmc := minMaxCount(tc.lines); true { // TODO: compare mmc and tc.mmc
			t.Errorf("minMaxCount(%v) = %v; want = %v", tc.lines, mmc, tc.mmc)
		}
	}
}

func TestCount(t *testing.T) {
	for _, tc := range []struct {
		path    string
		mmc     *MinMaxCount
		wantErr bool
	}{
		{"9", &MinMaxCount{1, 9, 9}, false},
		{"10", &MinMaxCount{1, 2, 10}, false},
		{"80", &MinMaxCount{0, 80, 9}, false},
		{"12", nil, true},
	} {
		mmc, err := Count(filepath.Join(testDir, tc.path))
		if tc.wantErr != (err != nil) {
			t.Errorf("Count(%q) err != nil is %v; want %v", tc.path, err != nil, tc.wantErr)
			continue
		}
		if true { // TODO: compare mmc and tc.mmc
			t.Errorf("Count(%q) = %v, _; want = %v, _", tc.path, mmc, tc.mmc)
		}
	}
}
