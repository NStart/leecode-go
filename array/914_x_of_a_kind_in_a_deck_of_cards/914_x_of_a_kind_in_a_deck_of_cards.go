package main

import "fmt"

func main() {
	fmt.Println(hasGroupSizeX([]int{1, 1, 1, 2, 2, 2, 3, 3}))
	fmt.Println(hasGroupSizeX([]int{1, 1, 1, 1, 2, 2, 2, 2, 2, 2}))
}

// 统计各个数字出现的频率，看下频率之间是否能分组(存在公约数)
// 关键要理解分组等价是否等于公约数
func hasGroupSizeX(deck []int) bool {
	if len(deck) <= 1 {
		return false
	}

	freqs := make(map[int]int)
	for _, num := range deck {
		freqs[num]++
	}

	counts := make([]int, 0, len(freqs))
	for _, freq := range freqs {
		counts = append(counts, freq)
	}

	for i := 0; i < len(counts)-1; i++ {
		if gcd(counts[i], counts[i+1]) == 1 { //不能分组
			return false
		}
	}
	return true
}

// 辗转相除法求最大公约数
func gcd(x, y int) int {
	//两种方法选一种
	// for y != 0 {
	// 	x, y = y, x%y
	// }
	// return x

	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}
