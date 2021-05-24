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
