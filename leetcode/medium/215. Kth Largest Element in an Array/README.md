# 215. Kth Largest Element in an Array

**K-й по величине элемент в массиве** · 🟡 Medium · Heap / Priority Queue · [LeetCode](https://leetcode.com/problems/kth-largest-element-in-an-array/description/)

## Условие

Дан массив целых чисел `nums` и целое число `k`. Верните *`k`-й по величине элемент массива*.

Обратите внимание, что это `k`-й по величине элемент в отсортированном порядке, а не `k`-й различный элемент.

Можете ли вы решить задачу без сортировки?

## Примеры

**Пример 1:**

```text
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
```

**Пример 2:**

```text
Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4
```

## Ограничения

- `1 ≤ k ≤ nums.length ≤ 10⁵`
- `-10⁴ ≤ nums[i] ≤ 10⁴`
