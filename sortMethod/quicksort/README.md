# algorithm


## quicksort Benchmark

1. qsort基本版
```$xslt
// 100个数
goos: darwin
goarch: amd64
pkg: algorithm/mysort/quicksort
BenchmarkQuickSort-4   	 1000000	      1113 ns/op
PASS
// 10000个数
goos: darwin
goarch: amd64
pkg: algorithm/mysort/quicksort
BenchmarkQuickSort-4   	   10000	    144329 ns/op
PASS
```

2.qsort代码简洁版
```$xslt
// 100个数
goos: darwin
goarch: amd64
pkg: algorithm/mysort/quicksort
BenchmarkQuickSort2-4   	  200000	      8618 ns/op
PASS
// 10000个数
goos: darwin
goarch: amd64
pkg: algorithm/mysort/quicksort
BenchmarkQuickSort2-4   	     100	  57582276 ns/op
PASS
```

3. qsort并发版
```$xslt
// 100个数
goos: darwin
goarch: amd64
pkg: algorithm/mysort/quicksort
BenchmarkQuickSort3-4   	  100000	     20526 ns/op
PASS
// 10000个数
goos: darwin
goarch: amd64
pkg: algorithm/mysort/quicksort
BenchmarkQuickSort3-4   	   10000	    119549 ns/op
PASS
```


综上对三种快排写法做了对比，如果数据量很大，可以选择并发排序，效率会高一些，
千万不要为了一时代码的简洁而使用方法2，其实就是里面套了个冒泡排序，被他简洁的代码给骗了，。。。