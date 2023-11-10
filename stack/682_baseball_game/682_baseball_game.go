package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))                 // 30
	fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"})) // 27
}

func calPoints(ops []string) int {
	var s Stack
	for _, str := range ops {
		switch str {
		case "+":
			if !s.IsEmpty() {
				originPeek := s.Pop()
				sum := s.Peek() + originPeek
				s.Push(originPeek)
				s.Push(sum)
			}
		case "D":
			if !s.IsEmpty() {
				s.Push(2 * s.Peek())
			}
		case "C":
			if !s.IsEmpty() {
				s.Pop()
			}
		default:
			n, _ := strconv.Atoi(str)
			s.Push(n)
		}
	}

	sum := 0

	for !s.IsEmpty() {
		sum += s.Pop()
	}
	return sum
}

type Stack []int

// IsEmpty: 检查栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push: 向栈顶添加元素
func (s *Stack) Push(n int) {
	*s = append(*s, n)
}

// Pop: 移除并返回栈顶元素
func (s *Stack) Pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

// Peek: 返回栈顶元素，但不移除
func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}
