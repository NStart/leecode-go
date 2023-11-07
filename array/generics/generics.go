package main

import (
	"fmt"
	"strings"
)

func Swap[T any](a, b T) (T, T) {
	return b, a
}

// 泛型切片
func Filter[T any](s []T, f func(T) bool) []T {
	var result []T
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// 泛型堆栈
type Stack[T any] []T

func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

func (s *Stack[T]) Pop() T {
	fmt.Println(len(*s))
	if len(*s) == 0 {
		var zeroValue T
		return zeroValue
	}
	value := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return value
}

// 泛型映射Map
func Map[T, U any](items []T, f func(T) U) []U {
	var result []U
	for _, item := range items {
		result = append(result, f(item))
	}
	return result
}

func main() {
	x, y := 3, 4
	x, y = Swap(x, y)

	fmt.Println(x, y)

	s1, s2 := "Hello", "World"
	s1, s2 = Swap(s1, s2)
	fmt.Println(s1, s2)

	//泛型切片
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	isEven := func(x int) bool {
		return x%2 == 0
	}
	evens := Filter(numbers, isEven)
	fmt.Println(evens)

	words := []string{"apple", "banana", "cherry", "date", "elderberry"}
	hasA := func(s string) bool {
		return strings.Contains(s, "a")
	}
	aWords := Filter(words, hasA)
	fmt.Println(aWords)

	//泛型堆栈
	intStack := new(Stack[int])
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Println(intStack.Pop())
	fmt.Println(intStack.Pop())

	stringStack := new(Stack[string])
	stringStack.Push("apple")
	stringStack.Push("banana")

	fmt.Println(stringStack.Pop())

	//泛型映射Map
	numbers2 := []int{1, 2, 3, 4, 5}
	double := func(x int) int {
		return x * 2
	}
	doubleNumber := Map(numbers2, double)
	fmt.Println(doubleNumber)

	words2 := []string{"apple", "banana", "cherry"}
	length := func(s string) int {
		return len(s)
	}
	wordLengths := Map(words2, length)
	fmt.Println(wordLengths)

}
