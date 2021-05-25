package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

/*
 На вход подается строка. Нужно определить, является ли она правильной или нет.
 Правильная строка начинается с заглавной буквы и заканчивается точкой. Если строка правильная - вывести Right
 иначе - вывести Wrong
 */

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	rs := []rune(text)

	if unicode.IsUpper(rs[0]) && rs[len(rs)-1]==46 {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
