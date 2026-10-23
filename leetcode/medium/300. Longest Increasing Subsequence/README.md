# 300. Longest Increasing Subsequence

**Самая длинная возрастающая подпоследовательность** · 🟡 Medium · 1-D Dynamic Programming · [LeetCode](https://leetcode.com/problems/longest-increasing-subsequence/description/)

## Условие

Дан целочисленный массив `nums`. Верните *длину самой длинной **строго возрастающей подпоследовательности***.

## Примеры

**Пример 1:**

```text
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: Самая длинная возрастающая подпоследовательность — [2,3,7,101], поэтому длина равна 4.
```

**Пример 2:**

```text
Input: nums = [0,1,0,3,2,3]
Output: 4
```

**Пример 3:**

```text
Input: nums = [7,7,7,7,7,7,7]
Output: 1
```

## Ограничения

- `1 ≤ nums.length ≤ 2500`
- `-10⁴ ≤ nums[i] ≤ 10⁴`

## Дополнительный вопрос

Можете ли вы придумать алгоритм с временной сложностью `O(n log(n))`?
