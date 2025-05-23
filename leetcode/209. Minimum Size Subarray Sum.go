package leetcode

import "log"

func minSubArrayLen(target int, nums []int) int {
	min := len(nums) + 1

	total := 0
	begin := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return 1
		}

		total += nums[i]
		if total >= target {
			if i-begin+1 < min {
				min = i - begin + 1
			}

			begin = i - min + 3
			total = 0

			if begin > i {
				i = begin - 1 // jump
				continue
			}

			for j := begin; j <= i; j++ {
				total += nums[j]
			}
		}
	}

	if min == len(nums)+1 {
		return 0
	}

	total = 0
	begin = len(nums) - 1
	for i := len(nums) - 1; i > len(nums)-min; i-- {
		total += nums[i]
		if total >= target {
			if begin-i+1 < min {
				min = begin - i + 1
			}

			begin = i + min - 3
			total = 0

			if begin < i {
				i = begin + 1 // jump
				continue
			}

			for j := begin; j >= i; j-- {
				total += nums[j]
			}
		}

	}

	return min
}

func TestminSubArrayLen() {

	type Case struct {
		numbers []int
		target  int

		expected int
	}

	cases := []Case{
		{
			numbers:  []int{1, 2, 3, 4, 5},
			target:   11,
			expected: 3,
		},
		{
			numbers:  []int{2, 3, 1, 2, 4, 3},
			target:   7,
			expected: 2,
		},
		{
			numbers:  []int{1, 4, 4},
			target:   4,
			expected: 1,
		},
		{
			numbers:  []int{1, 1, 1, 1, 1, 1, 1, 1},
			target:   11,
			expected: 0,
		},
	}

	for _, testcase := range cases {
		res := minSubArrayLen(testcase.target, testcase.numbers)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.numbers, testcase.target, testcase.expected, res)
		}
		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.numbers, testcase.target, res, testcase.expected)
	}
}
