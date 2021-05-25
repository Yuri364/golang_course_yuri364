package main

import (
	"bufio"
	"fmt"
	"os"
)

// Дана строка, содержащая только десятичные цифры. Найти и вывести наибольшую цифру.

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	chars := []rune(text)

	n := int(chars[0])
	for _, v := range chars {
		if int(v) > n {
			n = int(v)
		}
	}

	fmt.Println(string(n))
}
