package main

import "fmt"

func main() {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
}

// 堆栈
type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() int {
	n := len(*s)
	res := (*s)[n-1]
	*s = (*s)[:n-1]
	return res
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

// 可以直接使用[]int  来实现queue
// 比较取巧的防范，用两个栈实现一个队列：x->s1,s2->
type Myqueue struct {
	s1, s2 Stack
}

func Constructor() Myqueue {
	return Myqueue{}
}

func (this *Myqueue) Push(x int) {
	this.s1.Push(x)
}

func (this *Myqueue) Pop() int {
	if this.s2.IsEmpty() {
		for !this.s1.IsEmpty() { //所有元素向前挪动
			this.s2.Push(this.s1.Pop())
		}
	}
	return this.s2.Pop()
}

func (this *Myqueue) Peek() int {
	if this.s2.IsEmpty() {
		for !this.s1.IsEmpty() { //所有元素向前挪动
			this.s2.Push(this.s1.Pop())
		}
	}
	return this.s2.Peek()
}

func (this *Myqueue) Empty() bool {
	return this.s2.IsEmpty() && this.s1.IsEmpty()
}
