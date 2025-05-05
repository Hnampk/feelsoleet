package leetcode

import "log"

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func addTwoNumbers(l1 *ListNode2, l2 *ListNode2) *ListNode2 {
	var head *ListNode2
	cur := head

	for l1 != nil || l2 != nil {
		n1 := 0
		n2 := 0

		if l1 != nil {
			n1 = l1.Val
		}
		if l2 != nil {
			n2 = l2.Val
		}
		sum := n1 + n2

		if sum >= 10 {
			if l2.Next == nil {
				l2.Next = &ListNode2{
					Val: 0,
				}
			}
			l2.Next.Val += 1
			sum -= 10
		}
		sumNode := &ListNode2{
			Val: sum,
		}

		if head == nil {
			head = sumNode
			cur = head
		} else {
			cur.Next = sumNode
			cur = cur.Next
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return head
}

func TestaddTwoNumbers() {
	type Case struct {
		i1 []int // input
		i2 []int

		expected []int
	}

	cases := []Case{
		{
			i1:       []int{2, 4, 3},
			i2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
	}

	for _, testcase := range cases {
		nodes1 := []*ListNode2{}
		for _, item := range testcase.i1 {
			nodes1 = append(nodes1, &ListNode2{
				Val: item,
			})
		}

		nodes2 := []*ListNode2{}
		for _, item := range testcase.i2 {
			nodes2 = append(nodes2, &ListNode2{
				Val: item,
			})
		}

		expected := []*ListNode2{}
		for _, item := range testcase.expected {
			expected = append(expected, &ListNode2{
				Val: item,
			})
		}

		for i := 0; i < len(nodes1)-1; i++ {
			nodes1[i].Next = nodes1[i+1]
		}
		for i := 0; i < len(nodes2)-1; i++ {
			nodes2[i].Next = nodes2[i+1]
		}
		for i := 0; i < len(expected)-1; i++ {
			expected[i].Next = expected[i+1]
		}

		res := addTwoNumbers(nodes1[0], nodes2[0])

		resItem := res
		expectedItem := expected[0]
		for {
			if resItem != nil && expectedItem != nil {
				log.Printf("%+v & %+v", resItem, expectedItem)
				if resItem.Val != expectedItem.Val {
					log.Panicf("FAILED 1. %+v, %+v, Expected %+v but got %+v", testcase.i1, testcase.i2, testcase.expected, res)
				}
				resItem = resItem.Next
				expectedItem = expectedItem.Next
			} else {
				if resItem != nil || expectedItem != nil {
					log.Panicf("FAILED 2. %+v, %+v, Expected %+v but got %+v", testcase.i1, testcase.i2, testcase.expected, res)
				} else {
					break
				}
			}
		}

		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.i1, testcase.i2, res, testcase.expected)
	}
}
