package Week_03

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(generateParenthesis(3))
	root := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	obj := Constructor()
	data := obj.serialize(root)
	ans := obj.deserialize(data)
	fmt.Println(data)
	bytes, _ := json.Marshal(*ans)
	fmt.Println(string(bytes))
}

//https://leetcode-cn.com/problems/generate-parentheses/submissions/
func generateParenthesis(n int) []string {
	var left, right int
	result := ""
	results := &[]string{}
	_generate(n, left, right, result, results)
	return *results
}

func _generate(max int, left int, right int, result string, results *[]string) {
	//fmt.Println(left, right, result)
	if left == right && right == max {
		*results = append(*results, result)
		return
	}
	if left < max {
		//result += "("
		_generate(max, left+1, right, result+"(", results)
	}
	if right < left {
		//result += ")"
		_generate(max, left, right+1, result+")", results)
	}

}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//中序排列为递增序列
func isValidBST(root *TreeNode) bool {
	pre := math.MinInt64
	return bst(root, &pre)
}

//pre 一定要用指针，指针相当与全局数据，pre的值会严格按照代码的执行顺序赋值， 不会收到递归函数栈的影响
func bst(root *TreeNode, pre *int) bool {
	if root != nil {
		if !bst(root.Left, pre) {
			return false
		}
		fmt.Println(root.Val, *pre)
		if root.Val <= *pre {
			return false
		}

		*pre = root.Val
		if !bst(root.Right, pre) {
			return false
		}
	}
	return true

}

//https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/
//前序遍历， 根左右
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "null,"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	data = strings.Trim(data, ",")
	sl := strings.Split(data, ",")
	return _deserialize(&sl)

}

func _deserialize(sl *[]string) *TreeNode {
	cur := (*sl)[0]
	*sl = (*sl)[1:]
	if cur == "null" {
		return nil
	}
	i, _ := strconv.ParseInt(cur, 10, 64)
	node := &TreeNode{
		Val: int(i),
	}
	node.Left = _deserialize(sl)
	node.Right = _deserialize(sl)
	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
