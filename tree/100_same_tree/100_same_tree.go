package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return same(p, q)
}

func same(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil { //先&&再||相当于1+2=3
		return false
	}

	return node1.Val == node1.Val && same(node1.Left, node2.Left) && same(node1.Right, node2.Right)
}
