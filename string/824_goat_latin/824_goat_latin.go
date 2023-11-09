package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toGoatLatin("I speak Goat Latin"))
	fmt.Println(toGoatLatin("Each word consists of lowercase and uppercase letters only"))
}

func toGoatLatin(S string) string {
	words := strings.Fields(S)
	newWords := make([]string, 0, len(words))
	for i, word := range words {
		switch word[0] {
		case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
			word += "ma"
		default:
			word = word[1:] + string(word[0]) + "ma"
		}
		i += 1
		for i > 0 {
			word += "a"
			i--
		}
		newWords = append(newWords, word)
	}
	return strings.Join(newWords, " ")
}
