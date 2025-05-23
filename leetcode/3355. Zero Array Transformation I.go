package leetcode

import (
	"log"
)

func isZeroArray(nums []int, queries [][]int) bool {
	counter := make([]int, len(nums)+1)
	for _, query := range queries {
		counter[query[0]]++
		counter[query[1]+1]--
	}

	totalCount := 0
	for i, num := range nums {
		totalCount += counter[i]

		if totalCount < num {
			return false
		}
	}

	return true
}

func isZeroArray1(nums []int, queries [][]int) bool {
	// 1. for each elm in queries, decrease directly from the nums
	//  after iterated through all queries, check all the elements to find none zero element
	//
	// 2. began with a zero array. Loop through elm in queries, increase each element by 1
	//   after iterated through all queries, compare all the elements of nums and the new array to make sure they're match with each other

	res := make([]int, len(nums))

	for _, query := range queries {
		for i := query[0]; i <= query[1]; i++ {
			res[i]++
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > res[i] {
			return false
		}
	}

	return true
}

func TestisZeroArray() {

	type Case struct {
		numbers []int
		queries [][]int

		expected bool
	}

	cases := []Case{
		{
			numbers:  []int{0, 6, 1, 7},
			queries:  [][]int{{2, 3}, {0, 2}, {0, 0}, {1, 3}, {2, 2}, {1, 3}, {0, 3}, {3, 3}, {0, 0}, {2, 3}, {2, 3}, {0, 2}, {3, 3}, {3, 3}},
			expected: false,
		},
		{
			numbers:  []int{1, 0, 1},
			queries:  [][]int{{0, 2}},
			expected: true,
		},
		{
			numbers:  []int{4, 3, 2, 1},
			queries:  [][]int{{0, 2}, {1, 3}},
			expected: false,
		},
		{
			numbers:  []int{1, 2, 3, 4},
			queries:  [][]int{{0, 3}, {2, 3}, {2, 3}, {2, 3}},
			expected: false,
		},
		{
			numbers:  []int{1, 2, 3, 4},
			queries:  [][]int{{0, 3}, {1, 3}, {2, 3}, {1, 3}},
			expected: true,
		},
		{
			numbers:  []int{1, 2, 3, 4},
			queries:  [][]int{{0, 3}, {1, 3}, {2, 3}},
			expected: false,
		},
	}

	for _, testcase := range cases {
		res := isZeroArray(testcase.numbers, testcase.queries)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.numbers, testcase.queries, testcase.expected, res)
		}
		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.numbers, testcase.queries, res, testcase.expected)
	}
}
