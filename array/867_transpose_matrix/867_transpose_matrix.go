package main

import "fmt"

func main() {
	fmt.Println(transpose([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
}

func transpose(A [][]int) [][]int {
	if len(A) <= 0 {
		return A
	}

	n, m := len(A), len(A[0])
	res := make([][]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Println(A[i][j])
			res[j] = append(res[j], A[i][j])
		}
	}
	return res
}
