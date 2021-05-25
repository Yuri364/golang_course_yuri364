package slice

type SortedSlice struct {
    nums []int
}

// Copy
func (s *SortedSlice) Insert(num int)  {
    for i, v := range s.nums {
        if num < v {
            s.nums = append(s.nums, 0)
            copy(s.nums[i+1:], s.nums[i:])
            s.nums[i] = num
            return
        }
    }
    s.nums = append(s.nums, num)
}

// Append
func (s *SortedSlice) InsertV2(num int) {
    for i, v := range s.nums {
        if num < v {
            s.nums = append(s.nums[:i], append([]int{num}, s.nums[i:]...)...)

            return
        }
    }

    s.nums = append(s.nums, num)
}

func (s *SortedSlice) Delete(num int) {
   for i, v := range s.nums {
       if v == num {
           s.nums = append(s.nums[:i], s.nums[i+1:]...)
       }
   }
}