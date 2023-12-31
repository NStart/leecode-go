package main

import (
	"fmt"
	"strings"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}

	//       1
	//  /   \
	// 2     3
	//  \
	//   5
	fmt.Println(binaryTreePaths(root)) // [1->2->5 1->3]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var paths [][]int
	traverse(root, nil, &paths)

	var res []string
	for _, path := range paths {
		fmt.Println(strings.Replace(fmt.Sprint(path), " ", "->", -1))
		res = append(res, strings.Trim(strings.Replace(fmt.Sprint(path), " ", "->", -1), "[]"))
	}
	return res
}

// 和113一样，递归取经
func traverse(root *TreeNode, curPath []int, paths *[][]int) {
	if root == nil {
		return
	}

	curPath = append(curPath, root.Val)
	if root.Left == nil && root.Right == nil {
		*paths = append(*paths, curPath)
	}

	lPath := make([]int, len(curPath))
	copy(lPath, curPath)
	rPath := make([]int, len(curPath))
	copy(rPath, curPath)

	traverse(root.Left, lPath, paths)
	traverse(root.Right, rPath, paths)
}
