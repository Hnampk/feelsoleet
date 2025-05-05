package leetcode

import (
	"log"
	"strconv"
)

func calculate(s string) int {
	startIndex := 0
	if s[0] == '(' {
		startIndex = 1
	}
	res, _ := subCalculate(s, startIndex)

	return res
}

// subCalculate is used for calculate the expression inside a bucket
func subCalculate(s string, startIndex int) (res, endIndex int) {
	stack := []int{}

	for i := startIndex; i < len(s); i++ {
		token := s[i]

		switch token {
		case '+', '-':
			n2 := 0
			if s[i+1] == '(' {
				res, endIndex := subCalculate(s, i+1)
				n2 = res
				i = endIndex
			} else {
				startDigit := i + 1
				endDigit := i + 1
				for j := i + 1; j < len(s); j++ {
					if s[j] == '+' || s[j] == '-' || s[j] == ')' {
						endDigit = j
						break
					}
				}
				n2, _ = strconv.Atoi(s[startDigit:endDigit])
				i = endDigit
			}

			n1 := stack[len(stack)-1]

			if token == '+' {
				stack = append(stack[:len(stack)-1], n1+n2)
			}
			if token == '-' {
				stack = append(stack[:len(stack)-1], n1-n2)
			}

		case ')':
			return stack[0], i

		default:
			stack = append(stack, int(token))
		}

	}

	return stack[0], len(s)
}

func Testcalculate() {

	type Case struct {
		s string

		expected int
	}

	cases := []Case{
		{
			s:        "",
			expected: 1,
		},
	}

	for _, testcase := range cases {
		res := calculate(testcase.s)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v, Expected %+v but got %+v", testcase.s, testcase.expected, res)
		}

		log.Printf("MATCHED. %+v, result: %+v, expected: %+v\n\n", testcase.s, res, testcase.expected)
	}
}
