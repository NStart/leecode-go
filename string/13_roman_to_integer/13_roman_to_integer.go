package main

import "fmt"

func main() {
	fmt.Println(romanInt("MCMXCIV"))
	fmt.Println(romanInt("LVIII"))
}

var commonNums = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

var preNums = map[string]int{
	"IV": 4,
	"IX": 9,
	"XL": 40,
	"XC": 90,
	"CD": 400,
	"CM": 900,
}

func romanInt(s string) int {
	n := len(s)
	if n == 1 {
		return commonNums[s]
	}

	res := 0
	for len(s) > 0 {
		if len(s) >= 2 {
			if num, ok := preNums[s[:2]]; ok {
				res += num
				s = s[2:]
				continue
			}
		}
		res += commonNums[string(s[:1])]
		s = s[1:]
	}
	return res
}
