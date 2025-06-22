package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func areToysBalanced(root TreeNode) bool {
	return bypass(root.Left, new(int)) == bypass(root.Right, new(int))
}

func bypass(n *TreeNode, count *int) int {
	if n == nil {
		return 0
	}

	if n.HasToy {
		*count++
	}
	bypass(n.Left, count)
	bypass(n.Right, count)

	return *count
}

func createTree() *TreeNode {
	root := &TreeNode{HasToy: true}
	root.Left = &TreeNode{HasToy: true}
	root.Right = &TreeNode{HasToy: true}
	root.Left.Left = &TreeNode{HasToy: false}
	root.Left.Right = &TreeNode{HasToy: true}
	root.Right.Left = &TreeNode{HasToy: true}
	root.Right.Right = &TreeNode{HasToy: true}
	root.Left.Right.Left = &TreeNode{HasToy: false}
	root.Right.Right.Left = &TreeNode{HasToy: false}

	return root
}

func printTreeNode(node *TreeNode, depth int) {
	if node == nil {
		// Печать NIL-узла
		// fmt.Printf("%sNIL\n", strings.Repeat("    ", depth))
		return
	}

	printTreeNode(node.Right, depth+1)

	status := "0"
	if node.HasToy {
		status = "1"
	}
	fmt.Printf("%s[%s]\n", strings.Repeat("    ", depth), status)

	printTreeNode(node.Left, depth+1)
}

func PrintTree(root *TreeNode) {
	printTreeNode(root, 0)
}

func main() {
	tree := createTree()

	PrintTree(tree)
	fmt.Println(areToysBalanced(*tree))
}
