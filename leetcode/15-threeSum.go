package leetcode

import (
	"log"
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		if nums[i] > 0 {
			// sorted array => sum of all positive numbers always > 0
			break
		}

		target, left, right := -nums[i], i+1, len(nums)-1

		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				res = append(res, []int{nums[left], nums[i], nums[right]})
				left++
				right--

				for left < right && nums[right] == nums[right+1] {
					right--
				}
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			} else if sum > target {
				right--
			} else if sum < target {
				left++
			}
		}
	}

	return res
}

func TestthreeSum() {
	type Case struct {
		nums []int

		expected [][]int
	}

	cases := []Case{
		{
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			nums:     []int{0, 0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			nums:     []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			nums:     []int{2, -3, 0, -2, -5, -5, -4, 1, 2, -2, 2, 0, 2, -4, 5, 5, -10},
			expected: [][]int{{-10, 5, 5}, {-5, 0, 5}, {-4, 2, 2}, {-3, -2, 5}, {-3, 1, 2}, {-2, 0, 2}},
		},
	}

	for _, testcase := range cases {
		res := threeSum(testcase.nums)

		if len(res) != len(testcase.expected) {
			log.Panicf("FAILED. %+v, Expected %+v but got %+v", testcase.nums, testcase.expected, res)
		}
		for i := 0; i < len(testcase.expected); i++ {
			sort.Ints(testcase.expected[i])
			sort.Ints(res[i])

			for j := 0; j < len(testcase.expected[i]); j++ {
				if testcase.expected[i][j] != res[i][j] {
					log.Panicf("FAILED. %+v, Expected %+v but got %+v", testcase.nums, testcase.expected, res)
				}

			}
		}

		log.Printf("MATCHED. %+v, result: %+v, expected: %+v\n\n", testcase.nums, res, testcase.expected)
	}
}
