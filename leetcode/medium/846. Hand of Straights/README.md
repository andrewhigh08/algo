# 846. Hand of Straights

**Разбиение карт на стриты** · 🟡 Medium · Greedy · [LeetCode](https://leetcode.com/problems/hand-of-straights/description/)

## Условие

У Алисы есть некоторое количество карт, и она хочет разложить их на группы так, чтобы каждая группа имела размер `groupSize` и состояла из `groupSize` последовательных карт.

Дан целочисленный массив `hand`, где `hand[i]` — значение, написанное на `i`-й карте, и целое число `groupSize`. Верните `true`, если она может так разложить карты, или `false` в противном случае.

**Примечание:** эта задача совпадает с задачей [1296. Divide Array in Sets of K Consecutive Numbers](https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/).

## Примеры

**Пример 1:**

```text
Input: hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
Output: true
Explanation: руку Алисы можно разложить на группы [1,2,3],[2,3,4],[6,7,8].
```

**Пример 2:**

```text
Input: hand = [1,2,3,4,5], groupSize = 4
Output: false
Explanation: руку Алисы нельзя разложить на группы по 4.
```

## Ограничения

- `1 ≤ hand.length ≤ 10⁴`
- `0 ≤ hand[i] ≤ 10⁹`
- `1 ≤ groupSize ≤ hand.length`
