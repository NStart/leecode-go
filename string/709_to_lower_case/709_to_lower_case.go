package main

import "fmt"

func main() {
	fmt.Println(toLowerCase("HelLo"))
}

func toLowerCase(str string) string {
	const charDiff = 'a' - 'A'
	newStr := ""
	for _, r := range []rune(str) {
		if r < 'A' || r > 'Z' {
			newStr += string(r)
			continue
		}
		newStr += string(r + charDiff)
	}

	return newStr
}
