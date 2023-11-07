package main

import "fmt"

func main() {
	fmt.Println(matrixReshape([][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}, 2, 3))
}

func matrixReshape(nums [][]int, r, c int) [][]int {
	if len(nums) < 0 || len(nums[0]) < 0 {
		return nil
	}

	row, column := len(nums), len(nums[0])
	if row*column != r*c {
		return nums
	}

	allNums := make([]int, 0, row*column)
	for _, rowNums := range nums {
		allNums = append(allNums, rowNums...)
	}

	matrix := make([][]int, r)
	for i := 0; i < r; i++ {
		matrix[i] = allNums[i*c : (i+1)*c]
	}
	return matrix
}
