func sumInt(r ...int) (int,int) {
	var s int

	for _,v := range r {
		s += v
	}

	return len(r),s
}