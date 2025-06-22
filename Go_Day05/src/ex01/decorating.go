package main

import (
	"fmt"
	"strings"
	"time"
)

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func createTree() *TreeNode {
	root := &TreeNode{HasToy: true}
	root.Left = &TreeNode{HasToy: false}
	root.Right = &TreeNode{HasToy: true}
	root.Left.Left = &TreeNode{HasToy: false}
	root.Left.Right = &TreeNode{HasToy: true}
	root.Right.Left = &TreeNode{HasToy: false}
	root.Right.Right = &TreeNode{HasToy: true}
	root.Left.Right.Left = &TreeNode{HasToy: false}
	root.Right.Right.Left = &TreeNode{HasToy: true}

	return root
}

func unrollGarland(root *TreeNode) []bool {
	if root == nil {
		return nil
	}
	var garland []bool

	queue := []*TreeNode{root} // currentLevelNodes == 0

	currentLevelNodes, nextLevelNodes, countLevel := 1, 0, 0

	for len(queue) > 0 {
		var node *TreeNode
		for currentLevelNodes > 0 {
			node = queue[0]
			queue = queue[1:]
			if countLevel%2 == 0 {
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			} else {
				if node.Left != nil {
					queue = append(queue, node.Right)
				}
				if node.Right != nil {
					queue = append(queue, node.Left)
				}
			}

			currentLevelNodes--

			if node.Right != nil {
				nextLevelNodes++
			}
			if node.Left != nil {
				nextLevelNodes++
			}

			garland = append(garland, node.HasToy)
		}

		currentLevelNodes = nextLevelNodes
		nextLevelNodes = 0
		countLevel++
	}

	return garland
}

func main() {
	start := time.Now()
	root := createTree()

	garland := unrollGarland(root)
	PrintTree(root)
	fmt.Println(garland)

	duration := time.Since(start)
	fmt.Println("Время выполнения:", duration)
}


func printTreeNode(node *TreeNode, depth int) {
	if node == nil {
		// Печать NIL-узла
		// fmt.Printf("%sNIL\n", strings.Repeat("    ", depth))
		return
	}

	printTreeNode(node.Right, depth+1)

	status := "false"
	if node.HasToy {
		status = "true"
	}
	fmt.Printf("%s[%s]\n", strings.Repeat("    ", depth), status)

	printTreeNode(node.Left, depth+1)
}

func PrintTree(root *TreeNode) {
	printTreeNode(root, 0)
}

// обход дерева в ширину
// func unrollGarland(root *TreeNode) []bool {
// 	if root == nil {
// 		return nil
// 	}
// 	var garland []bool

// 	queue := []*TreeNode{root}

// 	for len(queue) > 0 {
// 		node := queue[0]
// 		queue = queue[1:]

// 		garland = append(garland, node.HasToy)

// 		if node.Left != nil {
// 			queue = append(queue, node.Left)
// 		}

// 		if node.Right != nil {
// 			queue = append(queue, node.Right)
// 		}
// 	}

// 	return garland
// }
