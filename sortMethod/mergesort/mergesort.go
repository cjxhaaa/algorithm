package mergesort

/*
	归并排序以O(NlogN)最坏运行时间运行，而所用的比较次数几乎是最优的。
	这个算法的基本操作是合并两个已排序的表，因为这两个表是已经排序的，所以若将输出放到第三个表中时则该算法可用过对输入数据一趟排序来完成。
    基本的合并算法是取两个输入数组A和B，一个输出数组C，以及三个计数器Aptr，Bptr，Cptr，它们初始对应数组的开始端。
	A[Aptr]和B[Bptr]中的较小者被拷贝到C中的下一个位置，相关计数器向前推进一步。当两个输入表有一个用完的时候，则将剩余部分直接copy到C中。
	合并两个已经排序的表明显是线性的，因为最多进行了N-1次比较，其中N是元素总数。

*/

func merge(A []int, l, m, r int) {
	//fmt.Println("start merge:")
	//fmt.Println("l:", l,"m:", m,"r:",r)
	var L []int
	var R []int
	for _, v := range A[l:m] {
		L = append(L, v)
	}
	//fmt.Println("L:", L)
	for _, v := range A[m:r] {
		R = append(R, v)
	}
	//fmt.Println("R:", R)

	var i = 0
	var j = 0

	for k:= l;k<r;k+=1 {
		//fmt.Println("i:",i,"j:",j,"k:",k)
		if i >= len(L) {
			for j < len(R) {

				A[k] = R[j]
				j+=1
				k+=1
			}
			break
		} else if j >= len(R) {
			for i < len(L) {
				A[k] = L[i]
				i+=1
				k+=1
			}
			break
		} else if L[i] <= R[j] {
			A[k] = L[i]
			i+=1
		} else {
			A[k] = R[j]
			j+=1
		}
	}
	//fmt.Println(A)
}

func MergeSort(A []int, l, r int) {
	if l < r && r-l > 1 {
		m := (l + r) /2
		MergeSort(A, l , m)
		MergeSort(A, m, r)
		merge(A, l, m, r)
	}
}



