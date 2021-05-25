package main

import (
	"fmt"
	"strings"
)

// Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку

func main() {
	var s,_s  string
	fmt.Scan(&s)
	sl := []rune(s)
	var _sl []rune

	for _, v := range sl {
		if ! strings.ContainsRune(_s, v) {
			_s += string(v)
		} else {
			_sl = append(_sl, v)
		}
	}

	if len(_sl) > 0 {
		var _s1 string
		for _, v := range []rune(_s) {
			if ! isContainsRune(v, _sl) {
				_s1 += string(v)
			}
		}
		fmt.Println(_s1)
	} else {
		fmt.Println(_s)
	}
}

func isContainsRune(r rune, _r []rune) bool {
	for _, v := range _r {
		if v == r {
			return true
		}
	}

	return false
}
