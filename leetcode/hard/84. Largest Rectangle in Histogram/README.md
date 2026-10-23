# 84. Largest Rectangle in Histogram

**Самый большой прямоугольник в гистограмме** · 🔴 Hard · Stack · [LeetCode](https://leetcode.com/problems/largest-rectangle-in-histogram/description/)

## Условие

Дан массив целых чисел `heights`, представляющий высоты столбцов гистограммы, где ширина каждого столбца равна `1`. Верните *площадь самого большого прямоугольника в гистограмме*.

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2021/01/04/histogram.jpg)

```text
Input: heights = [2,1,5,6,2,3]
Output: 10
Explanation: выше изображена гистограмма, ширина каждого столбца которой равна 1.
Самый большой прямоугольник выделен красным, его площадь равна 10 единицам.
```

**Пример 2:**

![Пример 2](https://assets.leetcode.com/uploads/2021/01/04/histogram-1.jpg)

```text
Input: heights = [2,4]
Output: 4
```

## Ограничения

- `1 ≤ heights.length ≤ 10⁵`
- `0 ≤ heights[i] ≤ 10⁴`
