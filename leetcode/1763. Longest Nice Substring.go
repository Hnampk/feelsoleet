package leetcode

import "log"

func longestNiceSubstring(s string) string {
	if s == "" {
		return ""
	}

	cMap := make(map[rune][]byte)
	for _, c := range s {
		cMap[c] = []byte{}
	}
	// iterate through each char in s
	// 	if there's a "not nice char" => divide the strings into 2 parts, then run recursion on each part

	for i, c := range s {
		target := c - 32
		if c <= 'Z' {
			target = c + 32 // to uppercase
		}

		if _, ok := cMap[target]; !ok {
			left := longestNiceSubstring(s[:i])
			right := longestNiceSubstring(s[i+1:])

			if len(left) >= len(right) {
				return left
			}
			return right
		}
	}

	return s
}

func TestlongestNiceSubstring() {

	type Case struct {
		s        string
		expected string
	}

	cases := []Case{
		{
			s:        "YazaAay",
			expected: "aAa",
		},
		{
			s:        "Bb",
			expected: "Bb",
		},
		{
			s:        "c",
			expected: "",
		},
	}

	for _, testcase := range cases {
		res := longestNiceSubstring(testcase.s)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v Expected %+v but got %+v", testcase.s, testcase.expected, res)
		}

		log.Printf("MATCHED. %+v result: %+v, expected: %+v\n\n", testcase.s, res, testcase.expected)
	}
}
