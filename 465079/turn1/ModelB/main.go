package main  
import (  
    "fmt"
    "github.com/faln/go-tree"
    "github.com/fatih/color"
)  

//Node structure  
type Node struct {  
    Key int
    Value string
    Left *Node  
    Right *Node  
}  

func main() {  
    // Create a map to represent the binary tree data  
    treeData := map[int]string{  
        10: "Root",  
        5: "Left Child of 10",  
        15: "Right Child of 10",  
        3: "Left Child of 5",  
        7: "Right Child of 5",  
        13: "Left Child of 15",  
        17: "Right Child of 15",  
    }  

    // Convert the map-based data into a tree structure using buildTree function
    root := buildTree(treeData)

    // Display the tree visually
    showTree(root)
}
// buildTree function: Constructs the binary tree from the provided map.
func buildTree(treeData map[int]string) *Node {
    // Initialize the tree structure with the root node
    root := &Node{}

    // Iterate through the map to add nodes to the tree
    for key, value := range treeData {
        // If the key is not present in the tree, it is a new node, so add it as the root
        if _, found := findNode(root, key); !found {
            root = insertNode(root, key, value)
        }
    }
    return root
}

// Finds the node with the given key in the binary tree.
func findNode(root *Node, key int) (*Node, bool) {
    if root == nil || root.Key == key {
        return root, root != nil
    }
    if key < root.Key {
        return findNode(root.Left, key)
    }
    return findNode(root.Right, key)
}

// Inserts a new node with the given key and value into the binary tree.
func insertNode(root *Node, key int, value string) *Node {
    if root == nil {
        return &Node{Key: key, Value: value}
    }
    if key < root.Key {
        root.Left = insertNode(root.Left, key, value)
    } else {
        root.Right = insertNode(root.Right, key, value)
    }
    return root
}
  
// Display the tree visually using "github.com/faln/go-tree" package.
func showTree(root *Node) {
    t := tree.New()
    if root != nil {
        appendNode(t, root)
    }
    fmt.Println(t)
}

// Recursively append nodes to the tree structure
func appendNode(parent tree.Node, node *Node) {
    if node == nil {
        return
    }

    n := parent.AppendChild(color.HiYellowString(fmt.Sprintf("%d: %s", node.Key, node.Value)))
    appendNode(n, node.Left)
    appendNode(n, node.Right)
}