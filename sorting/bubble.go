package sorting

// Сортировка пузырьком
func BubbleSort(ar []int) []int {
	for i := 0; i < len(ar); i++ {
		for j := i; j < len(ar); j++ {
			if ar[i] > ar[j] {
				swap(ar, i, j)
			}
		}
	}

	return ar
}