# 572. Subtree of Another Tree

**Поддерево другого дерева** · 🟢 Easy · Trees · [LeetCode](https://leetcode.com/problems/subtree-of-another-tree/description/)

## Условие

Даны корни двух бинарных деревьев `root` и `subRoot`. Верните `true`, если существует поддерево `root` с той же структурой и теми же значениями узлов, что и `subRoot`, и `false` в противном случае.

Поддерево бинарного дерева `tree` — это дерево, состоящее из какого-либо узла `tree` и всех потомков этого узла. Дерево `tree` также можно считать поддеревом самого себя.

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2021/04/28/subtree1-tree.jpg)

```text
Input: root = [3,4,5,1,2], subRoot = [4,1,2]
Output: true
```

**Пример 2:**

![Пример 2](https://assets.leetcode.com/uploads/2021/04/28/subtree2-tree.jpg)

```text
Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
Output: false
```

## Ограничения

- Количество узлов в дереве `root` — в диапазоне `[1, 2000]`.
- Количество узлов в дереве `subRoot` — в диапазоне `[1, 1000]`.
- `-10⁴ ≤ root.val ≤ 10⁴`
- `-10⁴ ≤ subRoot.val ≤ 10⁴`
