// Package todoio handles importing/exporting of TODO entries from/to CSV file.
package todoio

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Type Entry describes a basic TODO entry.
type Entry struct {
	Done     bool
	Text     string
	Deadline time.Time
}

// LoadEntries imports TODO entries from a CSV file.
func Load(path string) ([]*Entry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}
	var entries []*Entry
	for i, line := range lines {
		if i == 0 {
			continue
		}
		if len(line) != 3 {
			return nil, fmt.Errorf("Line %d: wrong format error", i+1)
		}
		var entry Entry
		entry.Done, err = strconv.ParseBool(line[0])
		entry.Text = line[1]
		entry.Deadline, err = time.Parse("2006-01-02", line[2])
		if err != nil {
			return nil, err
		}
		entries = append(entries, &entry)
	}
	return entries, nil
}

// StoreEntries exports TODO entries to a CSV file.
func Store(path string, entries []*Entry) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString("DONE,TEXT,DEADLINE\n")
	for _, entry := range entries {
		_, err := file.WriteString(strconv.FormatBool(entry.Done) + "," + entry.Text + "," + entry.Deadline.Format("2006-01-02") + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}
