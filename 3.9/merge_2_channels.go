package main

/*
	Необходимо написать функцию
	func merge2Channels(fn func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int).

	Описание ее работы:

	n раз сделать следующее

	прочитать по одному числу из каждого из двух каналов in1 и in2, назовем их x1 и x2.
	вычислить f(x1) + f(x2)
	записать полученное значение в out
	Функция merge2Channels должна быть неблокирующей, сразу возвращая управление.

	Функция fn может работать долгое время, ожидая чего-либо или производя вычисления.



	Формат ввода:

	количество итераций передается через аргумент n.
	целые числа подаются через аргументы-каналы in1 и in2.
	функция для обработки чисел перед сложением передается через аргумент fn.
	Формат вывода:

	канал для вывода результатов передается через аргумент out.
*/

import (
	"fmt"  // пакет используется для проверки выполнения условия задачи, не удаляйте его
	"time" // пакет используется для проверки выполнения условия задачи, не удаляйте его
)

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	id := make(chan int)
	val := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			a := <-in1
			b := <-in2

			done := make(chan int)
			go func() {
				done <- fn(a)
			}()
			go func() {
				done <- fn(b)
			}()

			go func(i int) {
				id <- i
				val <- (<-done) + (<-done)
			}(i)
		}
	}()

	go func() {
		m := make(map[int]int)
		for i := 0; i < n; i++ {
			m[<-id] = <-val
		}

		for i := 0; i < n; i++ {
			out <- m[i]
		}
	}()
}
