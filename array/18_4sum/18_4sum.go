package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{-3, 0, 0, 1, 2}, 0))
	fmt.Println(fourSum([]int{0, 0, 0, 0, 0}, 0))
}

// N数之和的本质是，在有序的数组内，寻找N个数恰好是S
// 解决办法还是3sum 3sumclosest 的双指针法，不过需要外部N-2 层循环，内部双指针循环即可
// 注意双指针在遍历时外部所有循环要去重，指针移动时也要去重
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int

	for i := 0; i < n-1; i++ {
		if i > 0 && nums[i] == nums[i-1] { //去重
			continue
		}

		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] { //去重2 //注意条件：j>i+1与i>0相同都是为了排除第一个相同数
				continue
			}
			head, tail := j+1, n-1
			for head < tail {
				sum := nums[i] + nums[j] + nums[head] + nums[tail]
				switch {
				case sum < target: //向后走
					head++
				case sum > target: //向前走
					tail--
				case sum == target: //向前向后走
					res = append(res, []int{nums[i], nums[j], nums[head], nums[tail]})
					//去重3：注意for循环条件的判断，避开死循环
					for head < tail && nums[head] == nums[head-1] {
						head++
					}
					for head < tail && nums[tail] == nums[tail-1] {
						tail--
					}
					head++
					tail--
				}
			}
		}
	}
	return res
}
