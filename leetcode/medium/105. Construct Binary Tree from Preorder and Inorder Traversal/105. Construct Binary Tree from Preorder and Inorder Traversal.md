# 105. [Построить бинарное дерево из обходов Preorder и Inorder (Construct Binary Tree from Preorder and Inorder Traversal)]()

`Medium`

Даны два целочисленных массива `preorder` и `inorder`, где `preorder` — это обход в прямом порядке (preorder traversal) бинарного дерева, а `inorder` — это обход в симметричном порядке (inorder traversal) того же дерева. Постройте и верните *бинарное дерево*.

**Пример 1:**\
![](https://assets.leetcode.com/uploads/2021/02/19/tree.jpg)
```
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
```

**Пример 2:**
```
Input: preorder = [-1], inorder = [-1]
Output: [-1]
```

**Ограничения:**

*   1 ≤ `preorder.length` ≤ 3000
*   `inorder.length == preorder.length`
*   -3000 ≤ `preorder[i]`, `inorder[i]` ≤ 3000
*   `preorder` и `inorder` состоят из **уникальных** значений.
*   Каждое значение из `inorder` также появляется в `preorder`.
*   `preorder` **гарантированно** является обходом в прямом порядке дерева.
*   `inorder` **гарантированно** является обходом в симметричном порядке дерева.