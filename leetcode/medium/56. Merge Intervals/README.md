# 56. [Объединение Интервалов (Merge Intervals)](https://leetcode.com/problems/merge-intervals/description/)

`Medium`

Дан массив интервалов `intervals`, где `intervals[i] = [starti, endi]`. Объедините все пересекающиеся интервалы и верните массив непересекающихся интервалов, которые покрывают все интервалы во входных данных.

**Пример 1:**
```
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Поскольку интервалы [1,3] и [2,6] пересекаются, объедините их в [1,6].
```

**Пример 2:**
```
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Интервалы [1,4] и [4,5] считаются пересекающимися.
```

**Ограничения:**

*   1 ≤ `intervals.length` ≤ 10⁴
*   `intervals[i].length == 2`
*   0 ≤ `starti` ≤ `endi` ≤ 10⁴