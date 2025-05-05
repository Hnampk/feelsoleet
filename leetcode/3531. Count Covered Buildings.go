package leetcode

import (
	"log"
)

func countCoveredBuildings(n int, buildings [][]int) int {
	res := 0
	boundariesLeft := [][]int{}
	boundariesRight := [][]int{}
	boundariesTop := [][]int{}
	boundariesBottom := [][]int{}

	for i := 0; i <= n+1; i++ {
		boundariesLeft = append(boundariesLeft, []int{n + 1, i})
		boundariesBottom = append(boundariesBottom, []int{i, n + 1})
		boundariesRight = append(boundariesRight, []int{0, i})
		boundariesTop = append(boundariesTop, []int{i, 0})
	}

	log.Printf("Left: %+v", boundariesLeft)
	log.Printf("Right: %+v", boundariesRight)
	log.Printf("Top: %+v", boundariesTop)
	log.Printf("Bottom: %+v", boundariesBottom)

	for _, building := range buildings {
		x := building[0]
		y := building[1]

		xMostLeft := boundariesLeft[y][0]
		xMostRight := boundariesRight[y][0]
		yMostTop := boundariesTop[x][1]
		yMostBottom := boundariesBottom[x][1]

		if x < xMostLeft {
			boundariesLeft[y] = building
		}
		if x > xMostRight {
			boundariesRight[y] = building
		}

		if y < yMostBottom {
			boundariesBottom[x] = building
		}

		if y > yMostTop {
			boundariesTop[x] = building
		}
	}

	log.Println("=== after ===")
	log.Printf("Left: %+v", boundariesLeft)
	log.Printf("Right: %+v", boundariesRight)
	log.Printf("Top: %+v", boundariesTop)
	log.Printf("Bottom: %+v", boundariesBottom)
	for _, building := range buildings {
		x := building[0]
		y := building[1]
		xMostLeft := boundariesLeft[y][0]
		xMostRight := boundariesRight[y][0]
		yMostTop := boundariesTop[x][1]
		yMostBottom := boundariesBottom[x][1]

		if xMostLeft == 0 || xMostRight == n+1 || yMostBottom == 0 || yMostTop == n+1 {
			// not enough buildings around
			continue
		}

		if xMostLeft == x || xMostRight == x || yMostBottom == y || yMostTop == y {
			// on the boundary
			continue
		}

		res++
	}

	return res
}

func TestcountCoveredBuildings() {
	type Case struct {
		n         int
		buildings [][]int

		expected int
	}

	cases := []Case{
		{
			n:         3,
			buildings: [][]int{{1, 2}, {2, 2}, {3, 2}, {2, 1}, {2, 3}},
			expected:  1,
		},
		{
			n:         3,
			buildings: [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}},
			expected:  0,
		},
		{
			n:         5,
			buildings: [][]int{{1, 3}, {3, 2}, {3, 3}, {3, 5}, {5, 3}},
			expected:  1,
		},
		{
			n:         4,
			buildings: [][]int{{2, 4}, {1, 2}, {2, 1}, {3, 1}, {1, 4}, {2, 2}, {3, 2}, {1, 3}},
			expected:  1,
		},
	}

	for _, testcase := range cases {
		res := countCoveredBuildings(testcase.n, testcase.buildings)

		if res != testcase.expected {
			log.Panicf("FAILED. %+v, %+v, Expected %+v but got %+v", testcase.n, testcase.buildings, testcase.expected, res)
		}

		log.Printf("MATCHED. %+v, %+v, result: %+v, expected: %+v\n\n", testcase.n, testcase.buildings, res, testcase.expected)
	}

}
