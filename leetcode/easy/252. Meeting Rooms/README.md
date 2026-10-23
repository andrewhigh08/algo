# 252. Meeting Rooms

**Комнаты для переговоров** · 🟢 Easy · Intervals · [LeetCode 🔒](https://leetcode.com/problems/meeting-rooms/description/) · [LintCode 920](https://www.lintcode.com/problem/920/)

## Условие

Дан массив встреч `intervals`, где `intervals[i] = [startᵢ, endᵢ]`.

Человек может посетить все встречи, если никакие два интервала не пересекаются. Встречи, одна из которых заканчивается в момент `t`, а другая начинается в момент `t`, **не** пересекаются.

Верните `true`, если человек может посетить все встречи, и `false` в противном случае.

## Примеры

**Пример 1:**

```text
Input: intervals = [[0,30],[5,10],[15,20]]
Output: false
Explanation: встреча [0,30] пересекается и с [5,10], и с [15,20].
```

**Пример 2:**

```text
Input: intervals = [[7,10],[2,4]]
Output: true
```

## Ограничения

- `0 ≤ intervals.length ≤ 10⁴`
- `intervals[i].length == 2`
- `0 ≤ startᵢ < endᵢ ≤ 10⁶`
