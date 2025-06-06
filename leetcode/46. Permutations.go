package leetcode

import "log"

func permute(nums []int) [][]int {
	res := [][]int{}
	permuteBacktrack([]int{}, nums, len(nums), &res)
	return res
}

func permuteBacktrack(state []int, choices []int, n int, res *[][]int) {
	// 1. record if it's the solution
	if len(state) == n {
		item := make([]int, n)
		copy(item, state)
		*res = append(*res, item)
		return
	}

	for i, choice := range choices {
		// 2. make choice
		state = append(state, choice)

		// 3. backtrack
		innerChoices := append([]int{}, choices[:i]...)
		innerChoices = append(innerChoices, choices[i+1:]...)
		permuteBacktrack(state, innerChoices, n, res)

		// 4. undo choice
		state = state[:len(state)-1]
	}
}

func permuteBacktrack2(state []int, choices []int, selected map[int][]byte, res *[][]int) {
	// 1. record if it's the solution
	if len(state) == len(choices) {
		item := make([]int, len(choices))
		copy(item, state)
		*res = append(*res, item)
		return
	}

	for i, choice := range choices {
		if _, ok := selected[i]; ok {
			// prune
			continue
		}

		// 2. make choice
		state = append(state, choice)
		selected[i] = []byte{}

		// 3. backtrack
		permuteBacktrack2(state, choices, selected, res)

		// 4. undo choice
		state = state[:len(state)-1]
		delete(selected, i)
	}
}

func Testpermute() {

	type Case struct {
		numbers []int

		expected [][]int
	}

	cases := []Case{
		{
			numbers:  []int{5, 4, 6, 2},
			expected: [][]int{},
		},
		{
			numbers:  []int{1, 2, 3},
			expected: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}

	for _, testcase := range cases {
		res := permute(testcase.numbers)

		if len(res) != len(testcase.expected) {
			log.Panicf("FAILED. %+v Expected %+v but got %+v", testcase.numbers, testcase.expected, res)
		}
		for i := 0; i < len(testcase.expected); i++ {
			if len(res[i]) != len(testcase.expected[i]) {
				log.Panicf("FAILED. %+v Expected %+v but got %+v", testcase.numbers, testcase.expected, res)
			}

			for j, item := range res[i] {
				if testcase.expected[i][j] != item {
					log.Panicf("FAILED. %+v Expected %+v but got %+v", testcase.numbers, testcase.expected, res)
				}

			}
		}

		log.Printf("MATCHED. %+v result: %+v, expected: %+v\n\n", testcase.numbers, res, testcase.expected)
	}
}
