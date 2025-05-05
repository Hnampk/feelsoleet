package leetcode

import "log"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	letterCount := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		letterCount[s[i]]++
		letterCount[t[i]]--
	}

	for _, count := range letterCount {
		if count != 0 {
			return false
		}
	}

	return true
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[byte]int)
	tMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
		tMap[t[i]]++
	}

	if len(sMap) != len(tMap) {
		return false
	}

	for letter, sCount := range sMap {
		if tCount, ok := tMap[letter]; !ok {
			return false
		} else if sCount != tCount {
			return false
		}
	}

	return true
}

func TestisAnagram() {

	type Case struct {
		s string
		t string

		expected bool
	}

	cases := []Case{
		{
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			s:        "rat",
			t:        "car",
			expected: false,
		},
	}

	for _, testcase := range cases {
		res := isAnagram(testcase.s, testcase.t)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.s, testcase.t, testcase.expected, res)
		}

		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.s, testcase.t, res, testcase.expected)
	}
}
