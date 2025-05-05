package leetcode

import (
	"log"
	"sort"
	"strings"
)

type Key [26]byte

func getKey(s string) Key {
	k := Key{}

	for i := 0; i < len(s); i++ {
		k[s[i]-'a']++
	}

	return k
}

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}

	wordMap := make(map[Key][]string)

	for _, word := range strs {
		wordMap[getKey(word)] = append(wordMap[getKey(word)], word)
	}

	for _, val := range wordMap {
		res = append(res, val)
	}

	return res
}

func groupAnagrams2(strs []string) [][]string {
	wordMap := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		chars := strings.Split(strs[i], "")
		sort.Strings(chars)
		key := strings.Join(chars, "")
		wordMap[key] = append(wordMap[key], strs[i])
	}

	res := [][]string{}
	for _, val := range wordMap {
		res = append(res, val)
	}

	return res
}

func TestgroupAnagrams() {
	type Case struct {
		strs []string

		expected [][]string
	}

	cases := []Case{
		{
			strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},

			expected: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
	}

	for _, testcase := range cases {
		res := groupAnagrams(testcase.strs)

		if len(res) != len(testcase.expected) {
			log.Panicf("FAILED. %+v, Expected %+v but got %+v", testcase.strs, testcase.expected, res)
		}
		for i := 0; i < len(testcase.expected); i++ {
			if len(testcase.expected[i]) != len(res[i]) {
				log.Panicf("FAILED. %+v, Expected %+v but got %+v", testcase.strs, testcase.expected, res)
			}

			for j := 0; j < len(testcase.expected[i]); j++ {
				if testcase.expected[i][j] != res[i][j] {
					log.Panicf("FAILED. %+v, Expected %+v but got %+v", testcase.strs, testcase.expected, res)
				}
			}
		}

		log.Printf("MATCHED. %+v, result: %+v, expected: %+v\n\n", testcase.strs, res, testcase.expected)
	}
}
