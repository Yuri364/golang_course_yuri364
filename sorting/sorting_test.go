package sorting

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var testData = []struct {
		nums     []int
		expected []int
	} {
		{[]int{34,90,10,37,49}, []int{10,34,37,49,90}},
		{[]int{190,999,521,877,605}, []int{190,521,605,877,999}},
		{[]int{1547,2294,5284,9475,1203}, []int{1203,1547,2294,5284,9475}},
	}

	for _, item := range testData {
		s := QuickSort(item.nums)
		sliceEqual(s, item.expected, t)
	}
}

func TestBubbleSort(t *testing.T) {
	var testData = []struct {
		nums     []int
		expected []int
	} {
		{[]int{98,13,59,76,14}, []int{13,14,59,76,98}},
		{[]int{156,390,528,744,546}, []int{156,390,528,546,744}},
		{[]int{5791,4309,9300,8695,7299}, []int{4309,5791,7299,8695,9300}},
	}

	for _, item := range testData {
		s := BubbleSort(item.nums)

		sliceEqual(s, item.expected, t)
	}
}

func TestSelectSort(t *testing.T) {

	var testData = []struct {
		nums     []int
		expected []int
	} {
		{[]int{61,1,95,7,85}, []int{1,7,61,85,95}},
		{[]int{474,2,26,558,272}, []int{2,26,272,474,558}},
		{[]int{6648,2426,2837,1425,4849}, []int{1425,2426,2837,4849,6648}},
	}

	for _, item := range testData {
		s := SelectSort(item.nums)

		sliceEqual(s, item.expected, t)
	}
}

func TestInsertionSort(t *testing.T) {
	var testData = []struct {
		nums     []int
		expected []int
	} {
		{[]int{63,4,52,24,9}, []int{4,9,24,52,63}},
		{[]int{786,384,208,792,561}, []int{208,384,561,786,792}},
		{[]int{1036,8996,5566,5904,6506}, []int{1036,5566,5904,6506,8996}},
	}

	for _, item := range testData {
		s := InsertionSort(item.nums)

		sliceEqual(s, item.expected, t)
	}
}

func sliceEqual(actual []int, expected []int, t *testing.T)  {
	if ! reflect.DeepEqual(actual, expected) {
		t.Fatalf("Failed to confirm that the slices are equal (Expected: %v, Actual: %v)", expected, actual)
	}
}