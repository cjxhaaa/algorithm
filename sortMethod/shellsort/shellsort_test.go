package shellsort

import (
	"log"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestInsertSort(t *testing.T) {
	A := []int{3,2,6,1,7,0,3,4,6}
	ShellSort(A)
	if !reflect.DeepEqual(A, []int{0,1,2,3,3,4,6,6,7}) {
		log.Fatal("排序出错了哦")
	}
}

var test_num = 1000000

func BenchmarkInsertSort(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	num := test_num
	A := make([]int, 0, num)
	for i := 0; i < num; i++ {
		val := rand.Intn(200000000)
		A = append(A, val)
	}
	b.ResetTimer()
	for n :=0; n < b.N; n++ {
		ShellSort(A)
	}
}
