# 56. Merge Intervals

**Объединение интервалов** · 🟡 Medium · Intervals · [LeetCode](https://leetcode.com/problems/merge-intervals/description/)

## Условие

Дан массив интервалов `intervals`, где `intervals[i] = [startᵢ, endᵢ]`. Объедините все пересекающиеся интервалы и верните *массив непересекающихся интервалов, которые покрывают все интервалы из входных данных*.

## Примеры

**Пример 1:**

```text
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: интервалы [1,3] и [2,6] пересекаются, поэтому объединяем их в [1,6].
```

**Пример 2:**

```text
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: интервалы [1,4] и [4,5] считаются пересекающимися.
```

**Пример 3:**

```text
Input: intervals = [[4,7],[1,4]]
Output: [[1,7]]
Explanation: интервалы [1,4] и [4,7] считаются пересекающимися.
```

## Ограничения

- `1 ≤ intervals.length ≤ 10⁴`
- `intervals[i].length == 2`
- `0 ≤ startᵢ ≤ endᵢ ≤ 10⁴`
