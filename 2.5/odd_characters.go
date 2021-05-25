package main

import (
	"fmt"
)

// На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)

func main() {
	var s string
	fmt.Scan(&s)
	sl := []rune(s)
	var _sl []rune

	for k, v := range sl {
		if k % 2 != 0 {
			_sl = append(_sl, v)
		}
	}

	fmt.Println(string(_sl))
}
