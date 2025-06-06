package leetcode

import (
	"log"
)

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	// // method 1: Brute-force
	// return minPathSumDFSBruteForce(grid, len(grid)-1, len(grid[0])-1)

	// // method 2: Memoized search
	// passed := [][]int{}

	// for range len(grid) {
	// 	r := []int{}
	// 	for range len(grid[0]) {
	// 		r = append(r, -1)
	// 	}
	// 	passed = append(passed, r)
	// }
	// return minPathSumMemoizedSearch(grid, len(grid)-1, len(grid[0])-1, passed)

	// method 3: DP
	// return minPathSumDP(grid)

	// method 3+: DP space optimization
	return minPathSumDPSpaceOp(grid)
}

// method 1: Brute-force
func minPathSumDFSBruteForce(grid [][]int, i, j int) int {
	if i == 0 && j == 0 {
		return grid[0][0]
	}

	if i < 0 || j < 0 {
		return 9999999 // a very large value
	}

	return grid[i][j] + minInt(minPathSumDFSBruteForce(grid, i-1, j), minPathSumDFSBruteForce(grid, i, j-1))
}

// method 2: Memoized search
func minPathSumMemoizedSearch(grid [][]int, i, j int, passed [][]int) int {
	if i == 0 && j == 0 {
		return grid[0][0]
	}

	if i < 0 || j < 0 {
		return 9999999 // a very large value
	}

	if passed[i][j] > -1 {
		return passed[i][j]
	}

	minPath := grid[i][j] + minInt(minPathSumMemoizedSearch(grid, i-1, j, passed), minPathSumMemoizedSearch(grid, i, j-1, passed))
	passed[i][j] = minPath

	return minPath
}

// method 3: DP
func minPathSumDP(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)

	// init boundaries
	dp[0] = make([]int, n)
	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}
	for j := 1; j < n; j++ {
		dp[0][j] = grid[0][j] + dp[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = grid[i][j] + minInt(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[m-1][n-1]
}

// method 3+: DP with space optimization
func minPathSumDPSpaceOp(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	// init boundaries
	dp := make([]int, n)
	dp[0] = grid[0][0]
	for j := 1; j < n; j++ {
		dp[j] = grid[0][j] + dp[j-1]
	}

	for i := 1; i < m; i++ {
		dp[0] += grid[i][0]
		for j := 1; j < n; j++ {
			dp[j] = grid[i][j] + minInt(dp[j], dp[j-1])
		}
	}

	return dp[n-1]
}

func minInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func TestminPathSum() {

	type Case struct {
		grid [][]int

		expected int
	}

	cases := []Case{
		{
			grid:     [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			expected: 7,
		},
		{
			grid:     [][]int{{1, 2, 3}, {4, 5, 6}},
			expected: 12,
		},
	}

	for _, testcase := range cases {
		res := minPathSum(testcase.grid)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v, Expected %+v but got %+v", testcase.grid, testcase.expected, res)
		}

		log.Printf("MATCHED. %+v, result: %+v, expected: %+v\n\n", testcase.grid, res, testcase.expected)
	}
}
