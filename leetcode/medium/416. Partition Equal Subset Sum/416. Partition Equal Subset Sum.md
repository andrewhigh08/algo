416. [Разбиение на Равные Подмножества по Сумме (Partition Equal Subset Sum)](https://leetcode.com/problems/partition-equal-subset-sum/description/)

`Medium`

Дан целочисленный массив `nums`, верните `true`, если вы можете разбить массив на два подмножества так, чтобы сумма элементов в обоих подмножествах была равна, или `false` в противном случае.

**Пример 1:**
```
Input: nums = [1,5,11,5]
Output: true
Explanation: Массив можно разбить на [1, 5, 5] и [11].
```

**Пример 2:**
```
Input: nums = [1,2,3,5]
Output: false
Explanation: Массив нельзя разбить на два подмножества с равными суммами.
```

**Ограничения:**

*   1 ≤ `nums.length` ≤ 200
*   1 ≤ `nums[i]` ≤ 100