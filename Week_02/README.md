## hash表
- 哈希函数，计算key的hash值，根据hash值存储数据，hashmap就是key-value的形式
hashset 就只有 key
- 不同的key， hash函数算出的hash值有可能相同（hash冲突），通常的做法是在对应位置存一个链表
- 插入、查询、删除的时间复杂度都是O(1)

## 树
- 链表是特殊的树，只有一个子节点的树就是单链表
- 树是特殊的图，没有形成环的图就是树

```
type TreeNode struct {
 	Val int
 	Left *TreeNode
 	Right *TreeNode
 }
 ```


### 二叉树
- 一个节点有两个儿子节点
- 二叉树的遍历：前序（根-左-右） 中序（左-根-有）后序（左-右-根）

### 二叉搜素树（binary search tree）
- 左子树的所有节点小于根节点
- 右子树的所有节点大于根节点
- 中序遍历就是升序
- 查询、插入、删除时间复杂度都是O(logN)
- demo: https://visualgo.net/zh/bst
- 查询插入比较容易， 删除的时候分两种情况
    - 情况一：删除节点是叶子节点，直接删除
    - 情况二：删除节点是非叶子节点，删除节点后，找到比删除节点大且最接近的节点，替换被删除的节点 


## 图
- 邻接距阵和临接表两种表示方式
- 记住BFS和DFS代码模版

## 堆
- 大顶堆，父节点大于子节点
- 小顶堆，父节点小于子节点
- getMin 时间复杂度为O(1)
- push和pop时间复杂度为logN
- 实现， 在src/container/heap/heap.go文件中



### 备忘，刷遍数
https://leetcode-cn.com/problems/valid-anagram/description/
https://leetcode-cn.com/problems/chou-shu-lcof/