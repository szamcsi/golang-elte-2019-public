// Package todoio handles importing/exporting of TODO entries from/to CSV file.
package todoio

import (
	"time"
)

// Type Entry describes a basic TODO entry.
type Entry struct {
	Done     bool
	Text     string
	Deadline time.Time // optional
}

// LoadEntries imports TODO entries from a CSV file.
func Load(path string) ([]*Entry, error) {
	return nil, nil
}

// StoreEntries exports TODO entries to a CSV file.
func Store(path string, entries []*Entry) error {
	return nil
}
