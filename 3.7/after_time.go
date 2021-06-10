package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

/*
	На стандартный ввод подается строковое представление двух дат, разделенных запятой
	(формат данных смотрите в примере).

	Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода
	между меньшей и большей датами.
 */

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace(text)

	dates := strings.Split(text, ",")

	layout := "02.01.2006 15:04:05"

	str1 := strings.TrimSpace(dates[0])
	str2 := strings.TrimSpace(dates[1])

	date1, _ := time.Parse(layout, str1)
	date2, _ := time.Parse(layout, str2)

	if date1.After(date2) {
		fmt.Println(date1.Sub(date2))
	} else {
		fmt.Println(date2.Sub(date1))
	}
}
