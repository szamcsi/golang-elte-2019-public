package todoio

import (
	"fmt"
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

	//Store("../test.csv", got)
}
