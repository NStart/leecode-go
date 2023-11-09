package main

import "fmt"

func main() {
	fmt.Println(reverseString("hello"))
}

func reverseString(str string) string {
	s := []byte(str)
	left, right := 0, len(s)-1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return string(s)
}
