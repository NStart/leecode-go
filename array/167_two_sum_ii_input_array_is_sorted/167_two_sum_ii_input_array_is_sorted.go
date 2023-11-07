package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
	fmt.Println(twoSum([]int{-1, 0}, -1))
}

func twoSum(numbers []int, target int) []int {
	head, tail := 0, len(numbers)-1
	for head < tail {
		switch {
		case numbers[head]+numbers[tail] > target:
			tail--
		case numbers[head]+numbers[tail] < target:
			head++
		default:
			return []int{head + 1, tail + 1}
		}
	}
	return nil
}
