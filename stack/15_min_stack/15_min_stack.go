package main

import "fmt"

func main() {
	s := MinStack{}
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Top())
	fmt.Println(s.GetMin())
	s.Pop()
	s.Pop()
	fmt.Println(s.GetMin())
	//fmt.Println(s.Top())
}

// 设计具有GetMin操作的的最小栈
type MinStack struct {
	min   int
	stack []int
}

func Construct() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 {
		this.min = x
	}
	if x < this.min {
		this.min = x
	}
	this.stack = append(this.stack, x)
}

// pop操作后需要解决min的操作
func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	if len(this.stack) > 0 {
		this.min = this.stack[0]
	}
	for _, elem := range this.stack {
		if elem < this.min {
			this.min = elem
		}
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.min
}
