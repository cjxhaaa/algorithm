package redblack

import "fmt"

// 参考《算法导论》第13章
/*
一颗红黑树是满足红黑性质的二叉搜索树：
1.每个结点或是红色，或是黑色
2.根结点是黑色
3.每个叶结点(nil)是黑色
4.如果一个结点是红色的，则它的两个子结点都是黑色
5.对每个结点，从该结点到其所有后代叶结点的简单路径上，均包含相同数目的黑色结点

** 使用一个哨兵T.nil代表所有叶节点和根节点的父节点

从某个结点x出发(不含该结点)到达一个叶结点的任意一条简单路径上的黑色结点个数称为该结点的 黑高，记bh(x)。
则红黑树的黑高为其根结点的黑高。

引理：一棵有n个内部结点的红黑树的高度h至多为2lg(n+1)

证明：
先证明任一结点x为根的子树中至少包含2^bh(x)-1个内部结点。

1.如果x的高度为0，则x必为叶结点T.nil，且以x为根结点的子树至少包含2^bh(x)-1 = 2^0-1 = 0个内部结点。成立。
2.考虑一个高度大于0且有两个子结点的内部结点x情况。当x为红色，该结点黑高bh(x)。当x为黑色，该结点黑高bh(x)-1。
由于x子结点的高度比x本身的高度要低，可以利用归纳假设得出x子结点至少有2^(bh(x)-1)-1个内部结点。
那么，x为根的子树至少包含两个子结点内部结点之和与本身，即2*(2^(bh(x)-1)-1) + 1 = 2^bh(x)-1个内部结点。成立

又根据红黑性质4，从根到叶结点任何简单路径上至少有一半为黑色。
设h为树高，则bh(x) >= h/2
n >= 2^bh(x)-1 >= 2^(h/2)-1
n >= 2^(h/2)-1
得h <= 2lg(n+1)

因为动态集合操作search,minium,maxinum,successor和predecessor在一棵高度为h的二叉搜索树上运行时间为O(h),
基于这个结论，容易知道红黑树的这些操作运行时间为O(lgn)。
多么平衡，多么棒的红黑树啊!!!

红黑树的insert和delete操作需要保持红黑性质，这是与二叉搜索树的主要区别。
 */

type Color int

const (
	Black Color = iota
	Red
)


type Node struct {
	Color  Color
	Key    int
	Left   *Node
	Right  *Node
	P      *Node
}

type Tree struct {
	Root *Node
	Nil  *Node
}

var Nil = &Node{Color:Black}

// 先序遍历
func PreorderTreeWalk(x *Node) {
	if x != nil {
		fmt.Println(x.Key,"*",x.Color)
		PreorderTreeWalk(x.Left)
		PreorderTreeWalk(x.Right)
	}
}


// 中序遍历
func InorderTreeWalk(x *Node) {
	if x != nil {
		InorderTreeWalk(x.Left)
		fmt.Println(x.Key)
		//fmt.Println(x.Key,"*",x.Color)
		InorderTreeWalk(x.Right)
	}
}

// 后序遍历
func PostorderTreeWalk(x *Node) {
	if x != nil {
		PostorderTreeWalk(x.Left)
		PostorderTreeWalk(x.Right)
		fmt.Println(x.Key,"*",x.Color)
	}
}

// 左旋
func (t *Tree)LeftRotate(x *Node) {
	y := x.Right       // 找到右结点y
	x.Right = y.Left
	if y.Left != t.Nil {
		y.Left.P = x
	}

	y.P = x.P          // 将y父结点变为原x父结点
	if x.P == t.Nil{     // x为根结点的情况
		t.Root = y
	} else {
		if x == x.P.Left {  // 如果x原来是左孩子
			x.P.Left = y
		} else {
			x.P.Right = y          // 如果x原来是右孩子
		}
	}
	y.Left = x                 // 将x变为y左孩子
	x.P = y
}

// 右旋
func (t *Tree)RightRotate(x *Node) {
	y := x.Left
	x.Left = y.Right
	if y.Right != t.Nil {

		y.Right.P = x
	}

	y.P = x.P
	if x.P == t.Nil {
		t.Root = y
	} else {
		if x == x.P.Left {
			x.P.Left = y
		} else {
			x.P.Right = y
		}
	}
	y.Right = x
	x.P = y
}

func (t *Tree)RBInsert(z *Node) {
	y := t.Nil
	x := t.Root
	for x != t.Nil {     // 正常的遍历找到结点该插入的位置
		y = x
		if z.Key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	z.P = y    // 找到合适的y后，y成为z的父亲
	if y == Nil {    // 如果y为空，即空树，则z直接是根结点
		t.Root = z
	} else {
		if z.Key < y.Key {    // 判断z是成为左孩子还是右孩子
			y.Left = z
		} else {
			y.Right = z
		}
	}
	z.Left = Nil    // 因为z是新插入的，他的孩子肯定都是空
	z.Right = Nil
	z.Color = Red   // 抹红，以满足红黑性质5。因为插入之前所有根至外部结点路径上黑色结点数目相同，如果插入黑色，肯定导致黑色数目不同。所以先将其置为红色，方便后续调整
	t.RBInsertFixup(z)
}


// 插入后因为多了一个红色结点，可能会破坏红黑性质，进行修复
func (t *Tree)RBInsertFixup(z *Node) {

	for z.P.Color == Red {  // 插入的z一开始已经被抹红了，所以要判断他破坏了哪些红黑性质，并进行相应修复

		if z.P == z.P.P.Left {   // 当z的父亲是z爷爷的左孩子时
			y := z.P.P.Right      // 找到z的叔叔结点y
			if y.Color == Red {     // 如果y为红色，将z父亲和叔叔抹黑，将z爷爷抹红。保持红黑性质的同时，z指针上升
				z.P.Color = Black
				y.Color = Black
				z.P.P.Color = Red
				z = z.P.P
			} else {  // 当叔叔结点为黑色情况，为了保持红黑性质，进行旋转修复
				if z == z.P.Right {
					z = z.P
					t.LeftRotate(z)
				}
				z.P.Color = Black
				z.P.P.Color = Red
				t.RightRotate(z.P.P)
			}

		} else if z.P == z.P.P.Right {
			y := z.P.P.Left
			if y.Color == Red {
				z.P.Color = Black
				y.Color = Black
				z.P.P.Color = Red
				z = z.P.P
			} else {
				if z == z.P.Left {
					z = z.P
					t.RightRotate(z)
				}
				z.P.Color = Black
				z.P.P.Color = Red
				t.LeftRotate(z.P.P)
			}
		}
	}
	t.Root.Color = Black
}

