package slice

type SortedSlice struct {
    nums []int
}

// Insert copy
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

// InsertV2 append
func (s *SortedSlice) InsertV2(num int) {
    for i, v := range s.nums {
        if num < v {
            s.nums = append(s.nums[:i], append([]int{num}, s.nums[i:]...)...)

            return
        }
    }

    s.nums = append(s.nums, num)
}


func (s *SortedSlice) getMax() int {
    max := s.nums[0]
    for _, v := range s.nums {
        if v > max {
            max = v
        }
    }

    return max
}

func (s *SortedSlice) getMin() int {
    min := s.nums[0]
    for _, v := range s.nums {
        if v < min {
            min = v
        }
    }

    return min
}

func (s *SortedSlice) Delete(num int) {
   for i, v := range s.nums {
       if v == num {
           s.nums = append(s.nums[:i], s.nums[i+1:]...)
       }
   }
}