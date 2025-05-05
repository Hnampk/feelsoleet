package leetcode

import "log"

func strStr(haystack string, needle string) int {
	if len(haystack) == 1 && haystack[0] == needle[0] {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] != needle[0] {
			continue
		}

		if len(needle) == 1 {
			return i
		}

		for j := 1; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}

			if j == len(needle)-1 {
				return i
			}
		}
	}

	return -1
}

func TeststrStr() {
	type Case struct {
		haystack string
		needle   string

		expected int
	}

	cases := []Case{
		{
			haystack: "sadbutsad",
			needle:   "sad",
			expected: 0,
		},
		{
			haystack: "leetcode",
			needle:   "leeto",
			expected: -1,
		},
		{
			haystack: "abcd",
			needle:   "bc",
			expected: 1,
		},
		{
			haystack: "abcdad",
			needle:   "ad",
			expected: 4,
		},
		{
			haystack: "a",
			needle:   "a",
			expected: 0,
		},
		{
			haystack: "ab",
			needle:   "ab",
			expected: 0,
		},
		{
			haystack: "abc",
			needle:   "c",
			expected: 2,
		},
		{
			haystack: "mississippi",
			needle:   "issip",
			expected: 4,
		},
	}

	for _, testcase := range cases {
		res := strStr(testcase.haystack, testcase.needle)

		if res != testcase.expected {
			log.Panicf("FAILED. %s - %s, Expected %+v but got %+v", testcase.haystack, testcase.needle, testcase.expected, res)
		}

		log.Printf("MATCHED. %s - %s, result: %+v, expected: %+v\n\n", testcase.haystack, testcase.needle, res, testcase.expected)
	}
}
