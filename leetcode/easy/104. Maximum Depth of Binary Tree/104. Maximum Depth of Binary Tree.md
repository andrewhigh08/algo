# 104. [Максимальная глубина бинарного дерева (Maximum Depth of Binary Tree)](https://leetcode.com/problems/maximum-depth-of-binary-tree/description/)

`Easy`

Дан `root` бинарного дерева, верните его *максимальную глубину*.

*Максимальная глубина* бинарного дерева — это количество узлов вдоль самого длинного пути от корневого узла до самого дальнего листового узла.

**Пример 1:**\
![](https://assets.leetcode.com/uploads/2020/11/26/tmp-tree.jpg)
```
Input: root = [3,9,20,null,null,15,7]
Output: 3
```

**Пример 2:**
```
Input: root = [1,null,2]
Output: 2
```

**Ограничения:**

*   Количество узлов в дереве находится в диапазоне [0, 10⁴].
*   -100 ≤ `Node.val` ≤ 100