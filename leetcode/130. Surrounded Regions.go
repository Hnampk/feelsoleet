package leetcode

import (
	"log"
)

type Vertice130 struct {
	row, col int
}

func newVertices(r, c int) Vertice130 {
	return Vertice130{
		row: r,
		col: c,
	}
}

func solve(board [][]byte) {
	checked := make(map[Vertice130][]byte)
	for i, row := range board {
		for j, cell := range row {
			if _, ok := checked[newVertices(i, j)]; ok {
				continue
			}

			if cell == 'O' {
				checking := make(map[Vertice130][]byte)
				isSurrounded := dfs130(board, i, j, checking)
				if isSurrounded {
					for v := range checking {
						board[v.row][v.col] = 'X'
					}
				} else {
					for v := range checking {
						checked[newVertices(v.row, v.col)] = []byte{}
					}
				}
			}
		}
	}
}

func dfs130(board [][]byte, row, col int, checked map[Vertice130][]byte) bool {
	res := true

	if row < 0 || row >= len(board) ||
		col < 0 || col >= len(board[row]) {
		return false
	}

	if board[row][col] == 'X' || board[row][col] == 'C' {
		return true
	}
	checked[newVertices(row, col)] = []byte{}

	board[row][col] = 'C' // checked

	offsets := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, offset := range offsets {
		newRow, newCol := row+offset[0], col+offset[1]
		if _, ok := checked[newVertices(newRow, newCol)]; ok {
			continue
		}

		status := dfs130(board, newRow, newCol, checked)
		res = res && status
	}

	board[row][col] = 'O'

	return res
}

func Testsolve() {

	type Case struct {
		board [][]byte

		expected [][]byte
	}

	cases := []Case{
		{
			board:    [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}},
			expected: [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}},
		},
		{
			board:    [][]byte{{'X'}},
			expected: [][]byte{{'X'}},
		},
		// {
		// 	board: [][]byte{
		// 		{'O', 'X', 'O', 'O', 'O', 'O', 'O', 'O', 'O'},
		// 		{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'X'},
		// 		{'O', 'X', 'O', 'X', 'O', 'O', 'O', 'O', 'X'},
		// 		{'O', 'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O'},
		// 		{'X', 'O', 'O', 'O', 'O', 'O', 'O', 'O', 'X'},
		// 		{'X', 'X', 'O', 'O', 'X', 'O', 'X', 'O', 'X'},
		// 		{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'O'},
		// 		{'O', 'O', 'O', 'X', 'O', 'O', 'O', 'O', 'O'},
		// 		{'O', 'O', 'O', 'O', 'O', 'X', 'X', 'O', 'O'},
		// 	},
		// 	expected: [][]byte{{'X'}},
		// },
		{
			board: [][]byte{
				{'O', 'X', 'X', 'O', 'X'},
				{'X', 'O', 'O', 'X', 'O'},
				{'X', 'O', 'X', 'O', 'X'},
				{'O', 'X', 'O', 'O', 'O'},
				{'X', 'X', 'O', 'X', 'O'}},
			expected: [][]byte{
				{'O', 'X', 'X', 'O', 'X'},
				{'X', 'X', 'X', 'X', 'O'},
				{'X', 'X', 'X', 'O', 'X'},
				{'O', 'X', 'O', 'O', 'O'},
				{'X', 'X', 'O', 'X', 'O'}},
		},
	}

	for _, testcase := range cases {
		solve(testcase.board)

		if len(testcase.board) != len(testcase.expected) {
			log.Panicf("FAILED. Expected %+v but got %+v", testcase.expected, testcase.board)
		}
		for i := 0; i < len(testcase.expected); i++ {
			for j := 0; j < len(testcase.expected[i]); j++ {
				if testcase.expected[i][j] != testcase.board[i][j] {
					log.Panicf("FAILED. Expected %+v but got %+v", testcase.expected, testcase.board)
				}
			}
		}

		log.Printf("MATCHED. result: %+v, expected: %+v\n\n", testcase.board, testcase.expected)
	}
}
