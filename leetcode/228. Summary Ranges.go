package leetcode

import (
	"log"
	"strconv"
)

func summaryRanges(nums []int) []string {
	res := []string{}
	if len(nums) == 0 {
		return res
	}
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}

	beginVal := nums[0]
	prevVal := nums[0]

	for i := 1; i < len(nums); i++ {
		curVal := nums[i]

		if curVal == prevVal+1 {
			prevVal = curVal
			if i < len(nums)-1 {
				continue
			}
		}

		newItem := strconv.Itoa(beginVal)
		if beginVal != prevVal {
			newItem += "->" + strconv.Itoa(prevVal)
		}
		res = append(res, newItem)

		beginVal = curVal
		prevVal = curVal
	}

	if nums[len(nums)-1] != nums[len(nums)-2]+1 {
		res = append(res, strconv.Itoa(nums[len(nums)-1]))
	}

	return res
}

func TestsummaryRanges() {
	type Case struct {
		numbers []int

		expected []string
	}

	cases := []Case{
		{
			numbers:  []int{0, 1, 2, 4, 5, 7},
			expected: []string{"0->2", "4->5", "7"},
		},
		{
			numbers:  []int{0, 2, 3, 4, 6, 8, 9},
			expected: []string{"0", "2->4", "6", "8->9"},
		},
	}

	for _, testcase := range cases {
		res := summaryRanges(testcase.numbers)

		if len(res) != len(testcase.expected) {
			log.Panicf("FAILED. %+v Expected %+v but got %+v", testcase.numbers, testcase.expected, res)
		}
		for i := 0; i < len(testcase.expected); i++ {
			if testcase.expected[i] != res[i] {
				log.Panicf("FAILED. %+v Expected %+v but got %+v", testcase.numbers, testcase.expected, res)
			}
		}

		log.Printf("MATCHED. %+v result: %+v, expected: %+v\n\n", testcase.numbers, res, testcase.expected)
	}
}
