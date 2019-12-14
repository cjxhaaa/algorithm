package shellsort


/*
	条件:一个数组A,长度为N
	又称缩小增量排序，不像插入排序依次比较，而是人为划分间隔比较，各趟比较的距离随着算法的进行而减小。
	示例的间隔策略是先取N/2的间隔，后面依次整除2
	[]int{3,2,6,1,7,0,3,4,6}

	《数据结构与算法分析--c语言描述》
	间隔策略实践中最优解为Sedgewick的9*4^i - 9*2^i + 1 或者 4^i - 3*2^i + 1的数列项
	其中还有常用的Hibbard的解1,3,7,...,2^k - 1。 其最坏运行时间θ(N^(3/2))
*/

func ShellSort(A []int) {
	for increment := len(A) / 2; increment > 0; increment /= 2 {
		for i := increment; i < len(A); i++ {
			tmp := A[i]
			var j int
			for j = i; j >= increment; j -= increment {
				if tmp < A[j - increment] {
					A[j] = A[j - increment]
				} else {
					break
				}
			}
			A[j] = tmp
		}
	}
}