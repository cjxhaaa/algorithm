package quicksort

import (
	"algorithm/mysort/insertsort"
)

const cutOff = 3

/*
	类似归并排序，快速排序也是分治递归。
	方法：
	>1.如果数组A元素长度是0或1，则返回
	>2.取A中任一元素v，称为枢纽元pivot
	>3.将A基于v分成两个不想交的集合A1,A2
	>4.对qsort(A1),qsort(A2),goto:2

	快排对选取枢纽元有一定策略，一般来说随机选取，不要直接取第一个。还有就是三数中值分割法（随机取三个数，取中间值）

	分割策略：
	>1.将枢纽元p与最后的元素交换使得枢纽元p离开要被分割的数据段。
	>2.i从第一个元素开始，j从倒数第二个元素开始

	分割截断要做的是把所有小元素移到数组的左边而把所有大元素移到数组的右边，当然大小是相对枢纽元p。
	当i在j的左边时，将i右移，移过小于p的元素；当j在i的右边时，将j左移，移过大于p的元素。
	当移动停止时，i指向一个大元素（大于p），j指向一个小元素（小与p），如果i在j左边，那么将两个元素互换。
	具体步骤示例：
	>1.i右移直到遇到一个大于p的元素，j左移直到遇到一个小与p的元素，这时若i,j未相遇，将i,j指向的数互换。
	>2.重复步骤1，最后i，j相遇时，可知已将小与p的元素与大于p的元素分开
	>3.对分开后的两个数组继续先取一个p，重复1,2,递归实现排序，到每个排序的数组len为0或1时，结束排序。

*/
func QuickSort(A []int) {
	qsort(A, 0, len(A) - 1)
}

func QuickSort_GO(A []int) {
	done := make(chan struct{})
	go qsort_go(A, 0, len(A) - 1, done, 5)
	<-done
}

func QuickSort2(A []int) {
	qsort2(A, 0, len(A) - 1)
}

func qsort(A []int, left, right int) {
	if len(A) <= 1 {
		return
	}

	if right - left >= cutOff {
		p := median3(A, left, right)
		pv := A[p]
		//i := arrange(A, left+1, right - 1, p)
		i,j := left, right
		for i <= j {
			for j >= p && A[j] > pv {
				j--
			}
			if j >= p {
				A[p] = A[j]
				p = j
			}

			for i <= p && A[i] <= pv {
				i++
			}
			if i <= p {
				A[p] = A[i]
				p = i
			}
		}
		A[p] = pv
		//A[i], A[right-1] = A[right-1], A[i]   // 将i, j 相交点与p交换，这时数组A中，A[i] 左侧的数都比p小，右侧数都比p大
		if p - left > 1 {
			qsort(A, left, p - 1)
		}

		if right - p > 1 {
			qsort(A, p + 1, right)
		}


	} else {  // 子集数量太少使用插入排序优化
		insertsort.InsertSort(A[left:], right-left+1)
	}

}

/*
	三值分割法。
	取三元素A[left],A[center],A[right]
	将三元素排序后，取A[center]的值为枢纽元，目前已知条件是A[left]<A[center]<A[right]
	所以将A[center]的值与A[right-1]交换，将i,j初始化为left+1,right-2
*/
func median3(A []int, left,right int) int  {
	center := (left + right) / 2

	if A[left] > A[center] {
		A[left], A[center] = A[center], A[left]
	}
	if A[left] > A[right] {
		A[left], A[right] = A[right], A[left]
	}
	if A[center] > A[right] {
		A[center], A[right] = A[right], A[center]
	}

	//A[center], A[right-1] = A[right-1], A[center]
	return center
}

/*
	并发的快排
 */
func qsort_go(A []int, left, right int, done chan struct{}, depth int) {
	if len(A) <= 1 {
		done <- struct{}{}
		return
	}

	if right - left >= cutOff {
		depth--
		p := median3(A, left, right)
		pv := A[p]
		//i := arrange(A, left+1, right - 1, p)
		i,j := left, right
		for i <= j {
			for j >= p && A[j] > pv {
				j--
			}
			if j >= p {
				A[p] = A[j]
				p = j
			}

			for i <= p && A[i] <= pv {
				i++
			}
			if i <= p {
				A[p] = A[i]
				p = i
			}
		}
		A[p] = pv
		//A[i], A[right-1] = A[right-1], A[i]   // 将i, j 相交点与p交换，这时数组A中，A[i] 左侧的数都比p小，右侧数都比p大

		if  depth > 0 {
			childDone := make(chan struct{}, 2)
			if p - left > 1 {
				go qsort_go(A, left, i-1, childDone, depth)
			} else {
				childDone <- struct{}{}
			}
			if right - p > 1 {
				go qsort_go(A, i+1, right, childDone, depth)
			} else {
				childDone <- struct{}{}
			}
			<- childDone
			<- childDone
		} else {
			if p - left > 1 {
				qsort(A, left, p - 1)
			}

			if right - p > 1 {
				qsort(A, p + 1, right)
			}
		}

	} else {
		insertsort.InsertSort(A[left:], right-left+1)
	}
	done <- struct{}{}
}

func qsort2(A []int, left, right int) {
	if len(A) <= 1 {
		return
	}

	if right - left >= cutOff {
		p := partition(A, left, right)
		qsort2(A, left, p-1)
		qsort2(A, p+1, right)

	} else {  // 子集数量太少使用插入排序优化
		insertsort.InsertSort(A[left:], right-left+1)
	}

}

func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {   // 这里其实就是个偷懒的冒泡，每次把两个大的元素往后调，垃圾
		if a[j] < pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}