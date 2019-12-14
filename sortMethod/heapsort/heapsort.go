package heapsort


/*
	类似基于二叉堆使用优先队列，建立N个元素二叉堆花费O(N)，执行N次DeleteMin，每次DeleteMin需要O(logN)，所以总运行时间O(NlogN)。
	但是它使用了一个附加数组，但是因此存储需求增加一倍。（数组拷贝额外时间消耗为O(N),不会显著影响运行时间，所以这是空间问题）

	为避免第二个数组不必要的空间开销，在每次DeleteMin之后，堆缩小了1。因此，位于堆中最后的单元可以用来存放刚刚删去的元素。
	使用这种策略，在最后一次DeleteMin后，如果以递减的顺序包含这些元素，这就是max堆。
	算法策略：先以线性时间建立一个堆，然后通过将堆中的最后元素与第一个元素交换，缩减堆的大小并进行下滤，来执行N - 1次DeleteMax操作。
*/
func HeapSort(A []int)  {
	var i int
	// 堆初始化，遍历二叉树的每个父节点进行堆序化
	for i = len(A)/2; i >= 0; i-- {
		percDown(A, i, len(A))
	}
	// 将堆顶最大值放到堆尾，重新对顶部元素进行堆序操作
	for i = len(A) - 1; i > 0; i-- {
		A[0], A[i] = A[i], A[0]
		percDown(A, 0, i)
	}
}


// 这里i是二叉树父节点的位置，child是父节点i的左子节点，child+1是右子节点
func percDown(A []int, i, N int)  {
	var child int
	var tmp int

	for tmp = A[i]; leftChild(i) < N; i = child {
		child = leftChild(i)

		// !=N-1 是为了判断是否尾部的右子节点是否存在，这里就是看看左右子节点哪个比较大
		if (child != N - 1 && A[child] < A[child+1]) {
			child++
		}
		// 这里判断较大的子节点与父节点的大小
		if tmp < A[child] {
			A[i] = A[child]
		} else {
			break
		}
	}
	A[i] = tmp
}

func leftChild(i int) int  {
	return 2 * i + 1
}