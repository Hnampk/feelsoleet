package leetcode

import (
	"log"
)

func reverseBits(num uint32) uint32 {
	var res uint32 = 0
	for i := 0; i < 32; i++ {
		res = res << 1   // Shift res left
		res += (num & 1) // Add the least significant bit of num to res
		num = num >> 1   // Shift num to the right
	}
	return res
}

func TestreverseBits() {

	type Case struct {
		num      uint32
		expected uint32
	}

	cases := []Case{
		{
			num:      2,
			expected: 2147483645,
		},
	}

	for _, testcase := range cases {
		res := reverseBits(testcase.num)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v Expected %+v but got %+v", testcase.num, testcase.expected, res)
		}
		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.num, res, testcase.expected)
	}
}
