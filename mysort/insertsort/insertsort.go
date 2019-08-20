package insertsort

/*
	插入排序的逻辑:
	条件:一个数组A,长度为N,需要排序N-1趟（即除了数组首位，要对后面N-1个元素进行顺序检查）
	每趟排序轮到的那个元素与前面已经排好序的元素反向依次比对，将元素插入到合适位置
	排序次数 1 + 2 + 3 + ... + (N - 1) = (1 + N - 1)*(N - 1)/2 = N*(N-1)/2 = θ(N²)
	如果已经预先排序，运行时间为O(N)
*/
func InsertSort(A []int, N int) {
	for p := 1; p < N; p++ {
		tmp := A[p]
		var j int
		for  j = p; j > 0 && A[j - 1] > tmp; j-- {
			A[j] = A[j-1]
		}
		A[j] = tmp
	}
}
