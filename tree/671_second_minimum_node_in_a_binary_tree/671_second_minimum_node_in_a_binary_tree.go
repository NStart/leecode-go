package main

import "fmt"

func main() {
	root := TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 6}
	fmt.Println(findSecondMiniMumValue(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const MAX = 1 << 31

func findSecondMiniMumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}

	q := []*TreeNode{root}
	min := MAX
	for len(q) > 0 {
		count := len(q)
		for i := 0; i < count; i++ {
			cur := q[0]
			q = q[1:]
			if cur.Val != root.Val && cur.Val < min {
				min = cur.Val
			}

			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}

	if min != MAX && min != root.Val {
		return min
	}

	return -1
}
