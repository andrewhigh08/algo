# 48. Rotate Image

**Повернуть изображение** · 🟡 Medium · Math & Geometry · [LeetCode](https://leetcode.com/problems/rotate-image/description/)

## Условие

Вам дана двумерная матрица `matrix` размером `n x n`, представляющая изображение. Поверните изображение на **90** градусов (по часовой стрелке).

Вы должны повернуть изображение [**на месте**](https://en.wikipedia.org/wiki/In-place_algorithm), то есть изменить входную двумерную матрицу напрямую. **НЕ** выделяйте другую двумерную матрицу, чтобы выполнить поворот.

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2020/08/28/mat1.jpg)

```text
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [[7,4,1],[8,5,2],[9,6,3]]
```

**Пример 2:**

![Пример 2](https://assets.leetcode.com/uploads/2020/08/28/mat2.jpg)

```text
Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]
```

## Ограничения

- `n == matrix.length == matrix[i].length`
- `1 ≤ n ≤ 20`
- `-1000 ≤ matrix[i][j] ≤ 1000`
