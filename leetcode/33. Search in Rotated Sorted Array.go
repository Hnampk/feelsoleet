package leetcode

import "log"

func search(nums []int, target int) int {
	// mid > right => rotate point on right
	//  target > mid => right
	//  target < mid
	//      target < right => right
	//      target > right => left

	// mid < right => rotate point on left
	//  target < mid => left
	//  target > mid
	//      target < right => right
	//      target > right => left

	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1

	for left < right {
		if target == nums[left] {
			return left
		}
		if target == nums[right] {
			return right
		}

		mid := left + (right-left)/2
		if target == nums[mid] {
			return mid
		}

		if nums[mid] > nums[right] {
			// rotate point is on the right part
			if target > nums[mid] {
				// target is on the right part
				left = mid + 1
				continue
			}

			if target < nums[right] {
				// target is on the right part
				left = mid + 1
				continue
			}
			if target > nums[right] {
				// target is on the left part
				right = mid - 1
				continue
			}

		}

		if nums[mid] < nums[right] {
			// rotate point is on the left part
			if target < nums[mid] {
				// target is on the left part
				right = mid - 1
				continue
			}

			if target < nums[right] {
				// target is on the right part
				left = mid + 1
				continue
			}
			if target > nums[right] {
				// target is on the left part
				right = mid - 1
				continue
			}

		}
	}

	if nums[left] == target {
		return left
	}
	return -1
}

func Testsearch() {
	type Case struct {
		numbers []int
		target  int

		expected int
	}

	cases := []Case{
		{
			numbers:  []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: 4,
		},
		{
			numbers:  []int{4, 5, 6, 7, 0, 1, 2},
			target:   3,
			expected: -1,
		},
		{
			numbers:  []int{1},
			target:   0,
			expected: -1,
		},
		{
			numbers:  []int{1, 3},
			target:   3,
			expected: 1,
		},
	}

	for _, testcase := range cases {
		res := search(testcase.numbers, testcase.target)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.numbers, testcase.target, testcase.expected, res)
		}

		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.numbers, testcase.target, res, testcase.expected)
	}
}
