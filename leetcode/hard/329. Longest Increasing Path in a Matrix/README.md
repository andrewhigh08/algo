# 329. Longest Increasing Path in a Matrix

**Самый длинный возрастающий путь в матрице** · 🔴 Hard · 2-D Dynamic Programming · [LeetCode](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/description/)

## Условие

Дана целочисленная матрица `matrix` размером `m x n`. Верните *длину самого длинного возрастающего пути в `matrix`*.

Из каждой ячейки можно двигаться в четырёх направлениях: влево, вправо, вверх или вниз. **Нельзя** двигаться **по диагонали** и выходить **за границы** матрицы (т. е. переходить через край на противоположную сторону нельзя).

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2021/01/05/grid1.jpg)

```text
Input: matrix = [[9,9,4],[6,6,8],[2,1,1]]
Output: 4
Explanation: самый длинный возрастающий путь — [1, 2, 6, 9].
```

**Пример 2:**

![Пример 2](https://assets.leetcode.com/uploads/2021/01/27/tmp-grid.jpg)

```text
Input: matrix = [[3,4,5],[3,2,6],[2,2,1]]
Output: 4
Explanation: самый длинный возрастающий путь — [3, 4, 5, 6]. Двигаться по диагонали нельзя.
```

**Пример 3:**

```text
Input: matrix = [[1]]
Output: 1
```

## Ограничения

- `m == matrix.length`
- `n == matrix[i].length`
- `1 ≤ m, n ≤ 200`
- `0 ≤ matrix[i][j] ≤ 2³¹ - 1`
