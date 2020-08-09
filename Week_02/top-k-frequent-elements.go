package Week_02

import "container/heap"

//https://leetcode-cn.com/problems/top-k-frequent-elements/
//统计元素出现的次数
//然后放入priorityQue中
//依次pop出k个元素
// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for _, n := range nums {
		if _, ok := m[n]; !ok {
			m[n] = 0
		}
		m[n]++
	}
	priorityQ := PriorityQueue{}
	i := 0
	for k, v := range m {
		priorityQ = append(priorityQ, &Item{value: k, priority: v, index: i})
		i++
	}
	ret := make([]int, 0, k)
	heap.Init(&priorityQ)
	for i := 0; i < k; i++ {
		ret = append(ret, heap.Pop(&priorityQ).(*Item).value)
	}
	return ret
}
