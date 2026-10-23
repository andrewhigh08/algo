# 102. Binary Tree Level Order Traversal

**Обход бинарного дерева по уровням** · 🟡 Medium · Trees · [LeetCode](https://leetcode.com/problems/binary-tree-level-order-traversal/description/)

## Условие

Дан корень бинарного дерева `root`. Верните *значения его узлов в порядке обхода по уровням* (то есть слева направо, уровень за уровнем).

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg)

```text
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]
```

**Пример 2:**

```text
Input: root = [1]
Output: [[1]]
```

**Пример 3:**

```text
Input: root = []
Output: []
```

## Ограничения

- Количество узлов в дереве — в диапазоне `[0, 2000]`.
- `-1000 ≤ Node.val ≤ 1000`
