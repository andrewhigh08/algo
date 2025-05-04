# 54. [Спиральная Матрица (Spiral Matrix)](https://leetcode.com/problems/spiral-matrix/description/)

`Medium`

Дана матрица `matrix` размером `m x n`, верните все элементы матрицы в спиральном порядке.

**Пример 1:**\
![](https://assets.leetcode.com/uploads/2020/11/13/spiral1.jpg)
```
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]
```

**Пример 2:**\
![](https://assets.leetcode.com/uploads/2020/11/13/spiral.jpg)
```
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
```

**Ограничения:**

*   `m == matrix.length`
*   `n == matrix[i].length`
*   1 ≤ m, n ≤ 10
*   -100 ≤ `matrix[i][j]` ≤ 100