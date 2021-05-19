func vote(x int, y int, z int) int {
	r := x+y+z

	if (r >= 2) {
		return 1
	}

	return 0
}