# 84. [Самый большой прямоугольник в гистограмме (Largest Rectangle in Histogram)](https://leetcode.com/problems/largest-rectangle-in-histogram/description/)

сложность: `Hard`

Дан массив целых чисел `heights`, представляющий высоты столбцов гистограммы, где ширина каждого столбца равна 1. Верните *площадь самого большого прямоугольника в гистограмме*.

**Пример 1:**

![Гистограмма для Примера 1](https://assets.leetcode.com/uploads/2021/01/04/histogram.jpg)

```
Input: heights = [2,1,5,6,2,3]
Output: 10
Explanation: The above is a histogram where width of each bar is 1.
The largest rectangle is shown in the red area, which has an area = 10 units.
```

**Пример 2:**

![Гистограмма для Примера 2](https://assets.leetcode.com/uploads/2021/01/04/histogram-1.jpg)

```
Input: heights = [2,4]
Output: 4
```

**Ограничения:**

*   `1 ≤ heights.length ≤ 10⁵`
*   `0 ≤ heights[i] ≤ 10⁴`