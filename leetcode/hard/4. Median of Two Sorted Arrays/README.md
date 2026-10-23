# 4. Median of Two Sorted Arrays

**Медиана двух отсортированных массивов** · 🔴 Hard · Binary Search · [LeetCode](https://leetcode.com/problems/median-of-two-sorted-arrays/description/)

## Условие

Даны два отсортированных массива `nums1` и `nums2` размеров `m` и `n` соответственно. Верните **медиану** двух отсортированных массивов.

Общая сложность алгоритма по времени должна быть `O(log (m+n))`.

## Примеры

**Пример 1:**

```text
Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: объединённый массив = [1,2,3], медиана равна 2.
```

**Пример 2:**

```text
Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: объединённый массив = [1,2,3,4], медиана равна (2 + 3) / 2 = 2.5.
```

## Ограничения

- `nums1.length == m`
- `nums2.length == n`
- `0 ≤ m ≤ 1000`
- `0 ≤ n ≤ 1000`
- `1 ≤ m + n ≤ 2000`
- `-10⁶ ≤ nums1[i], nums2[i] ≤ 10⁶`
