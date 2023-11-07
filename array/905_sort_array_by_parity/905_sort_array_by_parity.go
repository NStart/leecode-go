package main

import "fmt"

func main() {
	fmt.Println(sortArrayByParity([]int{1, 2, 3, 4, 5, 6, 7, 8}))
}

func sortArrayByParity(A []int) []int {
	isOdd := func(x int) bool {
		if x%2 == 0 {
			return false
		}
		return true
	}

	odds := make([]int, 0, len(A))
	evens := make([]int, 0, len(A))
	for _, a := range A {
		if isOdd(a) {
			odds = append(odds, a)
			continue
		}
		evens = append(evens, a)
	}
	return append(evens, odds...)
}
