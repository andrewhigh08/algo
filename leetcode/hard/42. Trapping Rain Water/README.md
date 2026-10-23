# 42. Trapping Rain Water

**Сбор дождевой воды** · 🔴 Hard · Two Pointers · [LeetCode](https://leetcode.com/problems/trapping-rain-water/description/)

## Условие

Даны `n` неотрицательных целых чисел, задающих карту высот, где ширина каждого столбца равна `1`. Вычислите, сколько воды она может удержать после дождя.

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2018/10/22/rainwatertrap.png)

```text
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: карта высот на рисунке выше (чёрная часть) задаётся массивом [0,1,0,2,1,0,1,3,2,1,2,1]. В этом случае удерживается 6 единиц дождевой воды (синяя часть).
```

**Пример 2:**

```text
Input: height = [4,2,0,3,2,5]
Output: 9
```

## Ограничения

- `n == height.length`
- `1 ≤ n ≤ 2 * 10⁴`
- `0 ≤ height[i] ≤ 10⁵`
