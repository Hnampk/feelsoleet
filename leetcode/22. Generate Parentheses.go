package leetcode

import "log"

func generateParenthesis(n int) []string {
	res := []string{}

	var backtrack func(state string, opened, closed int)
	backtrack = func(state string, opened, closed int) {
		// consider if it's a solution
		if len(state) == 2*n {
			res = append(res, state)
			return
		}

		if opened < n {
			backtrack(state+"(", opened+1, closed)
		}

		if closed < opened {
			backtrack(state+")", opened, closed+1)
		}
	}
	backtrack("", 0, 0)

	return res
}

func TestgenerateParenthesis() {
	type Case struct {
		number int

		expected []string
	}

	cases := []Case{
		{
			number:   3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			number:   1,
			expected: []string{"()"},
		},
	}

	for _, testcase := range cases {
		res := generateParenthesis(testcase.number)

		if len(res) != len(testcase.expected) {
			log.Panicf("FAILED. %+v, Expected %+v but got %+v", testcase.number, testcase.expected, res)
		}
		for i := 0; i < len(testcase.expected); i++ {
			if testcase.expected[i] != res[i] {
				log.Panicf("FAILED. %+v, Expected %+v but got %+v", testcase.number, testcase.expected, res)
			}
		}

		log.Printf("MATCHED. %+v, result: %+v, expected: %+v\n\n", testcase.number, res, testcase.expected)
	}
}
