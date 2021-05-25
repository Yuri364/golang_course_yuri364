package main

import "fmt"

/*
 На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является
 паллиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.
 */

func main() {
	var text string
	fmt.Scan(&text)

	rs := []rune(text)
	_rs := make([]rune, len(rs))

	var j int
	for i:=len(rs)-1; i>=0; i-- {
		_rs[j] = rs[i]
		j++
	}

	if string(rs) == string(_rs) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
