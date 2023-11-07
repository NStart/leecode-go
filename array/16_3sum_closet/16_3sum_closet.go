package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSumCloset1([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{0, 2, 1, -3}, 1))
	fmt.Println(threeSumClosest([]int{1, 2, 4, 8, 16, 32, 64, 128}, 84))
}

// 同样也是指针法
// 和15一样，排序预处理能知道双指针移动的方向，记录最小的abs
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	minAbs := 1<<31 - 1
	minSum := 0

	for i, num := range nums {
		if i > 0 && nums[i] == nums[i-1] { // 优化：可选的去重
			continue
		}

		l, r := i+1, n-1
		for l < r {
			sum := num + nums[l] + nums[r]
			if absInt(target-sum) < minAbs {
				minAbs = absInt(target - sum)
				minSum = sum
			}
			switch {
			case sum < target:
				l++
			case sum > target:
				r--
			default:
				return target
			}
		}
	}
	return minSum
}

// 暴力三层遍历，求解最接近的和
func threeSumCloset1(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	minAbs := absInt(nums[0] + nums[1] + nums[2] - target)
	minSum := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				sum := nums[i] + nums[j] + nums[k]
				if absInt(sum-target) <= minAbs {
					minSum = sum
					minAbs = absInt(sum - target)
				}
			}
		}
	}
	return minSum
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
