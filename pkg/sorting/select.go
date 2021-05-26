package sorting

// Сортировка выбором
func SelectSort(ar []int) []int  {
	for i := 0; i < len(ar)-1; i++ {
		min := i
		for j := i + 1; j < len(ar); j++ {
			if ar[min] > ar[j] {
				min = j
			}
		}

		if min != i {
			swap(ar, i, min)
		}
	}

	return ar
}