package leetcode

import "log"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[int]*Node)

	return bfs133(node, visited)
}

func dfs133(node *Node, visited map[int]*Node) *Node {
	if n, ok := visited[node.Val]; ok {
		return n
	}

	n := &Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}
	visited[node.Val] = n

	for _, neighbor := range node.Neighbors {
		n.Neighbors = append(n.Neighbors, dfs133(neighbor, visited))
	}

	return n
}

func bfs133(node *Node, visited map[int]*Node) *Node {
	q := []*Node{}
	q = append(q, node)

	visited[node.Val] = &Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		var n *Node = visited[curr.Val] // always found the node, because it's defined as a neighbor

		for _, neighbor := range curr.Neighbors {
			if found, ok := visited[neighbor.Val]; !ok {
				q = append(q, neighbor)
				newNode := &Node{
					Val:       neighbor.Val,
					Neighbors: []*Node{},
				}

				n.Neighbors = append(n.Neighbors, newNode)
				visited[neighbor.Val] = newNode
			} else {
				n.Neighbors = append(n.Neighbors, found)
			}
		}
	}

	return visited[1]
}

func cloneGraph2(node *Node) *Node {
	// recursively check through neighbor nodes to create nodes.
	//  we need a pointer to the parent nodes to add the edges after creating a new node

	if node == nil {
		return nil
	}

	visited := make(map[int]*Node)
	copiedNode := &Node{
		Val:       node.Val,
		Neighbors: []*Node{},
	}
	visited[node.Val] = copiedNode
	dfs1332(node, copiedNode, visited)
	return copiedNode

}
func dfs1332(node *Node, clone *Node, visited map[int]*Node) {
	for _, neighbor := range node.Neighbors {
		var newNode *Node
		if n, ok := visited[neighbor.Val]; !ok {
			newNode = &Node{
				Val:       neighbor.Val,
				Neighbors: []*Node{},
			}
			visited[neighbor.Val] = newNode
			dfs1332(neighbor, newNode, visited)
		} else {
			newNode = n
		}
		clone.Neighbors = append(clone.Neighbors, newNode)
	}
}

func TestcloneGraph() {
	type Case struct {
		adjList [][]int

		expected [][]int
	}

	cases := []Case{
		{
			adjList:  [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
			expected: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
		},
		{
			adjList:  [][]int{},
			expected: [][]int{},
		},
	}

	for _, testcase := range cases {
		listNode := make(map[int]*Node)
		for val, _ := range testcase.adjList {
			id := val + 1
			node := &Node{
				Val: id,
			}
			listNode[id] = node
		}

		for val, edges := range testcase.adjList {
			id := val + 1
			listNode[id].Neighbors = []*Node{}
			for _, e := range edges {
				listNode[id].Neighbors = append(listNode[id].Neighbors, listNode[e])
			}
		}

		res := cloneGraph(listNode[1])

		visited := make(map[int][]byte)
		print := func(node *Node) {
			log.Printf("==node", node)
			for i := range node.Neighbors {
				visited[node.Val] = []byte{}
				log.Printf("==res", res.Neighbors[i])
			}
		}

		for {
			if res == nil {
				break
			}
			if _, ok := visited[res.Val]; !ok {
				visited[res.Val] = []byte{}
				print(res)

				for _, n := range res.Neighbors {
					if _, ok := visited[n.Val]; !ok {
						res = n
						break
					}
				}
			} else {
				break
			}
		}
		// visited := make(map[int][]byte)

		// for _, list := range testcase.expected {
		// 	visited[res.Val] = []byte{}
		// 	if len(res.Neighbors) != len(list) {
		// 		log.Panicf("")
		// 	}

		// 	for i, n := range res.Neighbors {
		// 		if n.Val != list[i] {
		// 			log.Panicf("")
		// 		}
		// 	}

		// 	for _, n := range res.Neighbors {
		// 		if _, ok := visited[n.Val]; !ok {
		// 			visited[n.Val] = []byte{}
		// 			res = n
		// 			break
		// 		}
		// 	}
		// }

		// log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.numbers, testcase.target, res, testcase.expected)
	}

}
