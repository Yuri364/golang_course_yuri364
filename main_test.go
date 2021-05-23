package main

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	var testData = []struct {
		nums     []int
		num      int
		expected []int
	} {
		{[]int{3}, 1, []int{1,3}},
		{[]int{1,3}, 2, []int{1,2,3}},
		{[]int{1,3,12}, 6, []int{1,3,6,12}},
		{[]int{5,6,10}, 11, []int{5,6,10,11}},
	}

	// Insert V1
	for _, item := range testData {
		r := Insert(item.nums, item.num)

		if ! reflect.DeepEqual(r, item.expected) {
			t.Fatalf("Failed to confirm that the slices are equal (Expected: %v, Result: %v)", item.expected, r)
		}
	}

	// Insert V2
	for _, item := range testData {
		r := InsertV2(item.nums, item.num)

		if ! reflect.DeepEqual(r, item.expected) {
			t.Fatalf("Failed to confirm that the slices are equal (Expected: %v, Result: %v)", item.expected, r)
		}
	}

}

func TestDelete(t *testing.T) {
	var testData = []struct {
		nums     []int
		num      int
		expected []int
	} {
		{[]int{3}, 3, []int{}},
		{[]int{1,3}, 1, []int{3}},
		{[]int{1,3,6,12}, 6, []int{1,3,12}},
		{[]int{5,6,10,11}, 11, []int{5,6,10}},
	}

	for _, item := range testData {
		r := Delete(item.nums, item.num)

		if ! reflect.DeepEqual(r, item.expected) {
			t.Fatalf("Failed to confirm that the slices are equal (Expected: %v, Result: %v)", item.expected, r)
		}
	}

}

func BenchmarkInsert(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Insert([]int{12,32,12},3)
    }
}

func BenchmarkInsert2(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Insert([]int{123,32,12},33)
    }
}

func BenchmarkDelete(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Delete([]int{123,23,432,3212,1,43}, 23)
    }
}