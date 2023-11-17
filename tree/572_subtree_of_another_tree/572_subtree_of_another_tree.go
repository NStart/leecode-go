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

func main() {
	s := &TreeNode{Val: 3}
	s.Left = &TreeNode{Val: 4}
	s.Left.Left = &TreeNode{Val: 1}
	s.Left.Right = &TreeNode{Val: 2}
	//s.Left.Right.Left = &TreeNode{Val: 0} //false

	t := &TreeNode{Val: 4}
	t.Left = &TreeNode{Val: 1}
	t.Right = &TreeNode{Val: 2}

	fmt.Println(isSubTree(s, t))
}

func isSubTree(s *TreeNode, t *TreeNode) bool {
	if isSame(s, t) {
		return true
	}

	if s == nil {
		return false //向下不会再相等
	}

	return isSubTree(s.Left, t) || isSubTree(s.Right, t)
}

func isSame(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	return t1.Val == t2.Val && isSame(t1.Left, t2.Left) && isSame(t1.Right, t2.Right)
}
