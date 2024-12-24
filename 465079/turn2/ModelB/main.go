
package main

import (
        "fmt"
        "math/rand"
        "time"

        "github.com/gizak/termui/v3"
        "github.com/gizak/termui/v3/widgets"
)

// TreeNode represents a node in a binary tree.
type TreeNode struct {
        Val   int
        Left  *TreeNode
        Right *TreeNode
}

// buildTree builds a random binary tree with a given number of nodes.
func buildTree(n int) *TreeNode {
        if n == 0 {
                return nil
        }
        val := rand.Intn(100)
        return &TreeNode{
                Val:   val,
                Left:  buildTree(n / 2),
                Right: buildTree(n - n/2 - 1),
        }
}

// maxDepth calculates the maximum depth of a binary tree.
func maxDepth(root *TreeNode) int {
        if root == nil {
                return 0
        }
        return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

// max returns the maximum of two integers.
func max(a, b int) int {
        if a >= b {
                return a
        }
        return b
}

// visualizeTree visualizes a binary tree using termui widgets.
func visualizeTree(root *TreeNode, width, height int) {
        // Initialize termui
        if err := termui.Init(); err != nil {
                panic(err)
        }
        defer termui.Close()

        // Create a new grid
        g := termui.NewGrid()
        termWidth, termHeight := termui.TerminalDimensions()
        g.SetRect(0, 0, termWidth, termHeight)

        // Create a text widget to display the tree
        t := widgets.NewParagraph()
        t.Title = "Binary Tree Visualization"
        t.BorderStyle.Fg = termui.ColorWhite
        t.TextStyle.Fg = termui.ColorWhite

        // Add the text widget to the grid
        g.Set(
                termui.NewRow(
                        0.05,
                        termui.NewCol(1.0, 0, t),
                ),
        )

        // Render the initial grid
        termui.Render(g)

        // Function to print the tree in a visual format to the text widget
        printTree := func(node *TreeNode, x, y, dx int) {
                if node == nil {
                        return
                }

                // Calculate the position to print the current node
                pos := x + (y*width)/height

                // Clear the previous character at the calculated position
                t.Text = t.Text[:pos] + " " + t.Text[pos+1:]

                // Print the current node's value
                t.Text = t.Text[:pos] + fmt.Sprintf("%d", node.Val) + t.Text[pos+1:]

                // Recursively print the left and right subtrees
                printTree(node.Left, x, y+1, dx/2)
                printTree(node.Right, x+dx/2, y+1, dx/2)
        }

        // Clear the text widget
        t.Text = ""

        // Calculate the width and height for the tree visualization
        maxDepth := maxDepth(root)
        treeWidth := int(math.Pow(2, float64(maxDepth))) - 1
        treeHeight := maxDepth

        // Adjust the width and height to fit within the terminal window
        width = int(float64(width) * float64(termWidth) / float64(treeWidth))
        height = int(float64(height) * float64(termHeight) / float64(treeWidth))

        // Print the tree in the text widget
        printTree(root, 0, 0, width)

        // Render the updated text widget
        termui.Render(g)