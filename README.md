## Slice: Insert, Delete, getMax, getMin

|                                                            |         |         |       |          |      |   |           |
|------------------------------------------------------------|---------|---------|-------|----------|------|---|-----------|
| BenchmarkSortedSlice_Insert_copy_hundreds-6                | 613929  | 1842    | ns/op | 16384    | B/op | 1 | allocs/op |
| BenchmarkSortedSlice_Insert_copy_thousand-6                | 124434  | 9351    | ns/op | 106496   | B/op | 1 | allocs/op |
| BenchmarkSortedSlice_Insert_copy_hundreds_of_thousands-6   | 774     | 1429923 | ns/op | 10002436 | B/op | 1 | allocs/op |
| BenchmarkSortedSlice_Insert_copy_million-6                 | 980     | 1178440 | ns/op | 10002432 | B/op | 1 | allocs/op |
| BenchmarkSortedSlice_Insert_append_hundreds-6              | 356404  | 2839    | ns/op | 24576    | B/op | 2 | allocs/op |
| BenchmarkSortedSlice_Insert_append_thousand-6              | 66414   | 16903   | ns/op | 188416   | B/op | 2 | allocs/op |
| BenchmarkSortedSlice_Insert_append_hundreds_of_thousands-6 | 716     | 1450990 | ns/op | 18006021 | B/op | 2 | allocs/op |
| BenchmarkSortedSlice_Insert_append_million-6               | 912     | 1179616 | ns/op | 10002432 | B/op | 1 | allocs/op |
| BenchmarkSortedSlice_Delete_hundreds-6                     | 4243228 | 273.5   | ns/op | 0        | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_Delete_thousand-6                     | 461283  | 2597    | ns/op | 0        | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_Delete_hundreds_of_thousands-6        | 2691    | 446737  | ns/op | 0        | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_getMax_hundreds-6                     | 2346525 | 510.4   | ns/op | 0        | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_getMax_thousand-6                     | 230547  | 5151    | ns/op | 0        | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_getMax_hundreds_of_thousands-6        | 1898    | 623008  | ns/op | 0        | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_getMin_hundreds-6                     | 2336359 | 513.3   | ns/op | 0        | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_getMin_thousand-6                     | 227737  | 5195    | ns/op | 0        | B/op | 0 | allocs/op |
| BenchmarkSortedSlice_getMin_hundreds_of_thousands-6        | 1929    | 621652  | ns/op | 0        | B/op | 0 | allocs/op |

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

