# 4. [Медиана двух отсортированных массивов (Median of Two Sorted Arrays)](https://leetcode.com/problems/median-of-two-sorted-arrays/description/)

сложность: `Hard`

Даны два отсортированных массива `nums1` и `nums2` размеров `m` и `n` соответственно. Верните **медиану** двух отсортированных массивов.

Общая сложность по времени выполнения должна быть O(log(m+n)).

```
Пример 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.

Пример 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
```


**Ограничения:**

*   `nums1.length == m`
*   `nums2.length == n`
*   `0 ≤ m ≤ 1000`
*   `0 ≤ n ≤ 1000`
*   `1 ≤ m + n ≤ 2000`
*   `-10⁶ ≤ nums1[i], nums2[i] ≤ 10⁶`