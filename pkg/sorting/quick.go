package sorting

// Быстрая сортировка
func QuickSort(ar []int) []int  {
	if len(ar) <= 1 {
		return ar
	}

	split := partition(ar)

	QuickSort(ar[:split])
	QuickSort(ar[split:])

	return ar
}

func partition(ar []int) int {
	pivot := ar[len(ar)/2]

	left := 0
	right := len(ar) - 1

	for {
		for ; ar[left] < pivot; left++ {
		}

		for ; ar[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}

		swap(ar, left, right)
	}
}
