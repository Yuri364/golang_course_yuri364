goos: linux
goarch: amd64
pkg: app/course/golang_course_yuri364/pkg/slice
cpu: Intel(R) Core(TM) i5-9400F CPU @ 2.90GHz
BenchmarkSortedSlice_Insert_copy_hundreds-6                  	  576160	      2029 ns/op	   16384 B/op	       1 allocs/op
BenchmarkSortedSlice_Insert_copy_thousand-6                  	  129798	     11991 ns/op	  106496 B/op	       1 allocs/op
BenchmarkSortedSlice_Insert_copy_hundreds_of_thousands-6     	     912	   1426470 ns/op	10002438 B/op	       1 allocs/op
BenchmarkSortedSlice_Insert_append_hundreds-6                	  455444	      3048 ns/op	   24576 B/op	       2 allocs/op
BenchmarkSortedSlice_Insert_append_thousand-6                	   75340	     14894 ns/op	  180224 B/op	       2 allocs/op
BenchmarkSortedSlice_Insert_append_hundreds_of_thousands-6   	     765	   1686915 ns/op	17203202 B/op	       2 allocs/op
BenchmarkSortedSlice_Delete_hundreds-6                       	 4202362	       274.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkSortedSlice_Delete_thousand-6                       	  448838	      2593 ns/op	       0 B/op	       0 allocs/op
BenchmarkSortedSlice_Delete_hundreds_of_thousands-6          	    2559	    467441 ns/op	       0 B/op	       0 allocs/op
BenchmarkSortedSlice_getMax_hundreds-6                       	 4364847	       273.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkSortedSlice_getMax_thousand-6                       	  455940	      2586 ns/op	       0 B/op	       0 allocs/op
BenchmarkSortedSlice_getMax_hundreds_of_thousands-6          	    2748	    448255 ns/op	       0 B/op	       0 allocs/op
BenchmarkSortedSlice_getMin_hundreds-6                       	 4377116	       272.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkSortedSlice_getMin_thousand-6                       	  455468	      2593 ns/op	       0 B/op	       0 allocs/op
BenchmarkSortedSlice_getMin_hundreds_of_thousands-6          	    2748	    443216 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	app/course/golang_course_yuri364/pkg/slice	20.241s
