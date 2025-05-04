# 15. [3Сумма (3Sum)](https://leetcode.com/problems/3sum/description/)

сложность: `Medium`

Дан массив целых чисел `nums`. Верните все триплеты `[nums[i], nums[j], nums[k]]` такие, что `i != j`, `i != k`, и `j != k`, и `nums[i] + nums[j] + nums[k] == 0`.

Обратите внимание, что набор решений не должен содержать повторяющихся триплетов.

```
Пример 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Объяснение: 
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.

Пример 2:

Input: nums = [0,1,1]
Output: []
Объяснение: The only possible triplet does not sum up to 0.

Пример 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Объяснение: The only possible triplet sums up to 0.
```

**Ограничения:**

*   `3 ≤ nums.length ≤ 3000`
*   `-10⁵ ≤ nums[i] ≤ 10⁵`