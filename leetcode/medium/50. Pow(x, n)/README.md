# 50. Pow(x, n)

**Возведение в степень** · 🟡 Medium · Math & Geometry · [LeetCode](https://leetcode.com/problems/powx-n/description/)

## Условие

Реализуйте функцию [`pow(x, n)`](http://www.cplusplus.com/reference/valarray/pow/), которая вычисляет `x`, возведённое в степень `n` (то есть `xⁿ`).

## Примеры

**Пример 1:**

```text
Input: x = 2.00000, n = 10
Output: 1024.00000
```

**Пример 2:**

```text
Input: x = 2.10000, n = 3
Output: 9.26100
```

**Пример 3:**

```text
Input: x = 2.00000, n = -2
Output: 0.25000
Explanation: 2⁻² = 1/2² = 1/4 = 0.25
```

## Ограничения

- `-100.0 < x < 100.0`
- `-2³¹ ≤ n ≤ 2³¹ - 1`
- `n` — целое число.
- Либо `x` не равно нулю, либо `n > 0`.
- `-10⁴ ≤ xⁿ ≤ 10⁴`
