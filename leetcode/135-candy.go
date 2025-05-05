package leetcode

import "log"

func candy(ratings []int) int {

	return -1
}

func Testcandy() {
	type Case struct {
		nums     []int
		expected int
	}
	cases := []Case{
		{
			nums:     []int{1, 0, 2},
			expected: 5,
		},
		{
			nums:     []int{1, 2, 2},
			expected: 4,
		},
		{
			nums:     []int{1, 0, 2, 3, 1, 4, 1, 1},
			expected: 13,
		},
		{
			nums:     []int{1, 3, 2, 2, 1},
			expected: 7,
		},
	}

	for _, testcase := range cases {

		res := candy(testcase.nums)
		if res != testcase.expected {
			log.Panicf("FAILED. %+v Expected %d but got %d", testcase.nums, testcase.expected, res)
		}

		log.Printf("MATCHED, result: %d, expected: %+v\n\n", res, testcase.expected)
	}
}
