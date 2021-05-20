func test(x1 *int, x2 *int) {

	var _x1, _x2 = *x1, *x2

	*x1 = _x2
	*x2 = _x1

	fmt.Println(_x2, _x1)
}