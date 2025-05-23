package contest

import "log"

func smallestIndex(nums []int) int {
	for i, num := range nums {
		sum := 0
		for num >= 10 {
			sum += num % 10
			num /= 10
		}
		sum += num

		if sum == i {
			return i
		}
	}

	return -1
}

func TestsmallestIndex() {

	type Case struct {
		numbers []int

		expected int
	}

	cases := []Case{
		{
			numbers:  []int{1, 3, 2},
			expected: 2,
		},
		{
			numbers:  []int{1, 10, 11},
			expected: 1,
		},
	}

	for _, testcase := range cases {
		res := smallestIndex(testcase.numbers)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v Expected %+v but got %+v", testcase.numbers, testcase.expected, res)
		}

		log.Printf("MATCHED. %+v result: %+v, expected: %+v\n\n", testcase.numbers, res, testcase.expected)
	}
}
