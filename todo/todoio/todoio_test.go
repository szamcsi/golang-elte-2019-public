package todoio

import (
	"testing"
)

func TestLoad(t *testing.T) {
	path := "../golang-todo.csv"
	got, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}

	want := []*Entry{}

	// TODO: implement the test
}
