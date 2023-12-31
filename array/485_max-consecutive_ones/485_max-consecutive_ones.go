package main

import "fmt"

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 0, 0, 1, 1, 1, 1}))
}

func findMaxConsecutiveOnes(nums []int) int {
	n := len(nums)
	count, maxCount := 0, 0
	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			count = 1
			for j := i + 1; j < n; j++ {
				if nums[j] == 0 {
					break
				}
				count++
				i++
			}
			maxCount = max(count, maxCount)
		}
	}
	return maxCount
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
