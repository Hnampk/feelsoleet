package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	// find max value for root node
	//  return nil node for empty array
	//  recursively build tree with left & right sub arrays

	if len(nums) == 0 {
		return nil
	}

	maxIndex := 0
	for i := range len(nums) {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}

	return &TreeNode{
		Val:   nums[maxIndex],
		Left:  constructMaximumBinaryTree(nums[:maxIndex]),
		Right: constructMaximumBinaryTree(nums[maxIndex+1:]),
	}
}

func TestconstructMaximumBinaryTree() {

}
