## 堆（Heap）
堆就是用数组实现的完全二叉树，所以它没有使用父指针或者子指针。

堆分为两种，将根节点最大的堆叫做最大堆或大根堆，根节点最小的堆叫做最小堆或小根堆：
- 最大堆：各个节点大于或等于其子节点
- 最小堆：各个节点小于或等于其子节点

节点与其子节点、父节点对应关系（i 是当前节点在数组中的索引）：
- 父节点 ![](http://latex.codecogs.com/svg.latex?floor((i-1)/2))
- 左子节点：![](http://latex.codecogs.com/svg.latex?i*2+1)
- 右子节点：![](http://latex.codecogs.com/svg.latex?i*2+2)

n 个结点的堆，深度 ![](https://latex.codecogs.com/svg.latex?h=%5Clog\%20n), 根为第0层，则第i层结点个数为 ![](https://latex.codecogs.com/svg.latex?2^i)，考虑一个元素在堆中向下移动的距离，这种算法时间代价为 Ο(n).

由于堆有 ![](https://latex.codecogs.com/svg.latex?\log%20n) 层深，插入结点、删除普通元素和删除最小元素的平均时间代价和时间复杂度都是 ![](https://latex.codecogs.com/svg.latex?O\log%20n).


如下图为一个最大堆：
![屏幕快照 2020-12-22 下午4.08.29.png](https://i.loli.net/2020/12/22/glfJnV697qCPZGm.png)

其在数组中的顺序为：

||||||||
|:----:|:----:|:----:|:----:|:----:|:----:|:----:|
| 10 | 7 | 9 | 5 | 1 | 2 | 8 |