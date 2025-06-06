package leetcode

import "log"

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}

	if len(candidates) == 0 {
		return res
	}

	var backtrack func(state []int, remaining, startIndex int)
	backtrack = func(state []int, remaining, startIndex int) {
		// consider if it's a solution
		if remaining == 0 {
			copied := make([]int, len(state))
			copy(copied, state)
			res = append(res, copied)
			return
		}

		for i := startIndex; i < len(candidates); i++ {
			if candidates[i] > remaining {
				// prune
				continue
			}

			// pre-check to save stack space
			duplicate := 0
			for candidates[i]*(duplicate+1) <= remaining {
				duplicate++

				// make a choice
				state = append(state, candidates[i])
				backtrack(state, remaining-candidates[i]*duplicate, i+1)
			}
			state = state[:len(state)-duplicate]
		}
	}

	backtrack([]int{}, target, 0)

	return res
}

func TestcombinationSum() {
	type Case struct {
		numbers []int
		target  int

		expected [][]int
	}

	cases := []Case{
		{
			numbers:  []int{2, 3, 6, 7},
			target:   7,
			expected: [][]int{{2, 2, 3}, {7}},
		},
	}

	for _, testcase := range cases {
		res := combinationSum(testcase.numbers, testcase.target)

		if len(res) != len(testcase.expected) {
			log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.numbers, testcase.target, testcase.expected, res)
		}
		for i := 0; i < len(testcase.expected); i++ {
			if len(res[i]) != len(testcase.expected[i]) {
				log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.numbers, testcase.target, testcase.expected, res)
			}
			for j := 0; j < len(testcase.expected[i]); j++ {
				if testcase.expected[i][j] != res[i][j] {
					log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.numbers, testcase.target, testcase.expected, res)
				}
			}
		}

		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.numbers, testcase.target, res, testcase.expected)
	}
}
