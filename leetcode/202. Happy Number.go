package leetcode

import "log"

func isHappy(n int) bool {
	tracking := make(map[int][]byte)

	for n != 1 {
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		n = sum

		if _, ok := tracking[n]; ok {
			return false
		}
		tracking[n] = []byte{}
	}

	return true
}

func TestisHappy() {
	type Case struct {
		n int

		expected bool
	}

	cases := []Case{
		{
			n:        19,
			expected: true,
		},
		{
			n:        2,
			expected: false,
		},
	}

	for _, testcase := range cases {
		res := isHappy(testcase.n)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v, Expected %+v but got %+v", testcase.n, testcase.expected, res)
		}

		log.Printf("MATCHED. %+v, result: %+v, expected: %+v\n\n", testcase.n, res, testcase.expected)
	}

}
