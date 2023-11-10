package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(backspaceCompare("a#c", "b"))
	fmt.Println(backspaceCompare("ab#c", "ad#c"))
}

func backspaceCompare(S, T string) bool {
	var s1, s2 Stack
	traverse := func(str string, s *Stack) {
		for _, r := range str {
			if r == '#' {
				if !s.IsEmpty() {
					s.Pop()
				}
			} else {
				fmt.Println(int(r), rune(int(r)), string((int(r))))
				s.Push(int(r))
			}
		}
	}

	traverse(S, &s1)
	traverse(T, &s2)
	var s, t string
	for !s1.IsEmpty() {
		s += string(rune(s1.Pop()))
	}
	for !s2.IsEmpty() {
		t += string(rune(s2.Pop()))
	}
	return strings.EqualFold(s, t)
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
