# 416. Partition Equal Subset Sum

**Разбиение на равные подмножества по сумме** · 🟡 Medium · 1-D Dynamic Programming · [LeetCode](https://leetcode.com/problems/partition-equal-subset-sum/description/)

## Условие

Дан целочисленный массив `nums`. Верните `true`, *если вы можете разбить массив на два подмножества так, чтобы суммы элементов в обоих подмножествах были равны*, или `false` *в противном случае*.

## Примеры

**Пример 1:**

```text
Input: nums = [1,5,11,5]
Output: true
Explanation: Массив можно разбить на [1, 5, 5] и [11].
```

**Пример 2:**

```text
Input: nums = [1,2,3,5]
Output: false
Explanation: Массив нельзя разбить на два подмножества с равными суммами.
```

## Ограничения

- `1 ≤ nums.length ≤ 200`
- `1 ≤ nums[i] ≤ 100`
