package leetcode

import (
	"log"
)

func hammingWeight(n int) int {
	bitSum := 0

	for n != 0 {
		bitSum += n & 1
		n = n >> 1
	}

	return bitSum
}

func TesthammingWeight() {

	type Case struct {
		n        int
		expected int
	}

	cases := []Case{
		{
			n:        11,
			expected: 3,
		},
		{
			n:        128,
			expected: 1,
		},
		{
			n:        2147483645,
			expected: 30,
		},
	}

	for _, testcase := range cases {
		res := hammingWeight(testcase.n)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v Expected %+v but got %+v", testcase.n, testcase.expected, res)
		}

		log.Printf("MATCHED. %+v result: %+v, expected: %+v\n\n", testcase.n, res, testcase.expected)
	}

}
