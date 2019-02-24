package main

import (
	"path/filepath"
	"testing"
)

const testDir = "../testdata/"

func TestLineCount(t *testing.T) {
	// TODO: add test for the file with 11 lines!
	path := "10"
	if got, want := lineCount(filepath.Join(testDir, path)), 10; got != want {
		t.Errorf("lineCount(%q) = %d; want = %d", path, got, want)
	}
}
