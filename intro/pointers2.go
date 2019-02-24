package main

import "fmt"

type treeNode struct {
	value       string
	left, right *treeNode
}

func toString(t *treeNode) string {
	if t == nil {
		return "nil"
	}
	return fmt.Sprintf("%q: {%s, %s}", t.value, toString(t.left), toString(t.right))
}

func deepCopyUpper(root *treeNode) *treeNode {
	// replace this function body. deepCopyUpper should create a deep copy
	// of the tree with all values converted to uppercase. Use strings.ToUpper.
	return nil
}

func main() {
	one := &treeNode{
		value: "foo",
		left: &treeNode{
			value: "bar",
			left: &treeNode{
				value: "baz",
			},
			right: &treeNode{
				value: "qux",
			},
		},
		right: &treeNode{
			value: "quux",
			right: &treeNode{
				value: "quuz",
			},
		},
	}
	two := deepCopyUpper(one)
	fmt.Println(toString(one))
	fmt.Println(toString(two))
}
