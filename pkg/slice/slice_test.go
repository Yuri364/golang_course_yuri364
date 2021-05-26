package slice

import (
    "math/rand"
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
        {[]int{512,632}, 111, []int{111,512,632}},
    }

    for _, item := range testData {
        s := SortedSlice{item.nums}

        s.InsertV2(item.num)

        if ! reflect.DeepEqual(s.nums, item.expected) {
            t.Fatalf("Failed to confirm that the slices are equal (Expected: %v, Actual: %v)", item.expected, s.nums)
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
            t.Fatalf("Failed to confirm that the slices are equal (Expected: %v, Actual: %v)", item.expected, s.nums)
        }
    }
}

func TestGetMax(t *testing.T) {
    var testData = []struct {
        nums     []int
        expected int
    } {
        {[]int{3,32,10,35,54,76}, 76},
        {[]int{343,232,410,776,535,654}, 776},
        {[]int{3443,7576,2632,4710,5535,6754}, 7576},
    }

    for _, item := range testData {
        s := SortedSlice{item.nums}
        r := s.getMax()

        if ! reflect.DeepEqual(r, item.expected) {
            t.Fatalf("Failed getMax test (Expected: %v, Actual: %v)", item.expected, r)
        }
    }
}

func TestGetMin(t *testing.T) {
    var testData = []struct {
        nums     []int
        expected int
    } {
        {[]int{3,32,10,35,54,76}, 3},
        {[]int{343,232,410,776,535,654}, 232},
        {[]int{3443,7576,2632,4710,5535,6754}, 2632},
    }

    for _, item := range testData {
        s := SortedSlice{item.nums}
        r := s.getMin()

        if ! reflect.DeepEqual(r, item.expected) {
            t.Fatalf("Failed getMin test (Expected: %v, Actual: %v)", item.expected, r)
        }
    }
}

// Benchmarks
func BenchmarkSortedSlice_Insert_copy_hundreds(b *testing.B) { benchmarkSortedSliceInsert(b, 999, 99) }
func BenchmarkSortedSlice_Insert_copy_thousand(b *testing.B) { benchmarkSortedSliceInsert(b, 9999, 999) }
func BenchmarkSortedSlice_Insert_copy_hundreds_of_thousands(b *testing.B) { benchmarkSortedSliceInsert(b, 999999, 99999) }
func BenchmarkSortedSlice_Insert_copy_million(b *testing.B) { benchmarkSortedSliceInsert(b, 1000000, 999999) }

func BenchmarkSortedSlice_Insert_append_hundreds(b *testing.B) { benchmarkSortedSliceInsertV2(b, 999, 99) }
func BenchmarkSortedSlice_Insert_append_thousand(b *testing.B) { benchmarkSortedSliceInsertV2(b, 9999, 999) }
func BenchmarkSortedSlice_Insert_append_hundreds_of_thousands(b *testing.B) { benchmarkSortedSliceInsertV2(b, 999999, 99999) }
func BenchmarkSortedSlice_Insert_append_million(b *testing.B) { benchmarkSortedSliceInsertV2(b, 1000000, 999999) }

func BenchmarkSortedSlice_Delete_hundreds(b *testing.B) { benchmarkSortedSliceDelete(b, 999, 99) }
func BenchmarkSortedSlice_Delete_thousand(b *testing.B) { benchmarkSortedSliceDelete(b, 9999, 999) }
func BenchmarkSortedSlice_Delete_hundreds_of_thousands(b *testing.B) { benchmarkSortedSliceDelete(b, 999999, 99999) }

func BenchmarkSortedSlice_getMax_hundreds(b *testing.B) { benchmarkSortedSliceGetMax(b, 999) }
func BenchmarkSortedSlice_getMax_thousand(b *testing.B) { benchmarkSortedSliceGetMax(b, 9999) }
func BenchmarkSortedSlice_getMax_hundreds_of_thousands(b *testing.B) { benchmarkSortedSliceGetMax(b, 999999) }

func BenchmarkSortedSlice_getMin_hundreds(b *testing.B) { benchmarkSortedSliceGetMin(b, 999) }
func BenchmarkSortedSlice_getMin_thousand(b *testing.B) { benchmarkSortedSliceGetMin(b, 9999) }
func BenchmarkSortedSlice_getMin_hundreds_of_thousands(b *testing.B) { benchmarkSortedSliceGetMin(b, 999999) }

func benchmarkSortedSliceInsert(b *testing.B, size int, num int) {
    testSlice := getTestSlice(size)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        s := SortedSlice{nums: testSlice}

        s.Insert(num)
    }
}

func benchmarkSortedSliceInsertV2(b *testing.B, size int, num int) {
    testSlice := getTestSlice(size)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        s := SortedSlice{nums: testSlice}

        s.InsertV2(num)
    }
}

func benchmarkSortedSliceDelete(b *testing.B, size int, num int) {
    testSlice := getTestSlice(size)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        s := SortedSlice{nums: testSlice}

        s.Delete(num)
    }
}

func benchmarkSortedSliceGetMax(b *testing.B, size int) {
    testSlice := getTestSlice(size)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        s := SortedSlice{nums: testSlice}

        s.getMax()
    }
}

func benchmarkSortedSliceGetMin(b *testing.B, size int) {
    testSlice := getTestSlice(size)

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        s := SortedSlice{nums: testSlice}

        s.getMin()
    }
}

// Preallocate
func getTestSlice(size int) []int {
    slice := make([]int, 0, size)

    for i := 0; i < size; i++ {
        slice = append(slice, i)
    }

    rand.Shuffle(len(slice), func(i, j int) { slice[i], slice[j] = slice[j], slice[i] })

    return slice
}