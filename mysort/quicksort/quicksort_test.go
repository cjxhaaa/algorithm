package quicksort

import (
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"testing"
	"time"
)


func TestQuickSort2(t *testing.T) {
	A := []int{3,2,6,1,7,0,3,4,6}
	QuickSort(A)
	if !reflect.DeepEqual(A, []int{0,1,2,3,3,4,6,6,7}) {
		fmt.Println(A)
		log.Fatal("A排序出错了哦")
	}
	fmt.Println(A)

	B := []int{3,4,3,3,6,3,3,3,3,6,3,3,4,3}
	QuickSort(B)
	if !reflect.DeepEqual(B, []int{3,3,3,3,3,3,3,3,3,3,4,4,6,6}) {
		fmt.Println(B)
		log.Fatal("B排序出错了哦")
	}
	fmt.Println(B)

	C := []int{3,3,3,3,3,3,3,3,3,3,3,3,3,3}
	QuickSort(C)
	if !reflect.DeepEqual(C, []int{3,3,3,3,3,3,3,3,3,3,3,3,3,3}) {
		fmt.Println(C)
		log.Fatal("C排序出错了哦")
	}
	fmt.Println(C)

	A = []int{3,2,6,1,7,0,3,4,6}
	QuickSort_GO(A)
	if !reflect.DeepEqual(A, []int{0,1,2,3,3,4,6,6,7}) {
		fmt.Println(A)
		log.Fatal("A排序出错了哦")
	}
	fmt.Println(A)

	B = []int{3,4,3,3,6,3,3,3,3,6,3,3,4,3}
	QuickSort_GO(B)
	if !reflect.DeepEqual(B, []int{3,3,3,3,3,3,3,3,3,3,4,4,6,6}) {
		fmt.Println(B)
		log.Fatal("B排序出错了哦")
	}
	fmt.Println(B)

	C = []int{3,3,3,3,3,3,3,3,3,3,3,3,3,3}
	QuickSort_GO(C)
	if !reflect.DeepEqual(C, []int{3,3,3,3,3,3,3,3,3,3,3,3,3,3}) {
		fmt.Println(C)
		log.Fatal("C排序出错了哦")
	}
	fmt.Println(C)
}

func TestQuickSort(t *testing.T) {




	rand.Seed(time.Now().UnixNano())
	num := 100000000
	testData1 := make([]int, 0, num)
	for i := 0; i < num; i++ {
		val := rand.Intn(200000000)
		testData1 = append(testData1, val)
	}

	start := time.Now()
	QuickSort_GO(testData1)
	fmt.Println("multi goroutine:", time.Now().Sub(start))
}

func BenchmarkQuickSort(b *testing.B) {
	A := []int{3,2,6,1,7,0,3,4,6}
	b.ResetTimer()
	for n :=0; n < b.N; n++ {
		QuickSort(A)
	}
}
