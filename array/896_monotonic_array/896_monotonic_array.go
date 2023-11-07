package main

import "fmt"

func main() {
	fmt.Println(isMonotonic([]int{1, 2, 3}))
	fmt.Println(isMonotonic([]int{2, 2, 4, 1}))
}

func isMonotonic(A []int) bool {
	isUp, isDown := true, true
	for i := 0; i < len(A)-1; i++ {
		if A[i] < A[i+1] {
			isDown = false
		}
		if A[i] > A[i+1] {
			isUp = false
		}

		if !isDown && !isUp {
			return false
		}
	}
	return isUp || isDown
}
