## Slice: Insert, Delete, getMax, getMin

|                                                            |         |         |       |          |      |   |           |
|------------------------------------------------------------|---------|---------|-------|----------|------|---|-----------|
| BenchmarkSortedSlice_Insert_copy_hundreds-6                | 547023  | 2021    | ns/op | 16384    | B/op | 1 | allocs/op |
| BenchmarkSortedSlice_Insert_copy_thousand-6                | 126088  | 12056   | ns/op | 106496   | B/op | 1 | allocs/op |
| BenchmarkSortedSlice_Insert_copy_hundreds_of_thousands-6   | 877     | 1390835 | ns/op | 10002436 | B/op | 1 | allocs/op |
| BenchmarkSortedSlice_Insert_append_hundreds-6              | 367572  | 2757    | ns/op | 24576    | B/op | 2 | allocs/op |
| BenchmarkSortedSlice_Insert_append_thousand-6              | 83095   | 16091   | ns/op | 180224   | B/op | 2 | allocs/op |
| BenchmarkSortedSlice_Insert_append_hundreds_of_thousands-6 | 861     | 1518232 | ns/op | 17203201 | B/op | 2 | allocs/op |
| BenchmarkSortedSlice_Delete_hundreds-6                     | 4347541 | 274.8   | ns/op | 0        | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_Delete_thousand-6                     | 453878  | 2630    | ns/op | 0        | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_Delete_hundreds_of_thousands-6        | 2792    | 449264  | ns/op | 0        | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_getMax_hundreds-6                     | 2335897 | 513.6   | ns/op | 0        | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_getMax_thousand-6                     | 233172  | 5157    | ns/op | 0        | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_getMax_hundreds_of_thousands-6        | 1911    | 618086  | ns/op | 0        | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_getMin_hundreds-6                     | 2338509 | 508.4   | ns/op | 0        | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_getMin_thousand-6                     | 232839  | 5154    | ns/op | 0        | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_getMin_hundreds_of_thousands-6        | 1995    | 597354  | ns/op | 0        | B/op | 0 | allocs/op |

## Other sorting algorithms

|                                               |         |          |       |   |      |   |           |
|-----------------------------------------------|---------|----------|-------|---|------|---|-----------|
| BenchmarkSortedSlice_BubbleSort_hundreds-6    | 6955    | 171225   | ns/op | 0 | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_BubbleSort_thousand-6    | 66      | 17629371 | ns/op | 0 | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_InsertionSort_hundreds-6 | 1557603 | 778.6    | ns/op | 0 | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_InsertionSort_thousand-6 | 139779  | 7744     | ns/op | 0 | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_QuickSort_hundreds-6     | 126780  | 9400     | ns/op | 0 | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_QuickSort_thousand-6     | 10000   | 107671   | ns/op | 0 | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_SelectSort_hundreds-6    | 2490    | 475788   | ns/op | 0 | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_SelectSort_thousand-6    | 26      | 47267859 | ns/op | 0 | B/op | 0 | allocs/op |

