package insertsort

import (
	"log"
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	A := []int{3,2,6,1,7,0,3,4,6}
	InsertSort(A, len(A))
	if !reflect.DeepEqual(A, []int{0,1,2,3,3,4,6,6,7}) {
		log.Fatal("排序出错了哦")
	}
}

func BenchmarkInsertSort(b *testing.B) {
	A := []int{3,2,6,1,7,0,3,4,6}
	b.ResetTimer()
	for n :=0; n < b.N; n++ {
		InsertSort(A, len(A))
	}
}
