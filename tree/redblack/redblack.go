package redblack


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
	Red Color = iota + 1
	Black
)


type Node struct {
	Color  Color
	Key    int
	Left   *Node
	Right  *Node
	P      *Node
}

// 左旋
func LeftRotate(root, x *Node) {
	y := x.Right       // 找到右结点y
	x.Right = y.Left
	if y.Left != nil {   // y有左孩子，将左孩子给x
		y.Left.P = x
	}
	y.P = x.P          // 将y父结点变为原x父结点
	if x.P == nil{     // x为根结点的情况
		root = y
	} else if x == x.P.Left {
		x.P.Left = y
	} else {
		x.P.Right = y
	}
	y.Left = x
	x.P = y

}

