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
	Deadline time.Time // optional
}

// LoadEntries imports TODO entries from a CSV file.
func Load(path string) ([]*Entry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil,err
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil,err
	}
	var output []*Entry
	// Loop through lines & turn into object
	for _, line := range lines {
		if len (line) < 3 {
			return nil, fmt.Errorf("Malformed error: %g", line)
		}
		b, _ := strconv.ParseBool(line[0])
		now := time.Now()
		now.Format(line[2])
		var data = Entry{
			Done:     b,
			Text:     line[1],
			Deadline: now,
		}
		output = append(output, &data)
	}

	return output, nil
}

// StoreEntries exports TODO entries to a CSV file.
func Store(path string, entries []*Entry) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, entry := range entries {
		n, err := file.WriteString(strconv.FormatBool(entry.Done) + "," + entry.Text + "," + entry.Deadline.Format("2006-01-02") + "\n")
		if err != nil {
			return err
		}
		println(" Line written to file", n)
	}
	return err
}
