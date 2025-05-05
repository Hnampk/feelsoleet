package leetcode

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast := head
	for fast != nil && fast.Next != nil {
		head = head.Next
		fast = fast.Next.Next

		if head == fast {
			return true
		}
	}

	return false
}

func TesthasCycle() {

	type Case struct {
		head []int
		pos  int

		expected bool
	}

	cases := []Case{
		{
			head:     []int{3, 2, 0, -4},
			pos:      1,
			expected: true,
		},
		{
			head:     []int{1, 2},
			pos:      0,
			expected: true,
		},
		{
			head:     []int{1},
			pos:      -1,
			expected: false,
		},
	}

	for _, testcase := range cases {
		nodes := []*ListNode{}
		for _, val := range testcase.head {
			nodes = append(nodes, &ListNode{
				Val:  val,
				Next: nil,
			})
		}

		for i := 0; i < len(nodes)-1; i++ {
			nodes[i].Next = nodes[i+1]
		}

		if testcase.pos >= 0 {
			nodes[len(nodes)-1].Next = nodes[testcase.pos]
		}

		res := hasCycle(nodes[0])

		if res != testcase.expected {
			log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.head, testcase.pos, testcase.expected, res)
		}
		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.head, testcase.pos, res, testcase.expected)
	}
}
