# 572. [Поддерево другого дерева (Subtree of Another Tree)](https://leetcode.com/problems/subtree-of-another-tree/description/)

`Easy`

Даны корни двух бинарных деревьев `root` и `subRoot`. Верните `true`, если существует поддерево `root` с той же структурой и значениями узлов, что и `subRoot`, и `false` в противном случае.

Поддерево бинарного дерева `tree` - это дерево, состоящее из узла в `tree` и всех потомков этого узла. Дерево `tree` также может рассматриваться как поддерево самого себя.

**Пример 1:**\
![](https://assets.leetcode.com/uploads/2021/04/28/subtree1-tree.jpg)
```
Input: root = [3,4,5,1,2], subRoot = [4,1,2]
Output: true
```

**Пример 2:**\
![](https://assets.leetcode.com/uploads/2021/04/28/subtree2-tree.jpg)
```
Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
Output: false
```

**Ограничения:**

*   Количество узлов в дереве `root` находится в диапазоне [1, 2000].
*   Количество узлов в дереве `subRoot` находится в диапазоне [1, 1000].
*   -10⁴ ≤ `root.val` ≤ 10⁴
*   -10⁴ ≤ `subRoot.val` ≤ 10⁴