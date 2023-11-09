package main

import "fmt"

func main() {
	fmt.Println(detectCapitalUse("USA"))
	fmt.Println(detectCapitalUse("Leetcode"))
	fmt.Println(detectCapitalUse("mL"))
}

func detectCapitalUse(word string) bool {
	if len(word) <= 1 {
		return true
	}

	firstUp := rune(word[0]) >= 'A' && rune(word[0]) <= 'Z'
	allUpCase, allLowCase := true, true
	for i := 1; i < len(word); i++ {
		if rune(word[i]) < 'A' || rune(word[i]) > 'Z' {
			allUpCase = false
		}
		if rune(word[i]) < 'a' || rune(word[i]) > 'z' {
			allLowCase = false
		}
	}
	if firstUp && (allLowCase || allUpCase) {
		return true
	}

	if !firstUp && allLowCase {
		return true
	}
	return false
}
