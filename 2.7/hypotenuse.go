package main

import (
	"fmt"
	"math"
)

// На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы

func main() {
	var a,b float64
	fmt.Scan(&a,&b)

	fmt.Println(math.Sqrt((a*a)+(b*b)))
}