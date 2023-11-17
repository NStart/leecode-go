package main

import "fmt"

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 0}
	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Left.Left = &TreeNode{Val: 6}
	root.Left.Left.Right = &TreeNode{Val: 8}
	root.Left.Left.Right.Left = &TreeNode{Val: 9}
	root.Left.Left.Right.Right = &TreeNode{Val: 10}
	root.Left.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Right = &TreeNode{Val: 1}
	fmt.Println(preorderTraversal(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var nums []int
	traverse(root, &nums)
	return nums
}

func traverse(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}

	*nums = append(*nums, root.Val)
	traverse(root.Left, nums)
	traverse(root.Right, nums)
}
