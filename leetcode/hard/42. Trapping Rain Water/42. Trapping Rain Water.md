# 42. [Сбор дождевой воды (Trapping Rain Water)](https://leetcode.com/problems/trapping-rain-water/description/)

сложность: `Hard`

Даны `n` неотрицательных целых чисел, представляющих карту высот, где ширина каждого столбца равна 1. Вычислите, сколько воды он может удержать после дождя.

**Пример 1:**

![Карта высот и удержанная вода для Примера 1](https://assets.leetcode.com/uploads/2018/10/22/rainwatertrap.png)

```
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. 
In this case, 6 units of rain water (blue section) are being trapped.

Пример 2:

Input: height = [4,2,0,3,2,5]
Output: 9
```

**Ограничения:**

*   `n == height.length`
*   `1 ≤ n ≤ 2 * 10⁴`
*   `0 ≤ height[i] ≤ 10⁵`