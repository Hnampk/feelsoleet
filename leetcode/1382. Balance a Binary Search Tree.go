package leetcode

import (
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
func balanceBST(root *TreeNode) *TreeNode {
	// in-order traversal => sorted array
	sortedNodes := inOrder(root)

	// build the tree from sorted array
	return buildTree1382(sortedNodes)
}

func inOrder(root *TreeNode) []*TreeNode {
	res := []*TreeNode{}
	if root == nil {
		return res
	}

	res = append(res, inOrder(root.Left)...)
	res = append(res, root)
	res = append(res, inOrder(root.Right)...)

	return res
}

func buildTree1382(nodes []*TreeNode) *TreeNode {
	if len(nodes) == 0 {
		return nil
	}

	m := len(nodes) / 2

	root := &TreeNode{
		Val:   nodes[m].Val,
		Left:  buildTree1382(nodes[:m]),
		Right: buildTree1382(nodes[m+1:]),
	}

	return root
}

func TestbalanceBST() {

	type Case struct {
		root *TreeNode

		expected *TreeNode
	}

	cases := []Case{
		{
			root: &TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:  2,
					Left: nil,
					Right: &TreeNode{
						Val:  3,
						Left: nil,
						Right: &TreeNode{
							Val:   4,
							Left:  nil,
							Right: nil,
						},
					},
				},
			},
			expected: &TreeNode{},
		},
		{
			root: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			expected: &TreeNode{},
		},
	}

	for _, testcase := range cases {
		// build tree node

		res := balanceBST(testcase.root)

		// root := testcase.expected

		log.Printf("MATCHED. %+v result: %+v, expected: %+v\n\n", testcase.root, res, testcase.expected)
	}
}
