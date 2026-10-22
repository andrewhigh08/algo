# 235. [Наименьший Общий Предок Бинарного Дерева Поиска (Lowest Common Ancestor of a Binary Search Tree)](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/)

`Medium`

Дано бинарное дерево поиска (BST), найдите наименьшего общего предка (LCA) двух заданных узлов в BST.

Согласно [определению LCA на Википедии](https://en.wikipedia.org/wiki/Lowest_common_ancestor): "Наименьший общий предок определяется между двумя узлами `p` и `q` как самый нижний узел в `T`, у которого `p` и `q` являются потомками (где мы разрешаем узлу быть потомком самого себя)."

**Пример 1:**\
![](https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png)
```
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: LCA узлов 2 и 8 равен 6.
```

**Пример 2:**\
![](https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png)
```
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: LCA узлов 2 и 4 равен 2, поскольку узел может быть потомком самого себя согласно определению LCA.
```

**Пример 3:**
```
Input: root = [2,1], p = 2, q = 1
Output: 2
```

**Ограничения:**

*   Количество узлов в дереве находится в диапазоне [2, 10⁵].
*   -10⁹ ≤ `Node.val` ≤ 10⁹
*   Все `Node.val` уникальны.
*   `p != q`
*   `p` и `q` будут существовать в BST.