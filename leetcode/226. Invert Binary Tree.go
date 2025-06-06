package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	return root
}

func bfs226(root *TreeNode) *TreeNode {
	q := []*TreeNode{root}

	for len(q) > 0 {
		n := q[0]
		q = q[1:]

		if n == nil {
			continue
		}
		q = append(q, n.Left, n.Right)
		n.Left, n.Right = n.Right, n.Left
	}

	return root
}

func dfs226(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	return root
}

func TestinvertTree() {

}
