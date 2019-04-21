package todoio

import (
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	path := "../golang-todo.csv"
	got, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	path2 := "../golang-todo2.csv"
	err = Store(path2, got)
	if err != nil {
		t.Fatal(err)
	}

	got2, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	println(reflect.DeepEqual(got2, got))
}
