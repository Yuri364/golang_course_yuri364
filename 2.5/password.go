package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

/*
 Сделать проверку подходит ли пароль вводимый пользователем под заданные требования.
 Длина пароля должна быть не менее 5 символов, он должен содержать только цифры и буквы латинского алфавита.
 На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"
 */

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	pass := []rune(text)
	if len(pass) < 5 || ! checkPass(pass) {
		fmt.Println("Wrong password")
	} else {
		fmt.Println("Ok")
	}
}

func checkPass(r []rune) bool {
	for _, v := range r {
		if ! unicode.Is(unicode.Latin, v) && ! unicode.IsDigit(v) {
			return false
		}
	}

	return true
}
