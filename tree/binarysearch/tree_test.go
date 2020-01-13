package binarysearch

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	var root = &Node{Value:2}
	TreeInsert(root, &Node{Value:3})
	TreeInsert(root, &Node{Value:4})
	TreeInsert(root, &Node{Value:9})
	TreeInsert(root, &Node{Value:1})
	InorderTreeWalk(root)
	v := TreeSearch(root, 3)
	fmt.Println(v)
	TreeDelete(root, v)

	InorderTreeWalk(root)
}