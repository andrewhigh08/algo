# 105. Construct Binary Tree from Preorder and Inorder Traversal

**Построить бинарное дерево из обходов preorder и inorder** · 🟡 Medium · Trees · [LeetCode](https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/)

## Условие

Даны два целочисленных массива `preorder` и `inorder`, где `preorder` — это обход бинарного дерева в прямом порядке (preorder traversal), а `inorder` — обход того же дерева в симметричном порядке (inorder traversal). Постройте и верните *бинарное дерево*.

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2021/02/19/tree.jpg)

```text
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
```

**Пример 2:**

```text
Input: preorder = [-1], inorder = [-1]
Output: [-1]
```

## Ограничения

- `1 ≤ preorder.length ≤ 3000`
- `inorder.length == preorder.length`
- `-3000 ≤ preorder[i], inorder[i] ≤ 3000`
- `preorder` и `inorder` состоят из **уникальных** значений.
- Каждое значение из `inorder` также встречается в `preorder`.
- `preorder` **гарантированно** является обходом дерева в прямом порядке.
- `inorder` **гарантированно** является обходом дерева в симметричном порядке.
