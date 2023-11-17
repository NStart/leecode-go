package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1, root2 *TreeNode) bool {
	var leafs1, leafs2 []int
	traverse(root1, &leafs1)
	traverse(root2, &leafs2)
	if len(leafs1) != len(leafs2) {
		return false
	}

	for i, v := range leafs1 {
		if v != leafs2[i] {
			return false
		}
	}
	return true
}

func traverse(root *TreeNode, leafs *[]int) {
	if root == nil {
		return
	}

	traverse(root.Left, leafs)
	traverse(root.Right, leafs)
	if root.Left == nil && root.Right == nil {
		*leafs = append(*leafs, root.Val)
	}
}
