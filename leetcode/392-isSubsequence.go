package leetcode

import "log"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}

	sIndex := 0
	for i := 0; i < len(t); i++ {
		if sIndex == len(s) {
			break
		}

		if t[i] == s[sIndex] {
			sIndex++
		}
	}

	if sIndex == len(s) {
		return true
	}

	return false
}

func TestisSubsequence() {
	type Case struct {
		s string
		t string

		expected bool
	}

	cases := []Case{
		{
			t:        "ahbgdc",
			s:        "abc",
			expected: true,
		},
		{
			t:        "ahbgdc",
			s:        "axc",
			expected: false,
		},
		{
			t:        "",
			s:        "axc",
			expected: false,
		},
		{
			t:        "axc",
			s:        "",
			expected: true,
		},
	}

	for _, testcase := range cases {
		res := isSubsequence(testcase.s, testcase.t)

		if res != testcase.expected {
			log.Panicf("FAILED. %s - %s, Expected %+v but got %+v", testcase.s, testcase.t, testcase.expected, res)
		}

		log.Printf("MATCHED. %s - %s, result: %+v, expected: %+v\n\n", testcase.s, testcase.t, res, testcase.expected)
	}
}
