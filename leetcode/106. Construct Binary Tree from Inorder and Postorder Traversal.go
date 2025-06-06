package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	// inorder: [left, root, right]
	// postorder: [left, right, root]

	if len(inorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{
			Val: inorder[0],
		}
	}

	root := postorder[len(postorder)-1]
	var rootIndex int
	for i, val := range inorder {
		if val == root {
			rootIndex = i
			break
		}
	}

	return &TreeNode{
		Val:   root,
		Left:  buildTree(inorder[:rootIndex], postorder[:rootIndex]),
		Right: buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1]),
	}
}

func TestbuildTree106() {

}
