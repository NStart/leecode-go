package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 3}
	fmt.Println(increasingBST(root))
}

//感觉下面的算法应该不对，忽略了的右节点？？？
// 可以先递归中序遍历把数记在数组，再创建右斜树
// 不过用stack 可以一遍遍历一遍建树
func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var newRoot *TreeNode
	var newCur *TreeNode

	var stack []*TreeNode
	cur := root
	for cur != nil || len(stack) > 0 {
		//先走到最左节点记下新头节点
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
			continue
		}

		if newRoot == nil {
			newRoot = stack[len(stack)-1] //leftest
			newCur = newRoot
		} else {
			newCur.Right = stack[len(stack)-1]
			newCur = newCur.Right
		}

		stack = stack[:len(stack)-1] // pop
		newCur.Left = nil            // 把原节点的左子树给砍掉
		cur = newCur.Right

	}
	newCur.Right = nil
	return newRoot
}
