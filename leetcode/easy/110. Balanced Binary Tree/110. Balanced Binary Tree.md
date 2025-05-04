110. [Сбалансированное бинарное дерево (Balanced Binary Tree)](https://leetcode.com/problems/balanced-binary-tree/description/)

`Easy`

Дано бинарное дерево, определите, является ли оно **сбалансированным по высоте**.

**Сбалансированное по высоте** бинарное дерево — это бинарное дерево, в котором глубина левого и правого поддеревьев каждого узла никогда не отличается более чем на 1.

**Пример 1:**\
![](https://assets.leetcode.com/uploads/2020/10/06/balance_1.jpg)
```
Input: root = [3,9,20,null,null,15,7]
Output: true
```

**Пример 2:**\
![](https://assets.leetcode.com/uploads/2020/10/06/balance_2.jpg)
```
Input: root = [1,2,2,3,3,null,null,4,4]
Output: false
```

**Пример 3:**
```
Input: root = []
Output: true
```

**Ограничения:**

*   Количество узлов в дереве находится в диапазоне [0, 5000].
*   -10⁴ ≤ `Node.val` ≤ 10⁴