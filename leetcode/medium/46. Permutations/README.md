# 46. Permutations

**Перестановки** · 🟡 Medium · Backtracking · [LeetCode](https://leetcode.com/problems/permutations/description/)

## Условие

Дан массив `nums` из различных целых чисел. Верните все возможные перестановки. Ответ можно вернуть в **любом порядке**.

## Примеры

**Пример 1:**

```text
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
```

**Пример 2:**

```text
Input: nums = [0,1]
Output: [[0,1],[1,0]]
```

**Пример 3:**

```text
Input: nums = [1]
Output: [[1]]
```

## Ограничения

- `1 ≤ nums.length ≤ 6`
- `-10 ≤ nums[i] ≤ 10`
- Все целые числа в `nums` **уникальны**.
