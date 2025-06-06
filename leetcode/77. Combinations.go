package leetcode

import "log"

func combine(n int, k int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	res := [][]int{}

	var backtrack func(start int, state []int)
	backtrack = func(start int, state []int) {
		if len(state) == k {
			comb := make([]int, k)
			copy(comb, state)
			res = append(res, comb)
			return
		}

		for i := start; i < n+1; i++ {
			// 2. make a choice
			state = append(state, i)

			// 3. call backtrack
			backtrack(i+1, state)

			// 4. undo choice
			state = state[:len(state)-1]
		}
	}

	// combineBacktrack(n, k, 1, []int{}, &res)
	backtrack(1, []int{})
	return res
}

func combineBacktrack(n, k int, start int, state []int, res *[][]int) {
	// 1. consider it as a solution if it meet the condition
	if len(state) == k {
		comb := make([]int, k)
		copy(comb, state)
		*res = append(*res, comb)
		return
	}

	for i := start; i < n+1; i++ {
		// 2. make a choice
		state = append(state, i)

		// 3. call backtrack
		combineBacktrack(n, k, i+1, state, res)

		// 4. undo choice
		state = state[:len(state)-1]
	}
}

func Testcombine() {
	type Case struct {
		n int
		k int

		expected [][]int
	}

	cases := []Case{
		{
			n:        4,
			k:        2,
			expected: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			n:        1,
			k:        1,
			expected: [][]int{{1}},
		},
		{
			n:        0,
			k:        0,
			expected: [][]int{},
		},
	}

	for _, testcase := range cases {
		res := combine(testcase.n, testcase.k)

		// if res != testcase.expected {
		// 	log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.numbers, testcase.target, testcase.expected, res)
		// }
		if len(res) != len(testcase.expected) {
			log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.n, testcase.k, testcase.expected, res)
		}
		for i := 0; i < len(testcase.expected); i++ {
			for j := 0; j < len(testcase.expected[i]); j++ {
				if testcase.expected[i][j] != res[i][j] {
					log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.n, testcase.k, testcase.expected, res)
				}
			}
		}

		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.n, testcase.k, res, testcase.expected)
	}
}
