package leetcode

import (
	"log"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	// split s into words by space " "
	// 	loop through splitted words, each word must match to the character of pattern which is the value in a hashmap
	//	if there's a new word, check if the character at the same index does not map to any word yet. Then set a new value to the hash map

	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	mapped := make(map[string]byte)
	existed := make(map[byte][]byte)

	for i := 0; i < len(words); i++ {
		if found, ok := mapped[words[i]]; ok {
			if found != pattern[i] {
				return false
			}
		} else {
			if _, ok := existed[pattern[i]]; ok {
				return false
			}
			mapped[words[i]] = pattern[i]
			existed[pattern[i]] = []byte{}
		}
	}

	return true
}

func TestwordPattern() {
	type Case struct {
		pattern string
		s       string

		expected bool
	}

	cases := []Case{
		{
			pattern:  "abba",
			s:        "dog cat cat dog",
			expected: true,
		},
		{
			pattern:  "aaaa",
			s:        "dog cat cat dog",
			expected: false,
		},
		{
			pattern:  "abbaa",
			s:        "dog cat cat dog bird",
			expected: false,
		},
	}

	for _, testcase := range cases {
		res := wordPattern(testcase.pattern, testcase.s)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.pattern, testcase.s, testcase.expected, res)
		}
		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.pattern, testcase.s, res, testcase.expected)
	}
}
