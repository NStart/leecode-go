package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello wordld"))
	fmt.Println(lengthOfLastWord("     "))
}

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)
	fmt.Println(words)
	if len(words) <= 0 {
		return 0
	}

	return len(words[len(words)-1])
}
