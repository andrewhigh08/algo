# 230. Kth Smallest Element in a BST

**K-й наименьший элемент в BST** · 🟡 Medium · Trees · [LeetCode](https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/)

## Условие

Дан корень бинарного дерева поиска `root` и целое число `k`. Верните `k`-е *наименьшее значение (**нумерация с 1**) среди значений всех узлов дерева*.

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2021/01/28/kthtree1.jpg)

```text
Input: root = [3,1,4,null,2], k = 1
Output: 1
```

**Пример 2:**

![Пример 2](https://assets.leetcode.com/uploads/2021/01/28/kthtree2.jpg)

```text
Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3
```

## Ограничения

- Количество узлов в дереве равно `n`.
- `1 ≤ k ≤ n ≤ 10⁴`
- `0 ≤ Node.val ≤ 10⁴`

## Дополнительный вопрос

Если BST часто изменяется (то есть выполняются операции вставки и удаления), а `k`-е наименьшее значение нужно находить часто, как бы вы оптимизировали решение?
