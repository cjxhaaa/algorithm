# 排序算法比较
## 100W数据

### heapsort
```$xslt
goos: darwin
goarch: amd64
pkg: algorithm/sortMethod/heapsort
BenchmarkHeapsort-4   	      10	 110767091 ns/op
PASS
```

### shellsort
```$xslt
goos: darwin
goarch: amd64
pkg: algorithm/sortMethod/shellsort
BenchmarkInsertSort-4   	      20	  54952254 ns/op
PASS
```

### mergesort
```$xslt
goos: darwin
goarch: amd64
pkg: algorithm/sortMethod/mergesort
BenchmarkMergeSort-4   	      10	 152582040 ns/op
PASS

```

### quicksort
```$xslt
goos: darwin
goarch: amd64
pkg: algorithm/sortMethod/quicksort
BenchmarkQuickSort-4   	      20	  56103500 ns/op
PASS
```

### insertsort
```$xslt
跑了半天没跑出来，看来O(n²)的算法在排序大数据时性能确实不行
```

