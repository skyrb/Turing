package main

import (
	"fmt"
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"math"
	"strings"
)

// Binary tree node structure
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// Create a binary tree map representation.
func createTreeMap() map[int]map[string]int {
	root := TreeNode{Value: 1}
	root.Left = &TreeNode{Value: 2}
	root.Right = &TreeNode{Value: 3}
	root.Left.Left = &TreeNode{Value: 4}
	root.Left.Right = &TreeNode{Value: 5}
	root.Right.Left = &TreeNode{Value: 6}

	treeMap := make(map[int]map[string]int)
	addToMap(&root, treeMap)
	return treeMap
}

func addToMap(node *TreeNode, treeMap map[int]map[string]int) {
	if node == nil {
		return
	}
	treeMap[node.Value] = map[string]int{}
	if node.Left != nil {
		treeMap[node.Value]["left"] = node.Left.Value
	}
	if node.Right != nil {
		treeMap[node.Value]["right"] = node.Right.Value
	}
	addToMap(node.Left, treeMap)
	addToMap(node.Right, treeMap)
}

// Calculate the width of the binary tree.
func treeWidth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + int(math.Max(float64(treeWidth(node.Left)), float64(treeWidth(node.Right))))
}

// Calculate the depth of the binary tree.
func treeDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + int(math.Max(float64(treeDepth(node.Left)), float64(treeDepth(node.Right))))
}

// Calculate the X position for a node in the tree.
func xPosition(node *TreeNode, x int, width int) int {
	if node == nil {
		return 0
	}
	return x + (width-1)/2
}

// Draw the binary tree in the terminal.
func drawTree(treeMap map[int]map[string]int, root int, x, y int, width, depth int, g *widgets.Grid) {
	if _, ok := treeMap[root]; !ok {
		return
	}
	label := fmt.Sprintf("%d", root)
	g.SetCell(x, y, &widgets.Cell{Text: label, Bg: termui.ColorBlue, Fg: termui.ColorWhite})

	leftChild, ok := treeMap[root]["left"]
	if ok {
		drawTree(treeMap, leftChild, x-int(math.Ceil(float64(width)/2)), y+1, width/2, depth-1, g)
		g.SetCell(x-int(math.Ceil(float64(width)/2)), y+1, &widgets.Cell{Text: "/", Bg: termui.ColorBlue, Fg: termui.ColorWhite})
	}

	rightChild, ok := treeMap[root]["right"]
	if ok {
		drawTree(treeMap, rightChild, x+int(math.Ceil(float64(width)/2)), y+1, width/2, depth-1, g)
		g.SetCell(x+int(math.Ceil(float64(width)/2)), y+1, &widgets.Cell{Text: "\\", Bg: termui.ColorBlue, Fg: termui.ColorWhite})
	}
}

func main() {
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	treeMap := createTreeMap()

	root := 1
	width := treeWidth(&TreeNode{Value: root})
	depth := treeDepth(&TreeNode{Value: root})
	x := xPosition(&TreeNode{Value: root}, 0, width)
	y := 0

	g := widgets.NewGrid()
	g.Rows = make([]int, depth)
	g.Cols = make([]int, width)
	g.SetCell(x, y, &widgets.Cell{Text: fmt.Sprintf("%d", root), Bg: termui.ColorBlue, Fg: termui.ColorWhite})
	drawTree(treeMap, root, x, y, width, depth, g)

	termui.Root.Add(g)
	termui.Render()

	for {
		e := termui.PollEvent()
		switch e.Type {
		case termui.KeyEvent:
			if e.Key == termui.KeyEsc {
				return
			}
		case termui.CloseEvent:
			return
		}
		termui.Render()
	}
}