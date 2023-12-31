package main

import "fmt"

func main() {
	fmt.Println(rotateDigits(2))
	fmt.Println(rotateDigits(20))
}

func rotateDigits(N int) int {
	counts := 0
	for n := 1; n <= N; n++ {
		if valid, res := rotate(n); valid && n != res {
			counts++
		}
	}
	return counts
}

func rotate(num int) (valid bool, res int) {
	valid = true
	var nums []int
	for num > 0 {
		nums = append(nums, num%10)
		num /= 10
	}

	//fmt.Println(nums)

	for i, n := range nums { //旋转
		switch n {
		case 0, 1, 8:
		case 2:
			nums[i] = 5
		case 5:
			nums[i] = 2
		case 6:
			nums[i] = 9
		case 9:
			nums[i] = 6
		default:
			valid = false
		}
	}

	//fmt.Println(nums)
	if valid {
		for i, n := range nums {
			res += n * pow(10, i)
		}
	}
	return
}

func pow(x, y int) int {
	res := 1
	for y > 0 {
		res *= x
		y--
	}
	return res
}
