package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	root *TreeNode
}

// BFSInsert inserts a value into the binary tree using BFS
func (t *Tree) BFSInsert(val int) {
	if t.root == nil {
		t.root = &TreeNode{Val: val}
		return
	}

	queue := []*TreeNode{t.root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left == nil {
			node.Left = &TreeNode{Val: val}
			return
		} else {
			queue = append(queue, node.Left)
		}

		if node.Right == nil {
			node.Right = &TreeNode{Val: val}
			return
		} else {
			queue = append(queue, node.Right)
		}
	}
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	//最后一个元素，lDpeth=0, rDepth = 0,所以长度为rDepth(0)+1
	//倒数第二个是lDpeth=1,rDepth=1,所以长度是rDepth(1)+1
	lDepth := maxDepth(root.Left)
	rDepth := maxDepth(root.Right)

	if lDepth > rDepth {
		return lDepth + 1
	}

	return rDepth + 1
}

func main() {
	tree := &Tree{}
	tree.root = &TreeNode{Val: 3}
	for _, v := range []int{9, 20, -1, -1, 16, 7, 8} {
		tree.BFSInsert(v)
	}
	fmt.Println(maxDepth(tree.root))
}
