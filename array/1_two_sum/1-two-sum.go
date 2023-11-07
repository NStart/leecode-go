package main

import "fmt"

func main() {
	fmt.Println(twoSum3([]int{2, 3, 4}, 6))
}

// 优化后的哈希表索引方式
// 看 twoSum2 会发现进行了2次for循环，可以进行合并优化，一边遍历，一边检查
func twoSum3(nums []int, target int) []int {
	num2index := make(map[int]int, len(nums))
	for i, num := range nums {
		pair := target - num

		if j, ok := num2index[pair]; ok && i != j {
			return []int{j, i} // 注意返回值顺序，向后遍历 nums，所以 i 在 j 后
		}
		num2index[num] = i

		fmt.Println("end ", i, num, num2index)
	}
	return nil
}
