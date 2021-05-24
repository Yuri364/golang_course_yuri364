package slice

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

    for _, item := range testData {
        s := SortedSlice{item.nums}

        s.Insert(item.num)

        if ! reflect.DeepEqual(s.nums, item.expected) {
            t.Fatalf("Failed to confirm that the slices are equal (Expected: %v, Result: %v)", item.expected, s.nums)
        }
    }
}

func TestInsertV2(t *testing.T) {
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

    for _, item := range testData {
        s := SortedSlice{item.nums}

        r := s.InsertV2(item.num)

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
        s := SortedSlice{item.nums}

        s.Delete(item.num)

        if ! reflect.DeepEqual(s.nums, item.expected) {
            t.Fatalf("Failed to confirm that the slices are equal (Expected: %v, Result: %v)", item.expected, s.nums)
        }
    }
}

func BenchmarkSortedSlice_Insert(b *testing.B) { benchmarkSortedSliceInsert(b, 10, 9) }
func BenchmarkSortedSlice_Insert2(b *testing.B) { benchmarkSortedSliceInsert(b, 100, 99) }
func BenchmarkSortedSlice_Insert3(b *testing.B) { benchmarkSortedSliceInsert(b, 1000, 999) }
func BenchmarkSortedSlice_Insert4(b *testing.B) { benchmarkSortedSliceInsert(b, 10000, 9999) }
func BenchmarkSortedSlice_Insert5(b *testing.B) { benchmarkSortedSliceInsert(b, 100000, 99999) }

func BenchmarkSortedSlice_InsertV2(b *testing.B) { benchmarkSortedSliceInsertV2(b, 10, 9) }
func BenchmarkSortedSlice_InsertV2_2(b *testing.B) { benchmarkSortedSliceInsertV2(b, 100, 99) }
func BenchmarkSortedSlice_InsertV2_3(b *testing.B) { benchmarkSortedSliceInsertV2(b, 1000, 999) }
func BenchmarkSortedSlice_InsertV2_4(b *testing.B) { benchmarkSortedSliceInsertV2(b, 10000, 9999) }
func BenchmarkSortedSlice_InsertV2_5(b *testing.B) { benchmarkSortedSliceInsertV2(b, 100000, 99999) }

func BenchmarkSortedSlice_Delete(b *testing.B) { benchmarkSortedSliceDelete(b, 10, 9) }
func BenchmarkSortedSlice_Delete2(b *testing.B) { benchmarkSortedSliceDelete(b, 100, 99) }
func BenchmarkSortedSlice_Delete3(b *testing.B) { benchmarkSortedSliceDelete(b, 1000, 999) }
func BenchmarkSortedSlice_Delete4(b *testing.B) { benchmarkSortedSliceDelete(b, 10000, 9999) }
func BenchmarkSortedSlice_Delete5(b *testing.B) { benchmarkSortedSliceDelete(b, 100000, 99999) }


func benchmarkSortedSliceInsert(b *testing.B, capacity int, num int) {
    testSlice := getTestSlice(capacity)
    for i := 0; i < b.N; i++ {
        s := SortedSlice{nums: testSlice}

        s.Insert(num)
    }
}

func benchmarkSortedSliceInsertV2(b *testing.B, capacity int, num int) {
    testSlice := getTestSlice(capacity)
    for i := 0; i < b.N; i++ {
        s := SortedSlice{nums: testSlice}

        s.InsertV2(num)
    }
}

func benchmarkSortedSliceDelete(b *testing.B, capacity int, num int) {
    testSlice := getTestSlice(capacity)
    for i := 0; i < b.N; i++ {
        s := SortedSlice{nums: testSlice}

        s.Delete(num)
    }
}

func getTestSlice(capacity int) []int {
    slice := make([]int, 0, capacity)

    for i := 0; i < capacity; i++ {
        slice = append(slice, i)
    }

    return slice
}