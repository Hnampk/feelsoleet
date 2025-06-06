package leetcode

import (
	"log"
	"strings"
)

func longestSubstring(s string, k int) int {
	frequencies := ['z' - 'a' + 1]int{}

	for _, letter := range s {
		letterIndex := letter - 'a'
		frequencies[letterIndex]++
	}

	subs := []string{s}
	for index, freq := range frequencies {
		if freq > 0 && freq < k {
			newSubs := []string{}
			for _, sub := range subs {
				newSubs = append(newSubs, strings.Split(sub, string('a'+index))...)
			}
			subs = newSubs
		}
	}

	longest := 0
	for _, sub := range subs {
		if len(sub) < k {
			continue
		}

		if isValidSubstring(sub, k) {
			if len(sub) > longest {
				longest = len(sub)
			}
		} else {
			subLongest := longestSubstring(sub, k)
			if subLongest > longest {
				longest = subLongest
			}
		}
	}

	return longest
}

func isValidSubstring(s string, k int) bool {
	if len(s) < k {
		return false
	}

	frequency := ['z' - 'a']int{}
	for _, letter := range s {
		letterIndex := letter - 'a'
		frequency[letterIndex]++
	}

	for _, count := range frequency {
		if count > 0 && count < k {
			return false
		}
	}

	return true
}

// - string s, return the longest sub string of s, with all characters appear >= k times

// 1. - count frequency of every char
// - loop through letter of s, check the frequency and find the longest sub string

// 2. observe on begin indexes of letters has frequency >= k

func TestlongestSubstring() {

	type Case struct {
		s string
		k int

		expected int
	}

	cases := []Case{
		{
			s:        "aaabb",
			k:        3,
			expected: 3,
		},
		{
			s:        "ababbc",
			k:        2,
			expected: 5,
		},
		{
			s:        "ababacb",
			k:        3,
			expected: 0,
		},
	}

	for _, testcase := range cases {
		res := longestSubstring(testcase.s, testcase.k)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.s, testcase.k, testcase.expected, res)
		}

		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.s, testcase.k, res, testcase.expected)
	}
}
