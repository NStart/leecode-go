package main

import "fmt"

func main() {
	fmt.Println(removeDuplicates1([]int{1, 1, 1, 2, 2}))
}

// 针对有序数组，双指针法是十分常见且有用的
func removeDuplicates1(nums []int) int {
	slow, fast := 0, 0
	for fast < len(nums)-1 {
		if nums[fast] != nums[fast+1] { //相邻的数不相等
			slow++
			fast++
			nums[slow] = nums[fast]
			continue
		}
		fast++
	}
	fmt.Println(nums)
	return slow + 1
}

//充分利用数组有序的已知条件
func removeDuplicates2(nums []int) int {
	n := len(nums)
	l, r := 0, 1
	for r < n-1 {
		if nums[l] < nums[r] { //比我大就放到我的下一个
			l++
			nums[l], nums[r] = nums[r], nums[l]
		}
		r++
	}
	return l + 1
}
