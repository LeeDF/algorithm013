package Week_01

import (
	"encoding/json"
	"fmt"
)

//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
//remove-duplicates-from-sorted-array
//1.快慢双指针, 慢指针i记录不重复index, 快指针j遍历数组
//2. nums[i] == nums[j] 则j++
//3. nums[i] != nums[j] 则将nums[j] 移到i+1位置, 同时i++
func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

//---------------------------------------------------------------------------------------------------------------------
//https://leetcode-cn.com/problems/merge-two-sorted-lists/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//根据l1移动l1和l2各自指针， 如果l2指针节点>l1指针节点 <= l1指针下一个节点，则将l2指针节点移到l1出
//O(n)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		tmp := l1
		l1 = l2
		l2 = tmp
	}

	p1 := l1
	p2 := l2
	for {
		if p2.Val >= p1.Val && (p1.Next == nil || p2.Val <= p1.Next.Val) {
			tp := p2
			p2 = p2.Next
			tp.Next = p1.Next
			p1.Next = tp
		}

		p1 = p1.Next
		if p2 == nil {
			break
		}
	}
	return l1
}

func print(l1 *ListNode) {
	bytes, _ := json.Marshal(l1)
	fmt.Println(string(bytes))
}

//---------------------------------------------------------------------------------------------------------------------
//https://leetcode-cn.com/problems/design-circular-deque/
//1.无随机访问，所以用双向链表
//2.切片替换链表
//3.静态数组
type DequeNode struct {
	Val  int
	Next *DequeNode
	Prev *DequeNode
}

type MyCircularDeque struct {
	cap  int
	len  int
	head *DequeNode
	tail *DequeNode
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		cap: k,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	newNode := &DequeNode{Val: value}
	if this.len == 0 {
		newNode.Next = newNode
		newNode.Prev = newNode
		this.head = newNode
		this.tail = newNode
	} else {
		newNode.Next = this.head
		this.head.Prev = newNode
		this.head = newNode
	}
	this.len++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	newNode := &DequeNode{Val: value}
	if this.len == 0 {
		this.head = newNode
		this.tail = newNode
	} else {
		newNode.Prev = this.tail
		this.tail.Next = newNode
		this.tail = newNode
	}
	this.len++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.len == 0 {
		return false
	}
	this.head = this.head.Next

	if this.head != nil {
		this.head.Prev = nil
	}

	this.len--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.len == 0 {
		return false
	}
	this.tail = this.tail.Prev
	if this.tail != nil {
		this.tail.Next = nil
	}
	this.len--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.head.Val
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.tail.Val
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if this.len == 0 {
		return true
	}
	return false
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if this.len == this.cap {
		return true
	}
	return false
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

//---------------------------------------------------------------------------------------------------------------------
//https://leetcode-cn.com/problems/trapping-rain-water/solution/xiang-xi-tong-su-de-si-lu-fen-xi-duo-jie-fa-by-w-8/
//1.暴力法O(n^2)：遍历所有元素，找出左右最高节点，
// 	如果左右最高元素大于当前元素高度，则面积位较矮的高节点减去当前节点的高度
//2.优化：单调递减栈，入栈元素比栈顶小
//依次入栈，如果入栈元素比栈顶大，则出栈，小的元素已经能够计算雨水了
//小的元素计算完之后出栈
//大于等于自己的元素保留在栈内，遍历一次O(n)
//

func trap_1(height []int) int {
	area := 0
	for i := 1; i < len(height)-1; i++ {
		//找左最高元素
		maxleft := 0
		for j := 0; j < i; j++ {
			maxleft = max(maxleft, height[j])
		}
		//找右最高元素
		maxright := 0
		for j := i + 1; j < len(height); j++ {
			maxright = max(maxright, height[j])
		}
		minHeight := min(maxleft, maxright)
		if minHeight > height[i] {
			area += minHeight - height[i]
		}
	}
	return area
}

//func trap_2(height []int) int {
//	if len(height) == 0 {
//		return 0
//	}
//	area := 0
//	stack := make([]int, 0, len(height)) //栈内存index
//	for i := 0; i < len(height); i++ {
//		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
//			fmt.Println(stack)
//			tmp := stack[len(stack)-1]
//			if height[tmp] > height[i] {
//				stack = stack[:len(stack)-1]
//				break
//			}
//			area += (min(height[i], stack[0]) - height[0]) * (i - stack[0]-1)
//			fmt.Println((min(height[i], stack[0]) - height[0]) * (i - stack[0]-1))
//			stack = stack[:len(stack)-1]
//		}
//		stack = append(stack, i)
//	}
//	return area
//}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

