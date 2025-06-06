package leetcode

import "log"

func sample(numbers []int, target int) []int {

	return []int{}
}

func Testsample() {
	type Case struct {
		numbers []int
		target  int

		expected []int
	}

	cases := []Case{
		{
			numbers:  []int{2, 7, 11, 15},
			target:   9,
			expected: []int{1, 2},
		},
	}

	for _, testcase := range cases {
		res := sample(testcase.numbers, testcase.target)

		// if res != testcase.expected {
		// 	log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.numbers, testcase.target, testcase.expected, res)
		// }
		// if len(res) != len(testcase.expected) {
		// 	log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.numbers, testcase.target, testcase.expected, res)
		// }
		// for i := 0; i < len(testcase.expected); i++ {
		// 	if testcase.expected[i] != res[i] {
		// 		log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.numbers, testcase.target, testcase.expected, res)
		// 	}
		// }

		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.numbers, testcase.target, res, testcase.expected)
	}
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Definition for a QuadTree node.
type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node

	Left  *Node // 117
	Right *Node // 117
	Next  *Node // 117
}
