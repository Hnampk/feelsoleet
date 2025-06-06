package leetcode

import "log"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {
	if root == nil {
		return ""
	}
	smallest := []int{}
	smallestFromLeafBacktrack(root, []int{root.Val}, &smallest)

	res := ""
	for _, c := range smallest {
		res += string('a' + c)
	}

	return res
}

func smallestFromLeafBacktrack(root *TreeNode, choices []int, smallest *[]int) {
	// 1. If it's a solution
	if root.Left == nil && root.Right == nil {
		if len(*smallest) > 0 {
			for i := 0; i < len(choices); i++ {
				if i >= len(*smallest) {
					// current smallest is shorter
					return
				}
				if choices[i] > (*smallest)[i] {
					return
				}
				if choices[i] < (*smallest)[i] {
					*smallest = choices
					return
				}
			}
		}

		*smallest = choices
		return
	}

	if root.Left != nil {
		// 2. make a choice
		choices = append([]int{root.Left.Val}, choices...)
		// 3. backtrack
		smallestFromLeafBacktrack(root.Left, choices, smallest)
		// 4. undo choice
		choices = choices[1:]
	}

	if root.Right != nil {
		// 2. make a choice
		choices = append([]int{root.Right.Val}, choices...)
		// 3. backtrack
		smallestFromLeafBacktrack(root.Right, choices, smallest)
		// 4. undo choice
		choices = choices[1:]
	}
}

// - tree with left and right sub tree
//  - leaf node has both left and right is null
//  - shortest has smaller total Val, and shorter
//  - dfs. Consider the result while reach the leaf node

func TestsmallestFromLeaf() {
	log.Printf("TestsmallestFromLeaf %s", smallestFromLeaf(&TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
	}))
}
