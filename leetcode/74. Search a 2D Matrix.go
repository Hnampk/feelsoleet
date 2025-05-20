package leetcode

import "log"

func searchMatrix(matrix [][]int, target int) bool {
	top, bottom := 0, len(matrix)-1

	// vertical search
	row := 0
	for top <= bottom {
		mid := top + (bottom-top)/2
		if matrix[mid][0] == target {
			return true
		}
		if matrix[mid][0] < target {
			top = mid + 1
			continue
		}
		if matrix[mid][0] > target {
			bottom = mid - 1
			continue
		}
	}
	if top == 0 {
		// target is less than the first value
		return false
	}

	// horizontal search
	row = top - 1
	left, right := 0, len(matrix[row])-1
	for left <= right {
		mid := left + (right-left)/2
		if matrix[row][mid] == target {
			return true
		}
		if matrix[row][mid] < target {
			left = mid + 1
			continue
		}
		if matrix[row][mid] > target {
			right = mid - 1
			continue
		}
	}

	return false
}

func TestsearchMatrix() {
	type Case struct {
		matrix [][]int
		target int

		expected bool
	}

	cases := []Case{
		{
			matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   3,
			expected: true,
		},
		{
			matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   11,
			expected: true,
		},
		{
			matrix:   [][]int{{1}},
			target:   0,
			expected: false,
		},
		{
			matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			target:   13,
			expected: false,
		},
	}

	for _, testcase := range cases {
		res := searchMatrix(testcase.matrix, testcase.target)
		if res != testcase.expected {
			log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.matrix, testcase.target, testcase.expected, res)
		}

		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.matrix, testcase.target, res, testcase.expected)
	}
}
