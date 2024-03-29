## 赫夫曼树
赫夫曼（Huffman）树，又称最优树，是一类带权路径长度最短的树。

树的路径和路径长度：从树中一个节点到另一个节点之间的分支构成这两个节点之间的路径，路径上的分支数目乘坐**路径长度**。**树的路径长度**是从树根到每一个节点的路径长度之和。

带权路径长度：从该节点到树根之间的路径长度与节点上权的乘积，**树的带权路径长度**为树中所有叶子结点的带权路径长度之和，通常记作：![](http://latex.codecogs.com/svg.latex?WPL=\sum_{k=1}^{n}w_kl_k) $ WPL=\sum_{k = 1}^{n}w_kl_k $。

例如下面三颗树的路径长度为：

![3tree.png](https://i.loli.net/2021/01/18/iN5Wx6jX7MPI2TS.png)

- (a) WPL = 7 * 2 + 5 * 2 + 2 * 2 + 4 * 2 = 36
- (b) WPL = 7 * 3 + 5 * 3 + 2 * 1 + 4 * 2 = 46
- (c) WPL = 7 * 1 + 5 * 2 + 2 * 3 + 4 * 3 = 35

其中（c）树的路径长度最小，其就是赫夫曼树，即带权路径长度在所有带权为 7、5、2、4 的 4 个叶子节点的二叉树中居最小。

### 构建赫夫曼树的过程
1. 根据给定的 n 个权值 {w1, w2, ..., wn} 构成 n 棵二叉树的集合 $ F = {T_1, T_2, ..., T_n} $, 其中每棵二叉树 $ T_i $ 中只有一个带权为 $ w_i $ 的跟节点，其中左右子树均为空。
2. 在 F 中选取两棵跟节点的权值最小的树作为左右子树构建一棵新的二叉树，且置新的二叉树的根节点的权值为其左、右子节点的权值之和。
3. 在 F 中删除这两棵树，同时将新得到的二叉树加入到 F 中。
- 重复 2、3，直到 F 中只有一棵树为止。这棵树就是赫夫曼树。

构造过程如下图所示：

![hfm.png](https://i.loli.net/2021/01/18/rg264bxYc5RvIGe.png)


### 赫夫曼树的应用 -- 数据压缩
霍夫曼编码(Huffman Coding)是一种基于最小冗余编码的压缩算法。最小冗余编码是指，如果知道一组数据中符号出现的频率，就可以用一种特殊的方式来表示符号从而减少数据需要的存储空间。

一个字符串为 `ABACCDA` 假设 A、B、C、D 的编码分别是 00，01，10，11，则上述7个字符的字符串可编译为 `00010010101100`（长度为14），接收时按照2bit一字符的方式译码。当然我们希望字符串可以尽可能的短，占用更少的空间，假设我们重新设计  A、B、C、D 的编码分别是 0，00，1，01，那么我们将上面7个字符编码为：`000011010`(长度为9)，但是这样的编码无法翻译，比如 `0000` 即可翻译为 AAAA 也可翻译为 ABA。因此若要设计长短不等的编码,则必须是任一个字符的编码都不是另一个字符的编码的前缀，这样的编码称做**前缀编码**。

假设每种字符在字符串中出现的频率为 $ w_i $, 该字符编码的长度为 $ l_i $，则字符串总长度为 ![](http://latex.codecogs.com/svg.latex?\sum_{i = 1}^{n}w_il_i) $ \sum_{i = 1}^{n}w_il_i $，这个是不是看着有点眼熟？没错就是赫夫曼树**带权路径长度**的公式。将字符在字符串中出现的频率作为权值，字符编码长度为根到叶子节点的路径长度。

根据上面的公式可以设计字符串 `ABACCDA`，中A、B、C、D 的编码。A、B、C、D 的出现的概率从小到大排序为 `D-B-C-A` 或 `B-D-C-A`，约定左右分支分别代表 0 和 1，根据排序的结果构建赫夫曼树，并得到编码集：

![tree.png](https://i.loli.net/2021/01/20/L1rjDen6RUJ9oWK.png)

编码和译码：
- 编码，可以从根部出发也可以从叶子节点出发遍历整棵树生成字符的赫夫曼编码。
- 译码，读取编码中的 0 或 1 ，从根部出发直到寻找到叶子节点，即可得到字符原文。

#### 例子 bmp 图像压缩
bmp 头部结构：
![bbbb-min.png](https://i.loli.net/2021/04/30/ibNshoZ6GtalpL5.png)