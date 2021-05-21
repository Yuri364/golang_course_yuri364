package main

import (
	"fmt"
)

func main() {

	var nums []int
	var num int

	for {
		fmt.Scan(&num)

		if num > 0 {
			nums = InsertV2(nums, num)
		} else {
			nums = Delete(nums, -num)
		}

		fmt.Println(nums)
	}
}

func Insert(nums []int, num int) []int {

	for i, v := range nums {
		if num < v {
			return append(nums[:i], append([]int{num}, nums[i:]...)...)
		}
	}

	return append(nums, num)
}

func InsertV2(nums []int, num int) []int {

	for i, v := range nums {
		if num < v {
			nums = append(nums, 0)
			copy(nums[i+1:], nums[i:])
			nums[i] = num
			return nums
		}
	}

	return append(nums, num)
}


func Delete(nums []int, num int) []int {
	for i, v := range nums {
		if v == num {
			return append(nums[:i], nums[i+1:]...)
		}
	}

	return nums
}
