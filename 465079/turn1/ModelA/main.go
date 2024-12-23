package main

import (
	"fmt"
	"strings"
)

// TreeNode represents a node in the binary tree.
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// Create a binary tree map representation.
func createTreeMap() map[int]map[string]int {
	root := 1
	treeMap := make(map[int]map[string]int)
	treeMap[root] = map[string]int{
		"left":  2,
		"right": 3,
	}
	treeMap[2] = map[string]int{
		"left":  4,
		"right": 5,
	}
	treeMap[3] = map[string]int{
		"left":  6,
	}
	return treeMap
}

// Print the binary tree in a visual format.
func printTree(treeMap map[int]map[string]int, root int, level int, prefix string) {
	if _, ok := treeMap[root]; !ok {
		return
	}

	if level > 0 {
		fmt.Println(prefix, root)
	}

	prefix += strings.Repeat(" ", 4) // Indent for the next level
	leftChild, ok := treeMap[root]["left"]
	if ok {
		printTree(treeMap, leftChild, level+1, prefix+"- ")
	}

	rightChild, ok := treeMap[root]["right"]
	if ok {
		printTree(treeMap, rightChild, level+1, prefix+"+ ")
	}
}

func main() {
	treeMap := createTreeMap()
	printTree(treeMap, 1, 0, "")
}