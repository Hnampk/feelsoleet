package leetcode

import "log"

var passed [][]byte

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}

	passed = [][]byte{}
	for range len(board) {
		passed = append(passed, make([]byte, len(board[0])))
	}

	rowCount := len(board)
	colCount := len(board[0])
	var backtrack func(charIndex int, choices [][2]int) bool
	backtrack = func(charIndex int, choices [][2]int) bool {
		// consider if it's the solution
		if charIndex == len(word) {
			return true
		}

		for _, choice := range choices {
			row, col := choice[0], choice[1]
			if word[charIndex] == board[row][col] {
				passed[row][col] = 1
				if backtrack(charIndex+1, nextCheckBuilder(row, col, rowCount, colCount)) {
					return true
				}
				passed[row][col] = 0
			}
		}

		return false
	}

	for r, cols := range board {
		for c, letter := range cols {
			if letter == word[0] {
				passed[r][c] = 1
				if backtrack(1, nextCheckBuilder(r, c, rowCount, colCount)) {
					return true
				}
				passed[r][c] = 0
			}
		}
	}

	return false
}

func nextCheckBuilder(row, col, rowCount, colCount int) [][2]int {
	nextCheck := [][2]int{}
	if row-1 >= 0 && passed[row-1][col] != 1 {
		nextCheck = append(nextCheck, [2]int{row - 1, col})
	}
	if row+1 < rowCount && passed[row+1][col] != 1 {
		nextCheck = append(nextCheck, [2]int{row + 1, col})
	}
	if col-1 >= 0 && passed[row][col-1] != 1 {
		nextCheck = append(nextCheck, [2]int{row, col - 1})
	}
	if col+1 < colCount && passed[row][col+1] != 1 {
		nextCheck = append(nextCheck, [2]int{row, col + 1})
	}
	return nextCheck
}

// find word in a board. Letters are connected vertically or horizontally
//  - find the starting character, then keep backtracking with the index of a character in board, compare with its neighbors with the next letter

func Testexist() {

	type Case struct {
		board [][]byte
		word  string

		expected bool
	}

	cases := []Case{
		{
			board:    [][]byte{{'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a'}, {'a', 'a', 'a', 'a'}},
			word:     "aaaaaaaaaaaaa",
			expected: false,
		},
		{
			board:    [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:     "ABCB",
			expected: false,
		},
		{
			board:    [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:     "ABA",
			expected: false,
		},
		{
			board:    [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:     "ABCCED",
			expected: true,
		},
		{
			board:    [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:     "SEE",
			expected: true,
		},
	}

	for _, testcase := range cases {
		res := exist(testcase.board, testcase.word)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.board, testcase.word, testcase.expected, res)
		}

		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.board, testcase.word, res, testcase.expected)
	}
}
