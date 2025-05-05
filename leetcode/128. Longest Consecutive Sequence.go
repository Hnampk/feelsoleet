package leetcode

import "log"

func longestConsecutive(nums []int) int {
	longestConsecutive := 0
	numMap := make(map[int][]byte)

	for _, num := range nums {
		numMap[num] = []byte{}
	}

	for num := range numMap {
		if _, ok := numMap[num-1]; !ok {
			max := num

			for numMap[max+1] != nil {
				max++
			}

			if max-num+1 > longestConsecutive {
				longestConsecutive = max - num + 1
			}
		}
	}

	return longestConsecutive
}

func longestConsecutive2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	longestConsecutive := 1
	gMin := 0
	gMax := 0
	numMap := make(map[int][]byte)

	for _, num := range nums {
		numMap[num] = []byte{}
	}

	for num := range numMap {
		if num >= gMin && num <= gMax {
			continue
		}

		min := num
		max := num

		for numMap[max+1] != nil {
			max++
		}

		for numMap[min-1] != nil {
			min--
		}

		if max-min+1 > longestConsecutive {
			longestConsecutive = max - min + 1
			gMax = max
			gMin = min
		}
	}

	return longestConsecutive
}

func TestlongestConsecutive() {

	type Case struct {
		nums     []int
		expected int
	}

	cases := []Case{
		{
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4,
		},
		{
			nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expected: 9,
		},
		{
			nums:     []int{1, 0, 1, 2},
			expected: 3,
		},
	}

	for _, testcase := range cases {
		res := longestConsecutive(testcase.nums)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v, Expected %+v but got %+v", testcase.nums, testcase.expected, res)
		}

		log.Printf("MATCHED. %+v, result: %+v, expected: %+v\n\n", testcase.nums, res, testcase.expected)
	}
}
