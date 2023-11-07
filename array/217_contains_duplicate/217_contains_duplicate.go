package main

import "fmt"

func main() {
	fmt.Println(containDuplicate([]int{1, 2, 3, 1}))
}

// 没什么好说的，哈希表走你
func containDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = 0
	}
	return false
}
