package kmp

import (
	"fmt"
	"testing"
)

func TestGetNext(t *testing.T) {
	fmt.Println(GetNext("ABCABCDE"))
}


func TestKmpSearchOne(t *testing.T) {
	fmt.Println(KmpSearchOne("ABC", "BCBABC"))
}
