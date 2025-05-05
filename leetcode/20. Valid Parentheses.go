package leetcode

import (
	"log"
)

func isValid(s string) bool {
	// mostOpened []rune, appending open brackets
	//  when found a close brackets, check if it pairs with the most opened one
	//      if not => false
	//      else pop from the most opened backet
	//  finally, if there's any remaining bracket => false

	mostOpened := []rune{}

	bracketMap := map[rune]rune{
		']': '[',
		')': '(',
		'}': '{',
	}

	for _, char := range s {
		switch char {
		case '[', '{', '(':
			mostOpened = append(mostOpened, char)
		case ']', '}', ')':
			if len(mostOpened) == 0 {
				return false
			}

			if mostOpened[len(mostOpened)-1] != bracketMap[char] {
				return false
			}

			mostOpened = mostOpened[:len(mostOpened)-1]
		default:
			return false
		}
	}

	return len(mostOpened) == 0
}

func TestisValid() {
	type Case struct {
		s string

		expected bool
	}

	cases := []Case{
		{
			s:        "()",
			expected: true,
		},
		{
			s:        "()[]{}",
			expected: true,
		},
		{
			s:        "([])",
			expected: true,
		},
		{
			s:        "([",
			expected: false,
		},
		{
			s:        "(]",
			expected: false,
		},
		{
			s:        "([)]",
			expected: false,
		},
	}

	for _, testcase := range cases {
		res := isValid(testcase.s)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v, Expected %+v but got %+v", testcase.s, testcase.expected, res)
		}

		log.Printf("MATCHED. %+v, result: %+v, expected: %+v\n\n", testcase.s, res, testcase.expected)
	}

}
