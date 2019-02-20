package main

import (
	"fmt"
	"strings"
)

type treeNode struct {
	value    string
	children []*treeNode
}

func toString(t *treeNode) string {
	if t == nil {
		return "nil"
	}
	var out []string
	for _, c := range t.children {
		out = append(out, toString(c))
	}
	return fmt.Sprintf("%q: {%s}", t.value, strings.Join(out, ", "))
}

// creates a deep copy of the treeNode with all values converted to uppercase
func deepCopyUpper(root *treeNode) *treeNode {
	// TODO: write function body
	return nil
}

func main() {
	one := &treeNode{
		value: "foo",
		children: []*treeNode{
			{
				value: "bar",
				children: []*treeNode{
					{
						value: "baz",
					},
					{
						value: "qux",
					},
				},
			},
			{
				value: "quux",
			},
		},
	}
	two := deepCopyUpper(one)
	fmt.Println(toString(one))
	fmt.Println(toString(two))
}
