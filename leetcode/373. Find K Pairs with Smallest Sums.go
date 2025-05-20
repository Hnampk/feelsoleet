package leetcode

import (
	"container/heap"
	"log"
)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := &pairheap{}
	heap.Init(h)

	for i := 0; i < len(nums1); i++ {
		heap.Push(h, Pair{
			i:   i,
			j:   0,
			sum: nums1[i] + nums2[0],
		})
		if h.Len() >= k {
			break
		}
	}

	res := [][]int{}
	for h.Len() > 0 && k > 0 {
		pair := heap.Pop(h).(Pair)
		res = append(res, []int{nums1[pair.i], nums2[pair.j]}) // the top item is always the smallest one

		if pair.j+1 < len(nums2) {
			// push new item into heap
			heap.Push(h, Pair{
				i:   pair.i,
				j:   pair.j + 1,
				sum: nums1[pair.i] + nums2[pair.j+1],
			})
		}

		k--
	}

	return res
}

func TestkSmallestPairs() {
	type Case struct {
		nums1 []int
		nums2 []int
		k     int

		expected [][]int
	}

	cases := []Case{
		{
			nums1:    []int{1, 7, 11},
			nums2:    []int{2, 4, 6},
			k:        3,
			expected: [][]int{{1, 2}, {1, 4}, {1, 6}},
		},
		{
			nums1:    []int{1, 1, 2},
			nums2:    []int{1, 2, 3},
			k:        2,
			expected: [][]int{{1, 1}, {1, 1}},
		},
	}

	for _, testcase := range cases {
		res := kSmallestPairs(testcase.nums1, testcase.nums2, testcase.k)

		if len(res) != len(testcase.expected) {
			log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.nums1, testcase.nums2, testcase.expected, res)
		}
		for i := 0; i < len(testcase.expected); i++ {
			if testcase.expected[i][0] != res[i][0] || testcase.expected[i][1] != res[i][1] {
				log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.nums1, testcase.nums2, testcase.expected, res)
			}
		}

		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.nums1, testcase.nums2, res, testcase.expected)
	}
}

type Pair struct {
	i, j int // index in array
	sum  int
}

type pairheap []Pair

func (h *pairheap) Push(elm any) {
	*h = append(*h, elm.(Pair))
}

func (h *pairheap) Pop() any {
	last := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return last
}

func (h *pairheap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *pairheap) Less(i, j int) bool {
	return (*h)[i].sum < (*h)[j].sum // min heap
}

func (h *pairheap) Len() int {
	return len(*h)
}
