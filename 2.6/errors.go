/*
 Нужно вызвать функцию divide которая делит два числа, но она возвращает не только результат деления,
 но и информацию об ошибке. В случае какой-либо ошибки вы должны вывести "ошибка", если ошибки нет - результат функции.
 */

func main() {
	var a,b int
	fmt.Scan(&a,&b)
	res, error := divide(a, b)

	if error == nil {
		fmt.Println(res)
	} else {
		fmt.Println("ошибка")
	}
}