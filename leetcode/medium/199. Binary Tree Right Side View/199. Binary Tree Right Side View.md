# 199. [Вид бинарного дерева справа (Binary Tree Right Side View)](https://leetcode.com/problems/binary-tree-right-side-view/description/)

`Medium`

Дан `root` бинарного дерева, представьте, что вы стоите с **правой стороны** от него, верните значения узлов, которые вы можете видеть, упорядоченные сверху вниз.

**Пример 1:**
```
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]
```
![](https://assets.leetcode.com/uploads/2024/11/24/tmpd5jn43fs-1.png)

**Пример 2:**
```
Input: root = [1,2,3,4,null,null,null,5]
Output: [1,3,5]
```
![](https://assets.leetcode.com/uploads/2024/11/24/tmpkpe40xeh-1.png)

**Пример 3:**
```
Input: root = [1,null,3]
Output: [1,3]
```

**Пример 4:**
```
Input: root = []
Output: []
```

**Ограничения:**

*   Количество узлов в дереве находится в диапазоне [0, 100].
*   -100 ≤ `Node.val` ≤ 100
```