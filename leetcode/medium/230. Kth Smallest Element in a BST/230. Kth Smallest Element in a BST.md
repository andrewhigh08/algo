# 230. [K-й наименьший элемент в BST (Kth Smallest Element in a BST)](https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/)

`Medium`

Дан `root` бинарного дерева поиска и целое число `k`, верните k-е наименьшее значение (**с индексом 1**) из всех значений узлов в дереве.

**Пример 1:**\
![](https://assets.leetcode.com/uploads/2021/01/28/kthtree1.jpg)
```
Input: root = [3,1,4,null,2], k = 1
Output: 1
```

**Пример 2:**\
![](https://assets.leetcode.com/uploads/2021/01/28/kthtree2.jpg)
```
Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3
```

**Ограничения:**

*   Количество узлов в дереве равно `n`.
*   1 ≤ `k` ≤ `n` ≤ 10⁴
*   0 ≤ `Node.val` ≤ 10⁴

**Дополнительный вопрос:** Если BST часто модифицируется (т.е. мы можем выполнять операции вставки и удаления) и вам нужно часто находить k-е наименьшее значение, как бы вы оптимизировали?