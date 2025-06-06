package leetcode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	// bfs
	// put left & right nodes to queue
	//  loop through nodes in queue, if there's items in queue, they're on the same line => make next pointers

	if root == nil {
		return nil
	}
	q := []*Node{root}

	for len(q) > 0 {

		size := len(q)
		for i := range size {
			n := q[i]

			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
			if i == size-1 {
				break
			}

			n.Next = q[i+1]
		}

		q = q[size:]
	}

	return root
}

func Testconnect117() {

}
