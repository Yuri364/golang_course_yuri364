func minimumFromFour() int {
	var a,b,c,d,r int
	fmt.Scan(&a,&b,&c,&d)

	s := [4]int{a,b,c,d}
	r = a

	for _, v := range s {
		if r > v {
			r = v
		}
	}
	return r
}