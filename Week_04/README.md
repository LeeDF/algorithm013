
## 深度优先和广度优先（dfs&bfs）

- 访问所有节点且每个节点被访问一次
- dfs通常是递归，也可以用栈模拟递归
- bfs使用deque， 在go中就是数组表示

## 贪心算法（Greedy）

- 是指在每一步都采取当前最优解，得到最后的最优解
- 贪心算法不会保存中间结果，不会回退，但是动态规划会保存中间结果，并根据结果选择，有回退功能
- 贪心算法一般难点在于如何证明贪心能够得到最优解
- 部分贪心的需要调整贪心的切入点

## 二分查找
- 三个条件，1. 目标具有单调性 2. 存在上下边界 3. 能够通过索引访问
- 类似于二叉搜索树，不过二分查找通常是在数组中

## 实战
- [二叉树层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/)
- [在每个树行中的最大值](https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/)
- [单词接龙](https://leetcode-cn.com/problems/word-ladder/description/)
- [单词接龙II](https://leetcode-cn.com/problems/word-ladder-ii/)
- [岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)
- [扫雷游戏](https://leetcode-cn.com/problems/minesweeper/description/)
- [跳跃游戏](https://leetcode-cn.com/problems/jump-game/)
- [跳跃游戏II](https://leetcode-cn.com/problems/jump-game-ii/)
- [x 的平方根](https://leetcode-cn.com/problems/sqrtx/)
- [搜索旋转数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)
- [搜索二维矩阵](https://leetcode-cn.com/problems/search-a-2d-matrix/)
