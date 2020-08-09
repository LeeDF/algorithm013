package Week_02

import (
	"fmt"
)

//大顶堆
type IntHeap []int

func (this *IntHeap) Push(val int) {
	*this = append(*this, val) //添加到最后
	//向上调整， 如果比父亲节点大，向上移动
	this.up(len(*this) - 1)
}

//func (this *IntHeap) Pop(val int) int {
//	//取出第一个， 并将最后一个节点放到第一个位置,并移除最后一个位置
//	ret := (*this)[0]
//	(*this)[0] = (*this)[len(*this)-1]
//	*this = (*this)[0 : len(*this)-1]
//	//堆顶向下移动
//}

func (this *IntHeap) Init() {

}

func (this *IntHeap) up(idx int) {
	for {
		pi := parent(idx)
		if pi < 0 || (*this)[idx] <= (*this)[pi] {
			break
		}
		//儿子节点比父节点大，交换
		(*this)[idx], (*this)[pi] = (*this)[pi], (*this)[idx]
		idx = pi
	}
}

func (this *IntHeap) down(idx int) {
	for {
		li := left(idx)
		if li > len(*this)-1 {
			break
		}
		j := li
		ri := right(idx)
		if ri <= len(*this)-1 && (*this)[ri] > (*this)[li] {
			//子节点最大节点
			j = ri
		}
		if (*this)[idx] < (*this)[j] {
			(*this)[idx], (*this)[j] = (*this)[j], (*this)[idx]
		}
		idx = j
	}
}

func left(i int) int {
	return 2*i + 1
}
func right(i int) int {
	return left(i) + 1
}

func parent(i int) int {
	return (i - 1) / 2
}

//https://leetcode-cn.com/problems/recover-binary-search-tree/

//1. 中序遍历为递增数组，交换i，j之后， 那么nums[i] > nums[i+1], nums[j] < nums[j-1]
//2. 时间复杂度为O(1), 查找的时候记录被替换的节点， 找到两个， 则交换两个节点
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func RecoverTree(root *TreeNode) {
	var inorder func(tn *TreeNode)
	var pre *TreeNode
	var i, j, count int

	inorder = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		inorder(tn.Left)
		if pre != nil && pre.Val > tn.Val {
			if count == 0{
				i = pre.Val
				count++
			}
			if count == 1{
				j = tn.Val
				fmt.Println(j)
				//return
			}
		}
		pre = tn
		inorder(tn.Right)
	}
	inorder(root)
	_recoverTree(root, i, j)
}

func _recoverTree(root *TreeNode, i int, j int) {
	if root == nil {
		return
	}
	_recoverTree(root.Left, i, j)
	if root.Val == i {
		root.Val = j
	} else if root.Val == j {
		root.Val = i
	}

	_recoverTree(root.Right, i, j)

}

func main() {
	tree := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 1,
		},
	}
	RecoverTree(tree)
	fmt.Println(tree.Left, tree,  tree.Right)
}
