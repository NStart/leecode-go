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
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	//    3
	//   / \
	//  9  20
	//    /  \
	//   15   7
	fmt.Println(sumOfLeftLeaves(root)) // 24 // 9+15
}

// 左叶子则退出递归
func sumOfLeftLeaves(root *TreeNode) int {
	var sum int
	traverse(root, &sum)
	return sum
}

func traverse(root *TreeNode, sum *int) {
	if root == nil {
		return
	}

	traverse(root.Left, sum)
	traverse(root.Right, sum)

	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil { //左叶子
		*sum += root.Left.Val
	}

}
