package main

import "fmt"

func main() {
	s := Mystack{}
	fmt.Println(s.Empey())
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println(s.Pop())
}

// 两个队列实现一个栈
// [...enque...] <-
// [...deque...] <-
type Mystack struct {
	enque, deque []int
}

func Constructor() Mystack {
	return Mystack{}
}

func (this *Mystack) Push(x int) {
	this.enque = append(this.enque, x)
}

func (this *Mystack) Pop() int {
	//把enqueue队列元素按顺序append到deque队列中
	n := len(this.enque)
	for i := 0; i < n-1; i++ {
		this.deque = append(this.deque, this.enque[0])
		this.enque = this.enque[1:]
	}
	peek := this.enque[0]
	this.enque = this.deque
	this.deque = nil
	return peek
}

func (this *Mystack) Top() int {
	peek := this.Pop()
	this.enque = append(this.enque, peek)
	return peek
}

func (this *Mystack) Empey() bool {
	return len(this.enque) == 0
}
