package leetcode

import (
	"fmt"
	"log"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	// backtrack
	return backtrack257(root, []*TreeNode{})
}

func backtrack257(root *TreeNode, choices []*TreeNode) []string {
	// 2.1 make a choice
	choices = append(choices, root)
	defer func() {
		// undo choice
		choices = choices[:len(choices)-1] // pop
	}()

	// 1. record if it's the solution
	if root.Left == nil && root.Right == nil {
		res := fmt.Sprintf("%d", choices[0].Val)
		for i := 1; i < len(choices); i++ {
			res = fmt.Sprintf("%s->%d", res, choices[i].Val)
		}

		return []string{res}
	}

	left := []string{}
	right := []string{}
	if root.Left != nil {
		// 2.2 recursively call backtrack
		left = backtrack257(root.Left, choices)
	}
	if root.Right != nil {
		// 2.2 recursively call backtrack
		right = backtrack257(root.Right, choices)
	}

	return append(left, right...)
}

func backtrack257Template(root *TreeNode, choices []*TreeNode) []string {
	// 1. record if it's a solution
	if root.Left == nil && root.Right == nil {
		choices = append(choices, root)

		res := fmt.Sprintf("%d", choices[0].Val)
		for i := 1; i < len(choices); i++ {
			res = fmt.Sprintf("%s->%d", res, choices[i].Val)
		}
		choices = choices[:len(choices)-1] // pop

		return []string{res}
	}

	// 2. make choice
	res := []string{}
	choices = append(choices, root)
	if root.Left != nil {
		// 3. backtrack
		res = append(res, backtrack257Template(root.Left, choices)...)
	}
	if root.Right != nil {
		// 3. backtrack
		res = append(res, backtrack257Template(root.Right, choices)...)
	}
	// 4. undo choice
	choices = choices[:len(choices)-1] // pop

	return res
}

func TestbinaryTreePaths() {
	r := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	res := binaryTreePaths(r)

	for _, r := range res {
		log.Printf(r)
	}
}
