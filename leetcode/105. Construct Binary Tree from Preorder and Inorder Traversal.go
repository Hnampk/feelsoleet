package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree105(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	// root value is the first element of preorder
	root := &TreeNode{
		Val: preorder[0],
	}

	// find index of root value in inorder, split the array into 2 parts: left & right
	//	use the len of 2 parts to find the preorder array of sub-left tree and sub-right tree
	//	recursively build the sub-left & sub-right tree from preorder and inorder

	var rootIndex int
	for i, val := range inorder {
		if val == root.Val {
			rootIndex = i
			break
		}
	}
	root.Left = buildTree105(preorder[1:rootIndex+1], inorder[:rootIndex])
	root.Right = buildTree105(preorder[rootIndex+1:], inorder[rootIndex+1:])

	return root
}

func TestbuildTree() {

}
