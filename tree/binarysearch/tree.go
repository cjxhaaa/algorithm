package binarysearch

import "fmt"

// 参考《算法导论》第12章

type Node struct {
	Left   *Node
	Right  *Node
	P      *Node
	Value  int
}

// 先序遍历
func PreorderTreeWalk(x *Node) {
	if x != nil {
		fmt.Println(x.Value)
		PreorderTreeWalk(x.Left)
		PreorderTreeWalk(x.Right)
	}
}


// 中序遍历
func InorderTreeWalk(x *Node) {
	if x != nil {
		InorderTreeWalk(x.Left)
		fmt.Println(x.Value)
		InorderTreeWalk(x.Right)
	}
}

// 后序遍历
func PostorderTreeWalk(x *Node) {
	if x != nil {
		PostorderTreeWalk(x.Left)
		PostorderTreeWalk(x.Right)
		fmt.Println(x.Value)
	}
}


// 查询递归版本，运行时间为O(h), h为树高度
func TreeSearch(x *Node, k int) *Node {
	if x == nil || k == x.Value {
		return x
	}
	if k < x.Value {
		return TreeSearch(x.Left, k)
	} else {
		return TreeSearch(x.Right, k)
	}
}

// 查找的迭代版本
func IterativeTreeSearch(x *Node, k int) *Node {
	for x != nil && k != x.Value {
		if k < x.Value {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	return x
}

func TreeMininum(x *Node) *Node {
	if x.Left != nil {
		return TreeMininum(x.Left)
	}
	return x
}

func IterativeTreeMininum(x *Node) *Node {
	for x.Left != nil {
		x = x.Left
	}
	return x
}

func TreeMaxinum(x *Node) *Node {
	if x.Right != nil {
		return TreeMaxinum(x.Left)
	}
	return x
}

func IterativeTreeMaxinum(x *Node) *Node {
	for x.Right != nil {
		x = x.Right
	}
	return x
}


// 按中序遍历查找后继节点。分两种情况：
// 1.如果节点x的右子树非空，那么x的后继恰好是x右子树的最左节点，直接通过TreeMininum(x.Right)找到。
// 2.如果节点x的右子树为空并有一个后继y，那么y就是x的最底层祖先，并且y的左孩子也是x的一个祖先。
func TreeSuccessor(x *Node) *Node {
	if x.Right != nil {
		return TreeMininum(x.Right)
	}
	y := x.P
	for y != nil && x == y.Right {
		x = y
		y = y.P
	}
	return y
}

// 查找前驱节点,与查后继镜像
func TreePredecessor(x *Node) *Node {
	if x.Left != nil {
		return TreeMaxinum(x.Left)
	}
	y := x.P
	for y != nil && x == y.Left {
		x = y
		y = y.P
	}
	return y
}

// 遍历查找插入位置，y始终为x的父节点
func TreeInsert(root, z *Node) {
	var x = root
	var y *Node

	for x != nil {
		y = x
		if z.Value < x.Value {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	z.P = y
	if y == nil {
		root = z
	} else {
		if z.Value < y.Value {
			y.Left = z
		} else {
			y.Right = z
		}
	}
}


// 将以u为根的子树替换为以v为根的子树
func transplant(root, u, v *Node) {
	if u.P == nil {   // u为树根的情况
		root = v
	} else if u == u.P.Left {  // 左孩子的情况
		u.P.Left = v
	} else {             // 右孩子的情况
		u.P.Right = v
	}
	if v != nil {
		v.P = u.P
	}
}

func TreeDelete(root, z *Node) {
	if z.Left == nil {                 // 没有左孩子，直接替换
		transplant(root, z, z.Right)
	} else if z.Right == nil {         // 没有右孩子，直接替换
		transplant(root, z, z.Left)
	} else {                             // 左右孩子都存在，找到z的后继节点来替换z
		y := TreeMininum(z.Right)        // 因为有右孩子，则后继y一定是右子树的最小节点, 而且y没有左孩子，所以将y直接移出来拼接
		if y.P != z {                    // 当y不是z的右孩子时，因为y没有左孩子，将y的右孩子替换y，将y移至z右子树的根节点
			transplant(root, y, y.Right)
			y.Right = z.Right
			y.Right.P = y
		}
		transplant(root, z, y)          // 现在直接将y替换为z
		y.Left = z.Left
		y.Left.P = y
	}
}

