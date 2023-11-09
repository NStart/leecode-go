package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
}

func firstUniqChar(s string) int {
	runeMap := make(map[rune]int)
	for _, r := range s {
		runeMap[r]++
	}

	for i, r := range s {
		if runeMap[r] == 1 {
			return i
		}
	}
	return -1
}
