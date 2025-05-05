package leetcode

import (
	"log"
)

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		leftChar := getAlphanumericByte(s[left])
		if leftChar == 0 {
			left++
			continue
		}

		rightChar := getAlphanumericByte(s[right])
		if rightChar == 0 {
			right--
			continue
		}

		if leftChar != rightChar {
			return false
		}

		left++
		right--
	}

	return true
}

func getAlphanumericByte(c byte) byte {
	if c >= 65 && c <= 90 {
		return c + 32 // to lower case
	} else if (c >= 97 && c <= 122) || (c >= 48 && c <= 57) {
		return c
	}

	return 0
}

func TestisPalindrome() {
	type Case struct {
		s string

		expected bool
	}

	cases := []Case{
		{
			s:        "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			s:        "race a car",
			expected: false,
		},
		{
			s:        "race  car",
			expected: true,
		},
		{
			s:        " ",
			expected: true,
		},
	}

	for _, testcase := range cases {
		res := isPalindrome(testcase.s)

		if res != testcase.expected {
			log.Panicf("FAILED. %s  , Expected %+v but got %+v", testcase.s, testcase.expected, res)
		}

		log.Printf("MATCHED. %s  , result: %+v, expected: %+v\n\n", testcase.s, res, testcase.expected)
	}
}
