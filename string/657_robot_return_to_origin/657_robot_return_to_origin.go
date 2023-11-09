package main

import "fmt"

func main() {
	fmt.Println(judgeCircle("UD"))
	fmt.Println(judgeCircle("LL"))
}

// 想要回到原点，向左必须抵消向右，向上必须抵消向下
func judgeCircle(s string) bool {
	horizontal, vertical := 0, 0
	for _, move := range []rune(s) {
		switch move {
		case 'R':
			horizontal++
		case 'L':
			horizontal--
		case 'U':
			vertical++
		case 'D':
			vertical--
		}
	}
	return horizontal == 0 && vertical == 0
}

// func judgeCircle(s string) bool {
// 	horizontal, vertical := 0, 0
// 	for _, move := range s {
// 		switch move {
// 		case 'R':
// 			horizontal++
// 		case 'L':
// 			horizontal--
// 		case 'U':
// 			vertical++
// 		case 'D':
// 			vertical--
// 		}
// 	}
// 	return horizontal == 0 && vertical == 0
// }
