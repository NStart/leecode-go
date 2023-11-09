package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(countSegments("hello world"))
}

func countSegments(s string) int {
	return len(strings.Fields(s))
}
