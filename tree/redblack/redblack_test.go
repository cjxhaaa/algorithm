package redblack

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	var rbTree = &Tree{Root:Nil, Nil:Nil}

	rbTree.RBInsert(&Node{Key:2})
	rbTree.RBInsert(&Node{Key:3})

	rbTree.RBInsert(&Node{Key:4})

	rbTree.RBInsert(&Node{Key:9})

	dd := rbTree.RBInsert(&Node{Key:1})
	fmt.Println(dd)

	rbTree.RBDelete(dd)

}
