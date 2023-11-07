package main

import "fmt"

func main() {
	fmt.Println(largeGroupPositions("aaa"))
	fmt.Println(largeGroupPositions("abcxxxxzzy"))
	fmt.Println(largeGroupPositions("abcdddeeeeaabbbcd"))
}

func largeGroupPositions(S string) [][]int {
	n := len(S)
	if n <= 2 {
		return nil
	}

	var locs [][]int
	slow, fast := 0, 0
	for slow < n {
		for fast < n && S[fast] == S[slow] {
			fast++
		}
		if fast-slow >= 3 {
			locs = append(locs, []int{slow, fast - 1})
		}
		if fast > slow {
			slow = fast - 1
		}
		slow++
		fast = slow
	}
	return locs
}
