package leetcode

import (
	"log"
	"strconv"
)

func evalRPN(tokens []string) int {
	// implement using stack. Put every number into the stack,
	//	when meets an operation token, pop out the numbers to calculate, then push back the value into the stack.
	//	The result will be the last value exists in the stack
	stacksValue := []int{}

	for _, token := range tokens {
		switch token {
		case "+":
			n := len(stacksValue)
			newValue := stacksValue[n-2] + stacksValue[n-1]
			stacksValue = append(stacksValue[:n-2], newValue)
		case "-":
			n := len(stacksValue)
			newValue := stacksValue[n-2] - stacksValue[n-1]
			stacksValue = append(stacksValue[:n-2], newValue)
		case "*":
			n := len(stacksValue)
			newValue := stacksValue[n-2] * stacksValue[n-1]
			stacksValue = append(stacksValue[:n-2], newValue)
		case "/":
			n := len(stacksValue)
			newValue := stacksValue[n-2] / stacksValue[n-1]
			stacksValue = append(stacksValue[:n-2], newValue)

		default:
			intValue, _ := strconv.Atoi(token)
			stacksValue = append(stacksValue, intValue)
		}
	}

	return stacksValue[0]
}

func TestevalRPN() {

	type Case struct {
		tokens []string

		expected int
	}

	cases := []Case{
		{
			tokens:   []string{"2", "1", "+", "3", "*"},
			expected: 9,
		},
		{
			tokens:   []string{"4", "13", "5", "/", "+"},
			expected: 6,
		},
		{
			tokens:   []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			expected: 22,
		},
	}

	for _, testcase := range cases {
		res := evalRPN(testcase.tokens)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v Expected %+v but got %+v", testcase.tokens, testcase.expected, res)
		}
		log.Printf("MATCHED. %+v, result: %+v, expected: %+v\n\n", testcase.tokens, res, testcase.expected)
	}
}
