package main

import "fmt"

func main() {
	tree := &Tree{}
	tree.root = &TreeNode{Val: 5}
	tree.root.Left = &TreeNode{Val: 4}
	tree.root.Right = &TreeNode{Val: 8}
	fmt.Println(hasPathSum(tree.root, 13)) //true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	root *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	var sums []int
	traverse(root, &sums)
	for _, v := range sums {
		if v == sum {
			return true
		}
	}
	return false
}

// 遍历将父节点的值存储到子节点中
// 若是叶子节点则算作是一个sum
func traverse(root *TreeNode, sums *[]int) {
	if root == nil {
		return
	}

	//叶子节点
	if root.Left == nil && root.Right == nil {
		*sums = append(*sums, root.Val)
	}

	if root.Left != nil {
		root.Left.Val += root.Val
	}

	if root.Right != nil {
		root.Right.Val += root.Val
	}
	traverse(root.Left, sums)
	traverse(root.Right, sums)

}
