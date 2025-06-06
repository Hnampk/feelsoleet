package leetcode

import (
	"log"
)

func solveNQueens(n int) [][]string {
	res := [][]string{}

	var backtrack func(rowIndex int, state []string, choices [][]int)
	backtrack = func(rowIndex int, state []string, choices [][]int) {
		// 1. consider if it's a solution
		if len(state) == n {
			r := make([]string, n)
			copy(r, state)
			res = append(res, r)
			return
		}

		for row := rowIndex; row < n; row++ {
			cols := choices[row]

			// prune
			if len(cols) == 0 {
				// there is no posible solution
				break
			}

			for _, col := range cols {
				curRow := ""
				// 2. make a choice
				for i := range n {
					if i == col {
						curRow += "Q"
						continue
					}
					curRow += "."
				}
				state = append(state, curRow)

				// 3. backtrack
				backtrack(row+1, state, buildAvailCols(choices, [2]int{row, col}))

				// 4. undo choice
				state = state[:len(state)-1]
			}
		}
	}

	cheers := [][]int{}
	for range n {
		col := make([]int, n)
		for i := range n {
			col[i] = i
		}
		cheers = append(cheers, col)
	}
	backtrack(0, []string{}, cheers)

	return res
}

func buildAvailCols(curState [][]int, pickedIndex [2]int) [][]int {
	pickedRow, pickedCol := pickedIndex[0], pickedIndex[1]
	newState := make([][]int, len(curState))

	for row, cols := range curState {
		if row == pickedRow {
			newState[row] = []int{}
			continue
		}

		if len(cols) == 0 {
			newState[row] = []int{}
			continue
		}

		newCols := []int{}
		for _, col := range cols {
			if col == pickedCol {
				continue
			}
			if row-col == pickedRow-pickedCol {
				continue
			}
			if row+col == pickedRow+pickedCol {
				continue
			}
			newCols = append(newCols, col)
		}
		newState[row] = newCols
	}

	return newState
}

func TestsolveNQueens() {
	type Case struct {
		n int

		expected [][]string
	}

	cases := []Case{
		{
			n:        4,
			expected: [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}},
		},
		{
			n:        1,
			expected: [][]string{{"Q"}},
		},
	}

	for _, testcase := range cases {
		res := solveNQueens(testcase.n)

		if len(res) != len(testcase.expected) {
			log.Panicf("FAILED. %+v Expected %+v but got %+v", testcase.n, testcase.expected, res)
		}
		for i := 0; i < len(testcase.expected); i++ {
			if len(res[i]) != len(testcase.expected[i]) {
				log.Panicf("FAILED. %+v Expected %+v but got %+v", testcase.n, testcase.expected, res)
			}
			for j := 0; j < len(testcase.expected[i]); j++ {
				if testcase.expected[i][j] != res[i][j] {
					log.Panicf("FAILED. %+v Expected %+v but got %+v", testcase.n, testcase.expected, res)
				}

			}
		}

		log.Printf("MATCHED. %+v result: %+v, expected: %+v\n\n", testcase.n, res, testcase.expected)
	}
}
