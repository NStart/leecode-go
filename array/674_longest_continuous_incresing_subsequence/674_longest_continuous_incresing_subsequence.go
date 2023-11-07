package main

import "fmt"

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))
	fmt.Println(findLengthOfLCIS([]int{2, 2, 2, 2, 2}))
}

// 依旧双指针
func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return len(nums)
	}

	slow := 0
	maxCount := 0
	for slow < n-1 {
		count := 1
		fast := slow
		for fast < n-1 && nums[fast] < nums[fast+1] {
			count++
			fast++
		}
		slow++
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}
