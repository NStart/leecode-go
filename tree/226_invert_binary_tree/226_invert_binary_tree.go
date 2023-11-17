package main

import "fmt"

func main() {
	tree := &Tree{}
	tree.root = &TreeNode{Val: 4}
	for _, v := range []int{2, 7, 1, 3, 6, 9} {
		tree.BFSInsert(v)
	}
	tree.BFSTravese(tree.root)
	fmt.Println()
	invertTree(tree.root)
	tree.BFSTravese(tree.root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	root *TreeNode
}

// 翻转二叉树其实是镜像翻转，swap节点操作
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	invertTree(root.Left)
	invertTree(root.Right)

	swapNode := root.Left
	root.Left = root.Right
	root.Right = swapNode
	return root
}

func (t *Tree) BFSInsert(val int) {
	if t.root == nil {
		t.root = &TreeNode{Val: val}
		return
	}

	queue := []*TreeNode{t.root}
	i := 1
	fmt.Println()
	fmt.Println()
	fmt.Println()
	for len(queue) > 0 {
		fmt.Println("len: ", len(queue))
		//第一次进来，这边node就是根节点
		//第二次进来，这边node也是根节点
		//第三次进来，这边node也是根节点
		//第四次进来，这边node也是根节点
		//不管第几次进来，这边node都是根节点
		node := queue[0]
		fmt.Println(node)

		//第一次进来queue是[]
		//第二次进来queue也是[]
		//第三次进来queue也是[]
		//第四次进来queue也是[]
		//不管第几次进来，他的第一次for循环都是[]
		queue = queue[1:]
		fmt.Println("queue: ", queue)
		//第一次进来node.Left是nil所以会把值赋值给左节点，返回，不会再次循环，整个函数跳出。因为这边是return,不是break或者continue
		//第二次进来node.Left!=nil，queue进行append,然后node.Right==nil，进行赋值，再次跳出函数
		//第三次进来node.Left!=nil,queue进行append，append了Left和Right
		//再次for循环，node变成左节点，左节点的node.Left==nil,进行赋值，return跳出整个函数。
		//第四次进来node.Left!=nil，queue进行append,node.Right!=nil，queue进行append,
		//再次进入for循环，node是节点的左节点，左节点的node.Left==nil，进行赋值，return跳出整个函数
		//第五次进来，queue进行左右节点赋值，
		//再次for循环，node变成node的左节点，左节点的node.Right==nil，进行赋值，return跳出整个函数
		//第六次进来，queue进行左右节点append,再次for循环，node变成node左节点，再进行左节点的左右节点append,
		//再次for循环，此时queue[0]已经从右节点开始，所以开始往有节点的左节点依次赋值
		//然后第7次...
		if node.Left == nil {
			node.Left = &TreeNode{Val: val}
			i++
			return
		} else {
			fmt.Println("队列添加左节点")
			queue = append(queue, node.Left)
		}

		if node.Right == nil {
			node.Right = &TreeNode{Val: val}
			i++
			return
		} else {
			fmt.Println("队列添加右节点")
			queue = append(queue, node.Right)
		}
		i++
	}
}

func (t *Tree) BFSTravese(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Println(node.Val, "")
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	fmt.Println()
}
