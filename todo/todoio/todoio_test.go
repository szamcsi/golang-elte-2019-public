package todoio

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func StrToDate(str string) time.Time {
	date, err := time.Parse("2006-01-02", str)
	if err != nil {
		fmt.Errorf("Error: date in wrong format: %s", str)
	}
	return date
}

func TestLoad(t *testing.T) {
	path := "../golang-todo.csv"
	got, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}

	want := []*Entry{
		{false, "display entries from a CSV import", StrToDate("2019-04-04")},
		{false, "web interface and CLI sent for review", StrToDate("2019-05-10")},
	}

	for i, line := range got {
		if !cmp.Equal(line, want[i]) {
			if line.Done != want[i].Done {
				t.Errorf("Line %d: got = %t; want = %t", i+1, line.Done, want[i].Done)
			}
			if line.Text != want[i].Text {
				t.Errorf("Line %d: got = %s; want = %s", i+1, line.Text, want[i].Text)
			}
			if !cmp.Equal(line.Deadline, want[i].Deadline) {
				t.Errorf("Line %d: got = %s; want = %s", i+1, line.Deadline.Format("2006-01-02"), want[i].Deadline.Format("2006-01-02"))
			}
		}
	}
}

func TestStore(t *testing.T) {
	entries := []*Entry{
		{false, "display entries from a CSV import", StrToDate("2019-04-04")},
		{false, "web interface and CLI sent for review", StrToDate("2019-05-10")},
	}

	path := "../test.csv"
	err := Store(path, entries)
	if err != nil {
		t.Fatal(err)
	}

	want := [][]string{
		{"DONE", "TEXT", "DEADLINE"},
		{"false", "display entries from a CSV import", "2019-04-04"},
		{"false", "web interface and CLI sent for review", "2019-05-10"},
	}

	file, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	got, err := csv.NewReader(file).ReadAll()
	if err != nil {
		t.Fatal(err)
	}
	for i, line := range got {
		if !cmp.Equal(line, want[i]) {
			if line[0] != want[i][0] {
				t.Errorf("Line %d, column %d: got = %s; want = %s", i+1, 1, line[0], want[i][0])
			}
			if line[1] != want[i][1] {
				t.Errorf("Line %d, column %d: got = %s; want = %s", i+1, 2, line[1], want[i][1])
			}
			if line[2] != want[i][2] {
				t.Errorf("Line %d, column %d: got = %s; want = %s", i+1, 3, line[2], want[i][2])
			}
		}
	}
}
