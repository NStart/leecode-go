package main

import "fmt"

//二叉搜索树特性
//二叉树结构： 每个节点最多有两个子节点，分别称为左子节点和右子节点。
//节点值大小关系： 对于树中的每个节点，其左子树中的所有节点的值都小于该节点的值，而右子树中的所有节点的值都大于该节点的值。这个特性使得在二叉搜索树中，通过比较节点值可以快速定位特定值。
//中序遍历有序： 对二叉搜索树进行中序遍历（Inorder Traversal）会按照升序（从小到大）访问树中的所有节点。

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 7}

	//        4
	//       / \
	//      2   7
	//     / \
	//    1   3
	fmt.Println(searchBST(root, 2)) // 2
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	switch {
	case root.Val > val:
		return searchBST(root.Left, val)
	case root.Val < val:
		return searchBST(root.Right, val)
	default:
		return root
	}
}
