package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
	fmt.Println(pivotIndex([]int{-1, -1, -1, -1, -1, 0}))
	fmt.Println(pivotIndex([]int{-1, -1, -1, -1, 0, -1}))
}

func pivotIndex(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return -1
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}

	leftSum := 0
	for i := 0; i < n; i++ {
		if leftSum == sum-nums[i]-leftSum {
			return i
		}
		leftSum += nums[i]
	}

	return -1
}
