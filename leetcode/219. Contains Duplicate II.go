package leetcode

import "log"

func containsNearbyDuplicate(nums []int, k int) bool {
	// implement sliding window
	// loop through the arrays, compare every element with its next k elm

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j <= i+k; j++ {
			if j == len(nums) {
				break
			}

			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

func containsNearbyDuplicate2(nums []int, k int) bool {
	// loop through nums, check if the number at index j exists in hashmap or not
	//	if number at index j exists at index i, calculate delta = abs(j - i).
	//		if delta <= k then return true
	//		else, replace the index at hashmap = j
	//	esle, insert index j into hash map with key = number

	existsIndex := make(map[int]int)

	for j, num := range nums {
		if i, existed := existsIndex[num]; existed {
			if j-i >= -k && j-i <= k {
				return true
			}
		}
		existsIndex[num] = j
	}

	return false
}

func TestcontainsNearbyDuplicate() {
	type Case struct {
		nums []int
		k    int

		expected bool
	}

	cases := []Case{
		{
			nums:     []int{1, 2, 3, 1},
			k:        3,
			expected: true,
		},
		{
			nums:     []int{1, 0, 1, 1},
			k:        1,
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 1, 2, 3},
			k:        2,
			expected: false,
		},
	}

	for _, testcase := range cases {
		res := containsNearbyDuplicate(testcase.nums, testcase.k)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.nums, testcase.k, testcase.expected, res)
		}

		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.nums, testcase.k, res, testcase.expected)
	}

}
