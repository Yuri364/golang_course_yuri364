package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
 Дана строка, содержащая только английские буквы (большие и маленькие). Добавить символ ‘*’ (звездочка)
 между буквами (перед первой буквой и после последней символ ‘*’ добавлять не нужно).
 */

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	var newString string
	chars := []rune(text)

	for i, v := range chars {
		newString += string(v)
		if i != len(chars)-1 {
			newString += "*"
		}
	}

	fmt.Print(newString)
}