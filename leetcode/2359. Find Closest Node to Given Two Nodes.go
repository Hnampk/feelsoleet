package leetcode

import (
	"log"
)

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	reachable := [2][]int{}

	nodes := []int{node1, node2}

	for i, node := range nodes {
		nodeIndex := node
		metNodes := make(map[int][]byte)

		for nodeIndex > -1 {
			if _, ok := metNodes[nodeIndex]; ok {
				break
			}

			metNodes[nodeIndex] = []byte{}
			reachable[i] = append(reachable[i], nodeIndex)

			nodeIndex = edges[nodeIndex]
		}
	}

	foundIndex := 1000000
	distance := 1000000
	for i, n2 := range reachable[1] {
		for j, n1 := range reachable[0] {
			if n2 == n1 {
				maxDistance := i
				if j > i {
					maxDistance = j
				}

				if distance > maxDistance {
					foundIndex = n2
					distance = maxDistance
				} else if distance == maxDistance {
					if n2 < foundIndex {
						foundIndex = n2
					}
				}
			}
		}
	}

	if foundIndex == 1000000 {
		return -1
	}

	return foundIndex
}

// - node i directs to node edges[i]
// => find a node can be reached by both node1 and node2
//  - maximum distance from node1
//  - minimum distance from node2
//  - return the smallest index if there's multiple answer

// 1. find reachable nodes of node1 and node2 (edges may contain cycles)
// 2. merge 2 lists
// 3. find shortest from node2, with smallest index
// 4. if there's no answer, return -1

func TestclosestMeetingNode() {

	type Case struct {
		edges []int
		node1 int
		node2 int

		expected int
	}

	cases := []Case{
		{
			edges:    []int{2, 2, 3, -1},
			node1:    0,
			node2:    1,
			expected: 2,
		},
		{
			edges:    []int{1, 2, -1},
			node1:    0,
			node2:    2,
			expected: 2,
		},
		{
			edges:    []int{4, 3, 0, 5, 3, -1},
			node1:    4,
			node2:    0,
			expected: 4,
		},
		{
			edges:    []int{5, 3, 1, 0, 2, 4, 5},
			node1:    3,
			node2:    2,
			expected: 3,
		},
	}

	for _, testcase := range cases {
		res := closestMeetingNode(testcase.edges, testcase.node1, testcase.node2)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v Expected %+v but got %+v", testcase.edges, testcase.expected, res)
		}

		log.Printf("MATCHED. %+v result: %+v, expected: %+v\n\n", testcase.edges, res, testcase.expected)
	}
}
