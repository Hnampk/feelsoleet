package leetcode

import "log"

// Definition for a QuadTree node.
type Node427 struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node427
	TopRight    *Node427
	BottomLeft  *Node427
	BottomRight *Node427
}

func construct(grid [][]int) *Node427 {
	n := len(grid)
	if n == 0 {
		return nil
	}

	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum += grid[i][j]
		}
	}

	if sum == n*n || sum == 0 {
		return &Node427{
			Val:    grid[0][0] == 1,
			IsLeaf: true,
		}
	}

	// divide into sub grids, until the size = 1
	topLeft := [][]int{}
	topRight := [][]int{}
	bottomLeft := [][]int{}
	bottomRight := [][]int{}
	size := n / 2 // sub grid size
	for i := 0; i < size; i++ {
		topLeft = append(topLeft, grid[i][:size])
		topRight = append(topRight, grid[i][size:])
		bottomLeft = append(bottomLeft, grid[size+i][:size])
		bottomRight = append(bottomRight, grid[size+i][size:])
	}

	return &Node427{
		IsLeaf:      false,
		TopLeft:     construct(topLeft),
		TopRight:    construct(topRight),
		BottomLeft:  construct(bottomLeft),
		BottomRight: construct(bottomRight),
	}
}

func Testconstruct() {
	// grid := [][]int{{1, 1, 2, 2}, {1, 1, 2, 2}, {3, 3, 4, 4}, {3, 3, 4, 4}}
	grid := [][]int{{0, 1}, {1, 0}}
	node := construct(grid)

	log.Printf("node %+v", node)

	if node.TopLeft != nil {
		log.Printf("node.TopLeft %+v", node.TopLeft)
		log.Printf("node.TopRight %+v", node.TopRight)
		log.Printf("node.BottomLeft %+v", node.BottomLeft)
		log.Printf("node.BottomRight %+v", node.BottomRight)

	}

}
