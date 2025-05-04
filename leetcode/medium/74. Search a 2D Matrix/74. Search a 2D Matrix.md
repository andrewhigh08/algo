# 74. [Поиск в 2D матрице (Search a 2D Matrix)](https://leetcode.com/problems/search-a-2d-matrix/description/)

сложность: `Medium`

Вам дана целочисленная матрица `matrix` размером `m x n` со следующими двумя свойствами:

*   Каждая строка отсортирована в неубывающем порядке.
*   Первое целое число каждой строки больше последнего целого числа предыдущей строки.

Дано целое число `target`, верните `true`, если `target` находится в `matrix`, и `false` в противном случае.

Вы должны написать решение со сложностью по времени выполнения O(log(m * n)).

**Пример 1:**

![Матрица для Примера 1](https://assets.leetcode.com/uploads/2020/10/05/mat.jpg)

```
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true
```
**Пример 2:**

![Матрица для Примера 2](https://assets.leetcode.com/uploads/2020/10/05/mat2.jpg)

```
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false
```

**Ограничения:**

*   `m == matrix.length`
*   `n == matrix[i].length`
*   `1 ≤ m, n ≤ 100`
*   `-10⁴ ≤ matrix[i][j], target ≤ 10⁴`