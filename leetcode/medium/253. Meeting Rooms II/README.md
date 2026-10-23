# 253. Meeting Rooms II

**Комнаты для переговоров II** · 🟡 Medium · Intervals · [LeetCode 🔒](https://leetcode.com/problems/meeting-rooms-ii/description/) · [LintCode 919](https://www.lintcode.com/problem/919/)

## Условие

Дан массив интервалов встреч `intervals`, где `intervals[i] = [startᵢ, endᵢ]`. Верните *минимальное количество переговорных комнат*, необходимое для проведения всех встреч.

## Примеры

**Пример 1:**

```text
Input: intervals = [[0,30],[5,10],[15,20]]
Output: 2
Explanation: нужны две комнаты: в первой проходит встреча [0,30], во второй — встречи [5,10] и [15,20].
```

**Пример 2:**

```text
Input: intervals = [[7,10],[2,4]]
Output: 1
```

## Ограничения

- `1 ≤ intervals.length ≤ 10⁴`
- `0 ≤ startᵢ < endᵢ ≤ 10⁶`
