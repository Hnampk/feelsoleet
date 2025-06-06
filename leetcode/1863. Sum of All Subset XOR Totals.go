package leetcode

import "log"

func subsetXORSum(nums []int) int {
	total := 0
	// subsetXORSumBacktrack([]int{}, nums, &total)
	subsetXORSumBacktrack2([]int{}, nums, 0, &total)
	return total
}

func subsetXORSumBacktrack(state, choices []int, total *int) {
	// 1. if it's a solution
	xor := 0
	for _, num := range state {
		xor ^= num
	}
	*total += xor

	for i, num := range choices {
		// 2. make choice
		state = append(state, num)

		// 3. backtrack
		subsetXORSumBacktrack(state, choices[i+1:], total)

		// 4. undo choice
		state = state[:len(state)-1]
	}
}

func subsetXORSumBacktrack2(state, choices []int, start int, total *int) {
	// 1. if it's a solution
	xor := 0
	for _, num := range state {
		xor ^= num
	}
	*total += xor
	for i := start; i < len(choices); i++ {
		num := choices[i]
		// 2. make choice
		state = append(state, num)

		// 3. backtrack
		subsetXORSumBacktrack2(state, choices, i+1, total)

		// 4. undo choice
		state = state[:len(state)-1]
	}
}

// [5,1,6]
// [5]				[1]				[6]
// [5,1] [5,6]		[1,6]
// [5,1,6]

func TestsubsetXORSum() {

	type Case struct {
		numbers []int

		expected int
	}

	cases := []Case{
		{
			numbers:  []int{1, 3},
			expected: 6,
		},
		{
			numbers:  []int{5, 1, 6},
			expected: 28,
		},
		{
			numbers:  []int{3, 4, 5, 6, 7, 8},
			expected: 480,
		},
	}

	for _, testcase := range cases {
		res := subsetXORSum(testcase.numbers)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v Expected %+v but got %+v", testcase.numbers, testcase.expected, res)
		}
		log.Printf("MATCHED. %+v result: %+v, expected: %+v\n\n", testcase.numbers, res, testcase.expected)
	}
}
