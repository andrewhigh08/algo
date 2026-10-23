# 74. Search a 2D Matrix

**Поиск в 2D-матрице** · 🟡 Medium · Binary Search · [LeetCode](https://leetcode.com/problems/search-a-2d-matrix/description/)

## Условие

Вам дана целочисленная матрица `matrix` размером `m x n` со следующими двумя свойствами:

- Каждая строка отсортирована в неубывающем порядке.
- Первое число каждой строки больше последнего числа предыдущей строки.

Дано целое число `target`. Верните `true`, *если `target` есть в `matrix`, и* `false` *в противном случае*.

Решение должно работать за время `O(log(m * n))`.

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2020/10/05/mat.jpg)

```text
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true
```

**Пример 2:**

![Пример 2](https://assets.leetcode.com/uploads/2020/10/05/mat2.jpg)

```text
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false
```

## Ограничения

- `m == matrix.length`
- `n == matrix[i].length`
- `1 ≤ m, n ≤ 100`
- `-10⁴ ≤ matrix[i][j], target ≤ 10⁴`
