package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return dfs101(root.Left, root.Right)

	return bfs101(root)
}

func dfs101(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}

	if left.Val != right.Val {
		return false
	}

	return dfs101(left.Left, right.Right) && dfs101(left.Right, right.Left)
}

func bfs101(root *TreeNode) bool {
	if root == nil {
		return false
	}

	q := []*TreeNode{}
	q = append(q, root.Left, root.Right)

	for len(q) > 0 {
		if len(q)%2 == 1 {
			return false
		}

		lq := []*TreeNode{}
		rq := []*TreeNode{}
		left, right := 0, len(q)-1
		for left < right {
			if q[left] == nil || q[right] == nil {
				if q[left] == nil && q[right] == nil {
					left++
					right--
					continue
				}
				return false
			}
			if q[left].Val != q[right].Val {
				return false
			}
			lq = append(lq, q[left].Left, q[left].Right)
			rq = append([]*TreeNode{q[right].Left, q[right].Right}, rq...)
			left++
			right--
		}
		q = append(lq, rq...)
	}

	return true
}

func TestisSymmetric() {

}
