package leetcode

// Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func sortedArrayToBST(nums []int) *TreeNode {
	return dfs108(nums)
}

func bfs108(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	type Node struct {
		val  *TreeNode
		nums []int
	}

	m := len(nums) / 2
	root := &TreeNode{
		Val:   nums[m],
		Left:  nil,
		Right: nil,
	}

	q := []*Node{{root, nums}}

	for len(q) > 0 {
		n := q[0]
		q = q[1:]

		m := len(n.nums) / 2
		left := n.nums[:m]
		right := n.nums[m+1:]

		if len(left) > 0 {
			// left sub-tree
			m := len(left) / 2
			leftNode := &TreeNode{
				Val:   left[m],
				Left:  nil,
				Right: nil,
			}
			n.val.Left = leftNode

			q = append(q, &Node{leftNode, left})
		}
		if len(right) > 0 {
			// right sub-tree
			m := len(right) / 2
			rightNode := &TreeNode{
				Val:   right[m],
				Left:  nil,
				Right: nil,
			}
			n.val.Right = rightNode

			q = append(q, &Node{rightNode, right})
		}
	}

	return root
}

func dfs108(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	m := len(nums) / 2
	return &TreeNode{
		Val:   nums[m],
		Left:  dfs108(nums[:m]),
		Right: dfs108(nums[m+1:]),
	}
}

func TestsortedArrayToBST() {

}
