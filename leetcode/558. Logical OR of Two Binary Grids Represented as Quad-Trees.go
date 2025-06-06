package leetcode

// Definition for a QuadTree node.
type Node558 struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node558
	TopRight    *Node558
	BottomLeft  *Node558
	BottomRight *Node558
}

func intersect(quadTree1 *Node558, quadTree2 *Node558) *Node558 {
	if quadTree1.IsLeaf {
		if quadTree1.Val {
			return quadTree1
		}
		return quadTree2
	}

	if quadTree2.IsLeaf {
		if quadTree2.Val {
			return quadTree2
		}
		return quadTree1
	}

	tl := intersect(quadTree1.TopLeft, quadTree2.TopLeft)
	tr := intersect(quadTree1.TopRight, quadTree2.TopRight)
	bl := intersect(quadTree1.BottomLeft, quadTree2.BottomLeft)
	br := intersect(quadTree1.BottomRight, quadTree2.BottomRight)

	if tl.IsLeaf && tr.IsLeaf && bl.IsLeaf && br.IsLeaf &&
		tl.Val == tr.Val && tr.Val == bl.Val && bl.Val == br.Val {
		return &Node558{
			Val:         tl.Val,
			IsLeaf:      true,
			TopLeft:     nil,
			TopRight:    nil,
			BottomLeft:  nil,
			BottomRight: nil,
		}
	}

	return &Node558{
		Val:         false, // whatever
		IsLeaf:      false,
		TopLeft:     tl,
		TopRight:    tr,
		BottomLeft:  bl,
		BottomRight: br,
	}
}

func Testintersect() {

}

// - 2 quadtree with the same size n*n
// - return a quadtree of size n*n = quadtree1 OR quadtree2
//  + for non-leaf nodes (isLeaf = false), val can be true or false
//  + for leaf nodes (has no child, or all its children have the same value), val must be set as its children's val

// - OR of 1 leaf
//	- val = 1 => return leaf node val = 1
//	- else return the other node
// - else, OR of every child
//	- if all children are leaf node with the same value => return leaf node, val = children's value
//	- else, it's not leaf node, val can be true or false. Recursively set all children
