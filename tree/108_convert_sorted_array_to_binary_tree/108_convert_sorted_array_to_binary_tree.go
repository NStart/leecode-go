package main

import "fmt"

/*
二叉搜索树特点
有序性： 二叉搜索树中的每个节点都有一个值，且左子树中所有节点的值都小于等于该节点的值，右子树中所有节点的值都大于等于该节点的值。这个有序性贯穿整个树的结构，使得查找、插入和删除等操作更加高效。

递归定义： 二叉搜索树的左子树和右子树都是二叉搜索树。这意味着可以通过递归的方式对整个树进行操作。

唯一性： 二叉搜索树中不存在两个节点具有相同的值，确保了树中每个值的唯一性。

查找效率高： 由于有序性，可以通过比较节点的值来快速定位目标值，从而提高查找效率。

插入和删除操作： 插入和删除操作都可以在O(log n)的时间内完成，其中n是树中节点的数量。这是在平均情况下的时间复杂度，最坏情况下可能会达到O(n)，例如当树退化成链表的情况。

中序遍历有序序列： 对二叉搜索树进行中序遍历会得到一个有序的序列，这对于排序和范围查询等操作很有用。

平衡性： 为了维持搜索树的高效性能，通常会引入平衡性的概念，即保持树的左右子树高度相差不大。平衡二叉搜索树（如 AVL 树、红黑树等）是为了解决二叉搜索树在某些情况下可能退化成链表的问题而设计的。

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := sortedArrayToBST([]int{0, 1, 2, 3, 4, 5})
	fmt.Println(root)
	fmt.Println(root.Left)
	fmt.Println(root.Left.Left)
	fmt.Println(root.Left.Right)
	fmt.Println(root.Right)
	fmt.Println(root.Right.Left) // bingo
}

func sortedArrayToBST(nums []int) *TreeNode {
	return construct(nums)
}

func construct(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	if n == 1 {
		return &TreeNode{Val: nums[0]}
	}

	//尽可能地取中间的值作为根节点
	mid := n / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = construct(nums[:mid])    //左子数组作为左子树
	root.Right = construct(nums[mid+1:]) //右子数组作为右子树 //达到AVL树特点
	return root
}
