package leetcode

import "log"

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	letterMap := make(map[byte]string)
	letterMap['2'] = "abc"
	letterMap['3'] = "def"
	letterMap['4'] = "ghi"
	letterMap['5'] = "jkl"
	letterMap['6'] = "mno"
	letterMap['7'] = "pqrs"
	letterMap['8'] = "tuv"
	letterMap['9'] = "wxyz"

	res := []string{}
	var backtrack func(state string, i int)
	backtrack = func(state string, i int) {
		// consider if it's a solution
		if i == len(digits) {
			res = append(res, state)
			return
		}

		letters := letterMap[digits[i]]
		for _, l := range letters {
			// make a choice
			state += string(l)
			backtrack(state, i+1)
			// redo choice
			state = state[:len(state)-1]
		}
	}

	backtrack("", 0)

	return res
}

func TestletterCombinations() {

	type Case struct {
		digits string

		expected []string
	}

	cases := []Case{
		{
			digits:   "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
	}

	for _, testcase := range cases {
		res := letterCombinations(testcase.digits)

		if len(res) != len(testcase.expected) {
			log.Panicf("FAILED. %+v Expected %+v but got %+v", testcase.digits, testcase.expected, res)
		}
		for i := 0; i < len(testcase.expected); i++ {
			if testcase.expected[i] != res[i] {
				log.Panicf("FAILED. %+v Expected %+v but got %+v", testcase.digits, testcase.expected, res)
			}
		}

		log.Printf("MATCHED. %+v result: %+v, expected: %+v\n\n", testcase.digits, res, testcase.expected)
	}
}
