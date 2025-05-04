# 329. [Самый Длинный Возрастающий Путь в Матрице (Longest Increasing Path in a Matrix)](https://leetcode.com/problems/longest-increasing-path-in-a-matrix/description/)

`Hard`

Дана целочисленная матрица `matrix` размером `m x n`, верните *длину самого длинного возрастающего пути* в `matrix`.

Из каждой ячейки вы можете двигаться в четырех направлениях: влево, вправо, вверх или вниз. Вы **не можете** двигаться **по диагонали** или выходить **за пределы** (т.е. обход по кругу не разрешен).

**Пример 1:**\
![](https://assets.leetcode.com/uploads/2021/01/05/grid1.jpg)
```
Input: matrix = [[9,9,4],[6,6,8],[2,1,1]]
Output: 4
Explanation: Самый длинный возрастающий путь - [1, 2, 6, 9].
```

**Пример 2:**\
![](https://assets.leetcode.com/uploads/2021/01/27/tmp-grid.jpg)
```
Input: matrix = [[3,4,5],[3,2,6],[2,2,1]]
Output: 4
Explanation: Самый длинный возрастающий путь - [3, 4, 5, 6]. Движение по диагонали не разрешено.
```

**Пример 3:**
```
Input: matrix = [[1]]
Output: 1
```

**Ограничения:**

*   `m == matrix.length`
*   `n == matrix[i].length`
*   1 ≤ m, n ≤ 200
*   0 ≤ `matrix[i][j]` ≤ 2³¹ - 1