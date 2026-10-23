# 110. Balanced Binary Tree

**Сбалансированное бинарное дерево** · 🟢 Easy · Trees · [LeetCode](https://leetcode.com/problems/balanced-binary-tree/description/)

## Условие

Дано бинарное дерево. Определите, является ли оно **сбалансированным по высоте**.

**Сбалансированное по высоте** бинарное дерево — это бинарное дерево, в котором глубины левого и правого поддеревьев каждого узла никогда не отличаются более чем на единицу.

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2020/10/06/balance_1.jpg)

```text
Input: root = [3,9,20,null,null,15,7]
Output: true
```

**Пример 2:**

![Пример 2](https://assets.leetcode.com/uploads/2020/10/06/balance_2.jpg)

```text
Input: root = [1,2,2,3,3,null,null,4,4]
Output: false
```

**Пример 3:**

```text
Input: root = []
Output: true
```

## Ограничения

- Количество узлов в дереве — в диапазоне `[0, 5000]`.
- `-10⁴ ≤ Node.val ≤ 10⁴`
