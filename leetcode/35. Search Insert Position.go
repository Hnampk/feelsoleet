package leetcode

import "log"

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
			continue
		}
		if nums[mid] > target {
			right = mid - 1
			continue
		}
	}

	return left
}

func TestsearchInsert() {
	type Case struct {
		numbers []int
		target  int

		expected int
	}

	cases := []Case{
		{
			numbers:  []int{1, 3, 5, 6},
			target:   5,
			expected: 2,
		},
		{
			numbers:  []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		{
			numbers:  []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
	}

	for _, testcase := range cases {
		res := searchInsert(testcase.numbers, testcase.target)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.numbers, testcase.target, testcase.expected, res)
		}
		// if len(res) != 2 {
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
