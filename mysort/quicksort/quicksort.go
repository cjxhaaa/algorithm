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

func qsort(A []int, left, right int) {
	if len(A) <= 1 {
		return
	}

	if right - left < cutOff {
		p := median3(A, left, right)
		i, j := left + 1, right - 2
		for {
			for A[i] < p {
				i++
			}
			for A[j] > p {
				j--
			}
			if i < j {
				A[i], A[j] = A[j], A[i]
			} else {
				break
			}

		}
		A[i], A[right - 1] = A[right - 1], A[i]
		qsort(A, left, i - 1)
		qsort(A, right + 1, right)

	} else {  // 子集数量太少使用插入排序优化
		insertsort.InsertSort(A[left: right-left + 1])
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

	A[center], A[right - 1] = A[right - 1], A[center]
	return A[right - 1]
}
