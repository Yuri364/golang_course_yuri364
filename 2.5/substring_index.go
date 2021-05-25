package main

import (
	"fmt"
	"strings"
)

/*
 Даются две строки X и S. Нужно найти и вывести первое вхождение подстроки S в строке X.
 Если подстроки S нет в строке X - вывести -1
 */

func main() {
	var x, s string

	fmt.Scan(&x, &s)
	fmt.Println(strings.Index(x, s))
}

