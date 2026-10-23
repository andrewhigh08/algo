# 235. Lowest Common Ancestor of a Binary Search Tree

**Наименьший общий предок бинарного дерева поиска** · 🟡 Medium · Trees · [LeetCode](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/)

## Условие

Дано бинарное дерево поиска (BST). Найдите наименьшего общего предка (LCA) двух заданных узлов в BST.

Согласно [определению LCA в Википедии](https://en.wikipedia.org/wiki/Lowest_common_ancestor): «Наименьший общий предок двух узлов `p` и `q` определяется как самый нижний узел в `T`, потомками которого являются и `p`, и `q` (где мы разрешаем **узлу быть потомком самого себя**)».

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png)

```text
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
Output: 6
Explanation: LCA узлов 2 и 8 равен 6.
```

**Пример 2:**

![Пример 2](https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png)

```text
Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
Output: 2
Explanation: LCA узлов 2 и 4 равен 2, поскольку по определению LCA узел может быть потомком самого себя.
```

**Пример 3:**

```text
Input: root = [2,1], p = 2, q = 1
Output: 2
```

## Ограничения

- Количество узлов в дереве — в диапазоне `[2, 10⁵]`.
- `-10⁹ ≤ Node.val ≤ 10⁹`
- Все `Node.val` **уникальны**.
- `p ≠ q`
- `p` и `q` существуют в BST.
