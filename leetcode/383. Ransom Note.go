package leetcode

import "log"

func canConstruct2(ransomNote string, magazine string) bool {
	mMap := make(map[byte]int)
	rMap := make(map[byte]int)

	for i := 0; i < len(magazine); i++ {
		mMap[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		rMap[ransomNote[i]]++
	}

	for char, _ := range rMap {
		if rMap[char] > mMap[char] {
			return false
		}
	}

	return true
}
func canConstruct(ransomNote string, magazine string) bool {
	charMap := make(map[byte]int)

	for i := 0; i < len(magazine); i++ {
		charMap[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		if charMap[ransomNote[i]] == 0 {
			return false
		}
		charMap[ransomNote[i]]--
	}

	return true
}

func TestcanConstruct() {
	type Case struct {
		ransomNote string
		magazine   string

		expected bool
	}

	cases := []Case{
		{
			ransomNote: "a",
			magazine:   "b",
			expected:   false,
		},
		{
			ransomNote: "aa",
			magazine:   "ab",
			expected:   false,
		},
		{
			ransomNote: "aa",
			magazine:   "aab",
			expected:   true,
		},
		{
			ransomNote: "aa",
			magazine:   "",
			expected:   false,
		},
		{
			ransomNote: "",
			magazine:   "aa",
			expected:   true,
		},
	}

	for _, testcase := range cases {
		res := canConstruct(testcase.ransomNote, testcase.magazine)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.ransomNote, testcase.magazine, testcase.expected, res)
		}

		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.ransomNote, testcase.magazine, res, testcase.expected)
	}
}
