# 704. [Бинарный поиск (Binary Search)](https://leetcode.com/problems/binary-search/description/)

сложность: `Easy`

Дан массив целых чисел `nums`, который отсортирован в возрастающем порядке, и целое число `target`. Напишите функцию для поиска `target` в `nums`. Если `target` существует, верните его индекс. В противном случае верните -1.

Вы должны написать алгоритм со сложностью по времени выполнения O(log n).

```
Пример 1:

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Объяснение: 9 exists in nums and its index is 4

Пример 2:

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1

Объяснение: 2 does not exist in nums so return -1
```

**Ограничения:**

*   `1 ≤ nums.length ≤ 10⁴`
*   `-10⁴ < nums[i], target < 10⁴`
*   Все целые числа в `nums` **уникальны**.
*   `nums` отсортирован в возрастающем порядке.