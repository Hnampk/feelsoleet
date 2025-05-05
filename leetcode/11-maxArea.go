package leetcode

import (
	"log"
)

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	max := 0

	for left < right {
		area := 0

		if height[left] < height[right] {
			area = (right - left) * height[left]
			left++
		} else {
			area = (right - left) * height[right]
			right--
		}

		if area > max {
			max = area
		}
	}

	return max
}

func TestmaxArea() {

	type Case struct {
		height []int

		expected int
	}

	cases := []Case{
		{
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			height:   []int{1, 1},
			expected: 1,
		},
	}

	for _, testcase := range cases {
		res := maxArea(testcase.height)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v, Expected %+v but got %+v", testcase.height, testcase.expected, res)
		}

		log.Printf("MATCHED. %+v, result: %+v, expected: %+v\n\n", testcase.height, res, testcase.expected)
	}
}
