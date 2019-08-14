package quicksort

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	A := []int{3,2,6,1,7,0,3,4,6}
	QuickSort(A)
	if !reflect.DeepEqual(A, []int{0,1,2,3,3,4,6,6,7}) {
		log.Fatal("排序出错了哦")
	}

	B := []int{3,4,3,3,6,3,3,3,3,6,3,3,4,3}
	QuickSort(B)
	if !reflect.DeepEqual(B, []int{3,3,3,3,3,3,3,3,3,3,4,4,6,6}) {
		log.Fatal("排序出错了哦")
	}
	fmt.Println(B)

	C := []int{3,3,3,3,3,3,3,3,3,3,3,3,3,3}
	QuickSort(C)
	if !reflect.DeepEqual(C, []int{3,3,3,3,3,3,3,3,3,3,3,3,3,3}) {
		log.Fatal("排序出错了哦")
	}
	fmt.Println(C)
}

func BenchmarkQuickSort(b *testing.B) {
	A := []int{3,2,6,1,7,0,3,4,6}
	b.ResetTimer()
	for n :=0; n < b.N; n++ {
		QuickSort(A)
	}
}
