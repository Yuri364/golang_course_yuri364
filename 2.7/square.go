package main

import (
	"fmt"
)

// На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.

func main() {
	var n, _n int
	fmt.Scan(&n)
	var s []int

	for n>0{
		_n = n%10
		s = append(s, _n)
		n = n/10
	}

	for i:=len(s)-1; i>=0; i-- {
		fmt.Print(s[i]*s[i])
	}
}
