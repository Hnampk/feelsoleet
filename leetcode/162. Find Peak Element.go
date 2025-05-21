package leetcode

import "log"

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			if nums[i] > nums[i+1] {
				return 0
			}
			continue
		}

		if i == len(nums)-1 {
			if nums[i] > nums[i-1] {
				return len(nums) - 1
			}
			continue
		}

		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}

	return -1
}

// func findPeakElement(nums []int) int {
// 	var n = len(nums)

// 	if n == 1 {
// 		// initial element of array of length 1
// 		return 0
// 	}

// 	// initial element of array
// 	if nums[0] > nums[1] {
// 		return 0
// 	}

// 	// last element of array
// 	if nums[n-1] > nums[n-2] {
// 		return n - 1
// 	}

// 	// if first or last element are not peak element
// 	// then use binary search for finding peak element
// 	left := 1
// 	right := n - 2

// 	for left <= right {
// 		mid := (left + right) / 2
// 		if nums[mid] > nums[mid+1] && nums[mid] > nums[mid-1] {
// 			// peak element found
// 			return mid
// 		}

//			if nums[mid] > nums[mid-1] {
//				// slope is increasing
//				// so peak will be on right
//				left = mid + 1
//			} else {
//				// slope is decreasing
//				// peak will be on left
//				right = mid - 1
//			}
//		}
//		return -1
//	}
func TestfindPeakElement() {

	type Case struct {
		numbers []int

		expected int
	}

	cases := []Case{
		{
			numbers:  []int{1, 2, 1, 3, 5, 6, 4},
			expected: 1,
		},
		{
			numbers:  []int{1, 2, 3, 1},
			expected: 2,
		},
		{
			numbers:  []int{1},
			expected: 0,
		},
		{
			numbers:  []int{2, 1},
			expected: 0,
		},
		{
			numbers:  []int{1, 2},
			expected: 1,
		},
		{
			numbers:  []int{1, 2, 0, 1, 2, 3, 1},
			expected: 1,
		},
	}

	for _, testcase := range cases {
		res := findPeakElement(testcase.numbers)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v Expected %+v but got %+v", testcase.numbers, testcase.expected, res)
		}

		log.Printf("MATCHED. %+v result: %+v, expected: %+v\n\n", testcase.numbers, res, testcase.expected)
	}
}
