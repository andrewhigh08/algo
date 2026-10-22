# 846. [Рука Стрейтов (Hand of Straights)](https://leetcode.com/problems/hand-of-straights/description/)

`Medium`

У Алисы есть некоторое количество карт, и она хочет переставить карты в группы так, чтобы каждая группа имела размер `groupSize` и состояла из `groupSize` последовательных карт.

Дан целочисленный массив `hand`, где `hand[i]` - это значение, написанное на i-й карте, и целое число `groupSize`. Верните `true`, если она может переставить карты, или `false` в противном случае.

**Пример 1:**
```
Input: hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
Output: true
Explanation: Руку Алисы можно переставить как [1,2,3], [2,3,4], [6,7,8].
```

**Пример 2:**
```
Input: hand = [1,2,3,4,5], groupSize = 4
Output: false
Explanation: Руку Алисы нельзя переставить в группы по 4.
```

**Ограничения:**

*   1 ≤ `hand.length` ≤ 10⁴
*   0 ≤ `hand[i]` ≤ 10⁹
*   1 ≤ `groupSize` ≤ `hand.length`

**Примечание:** Этот вопрос совпадает с 1296: [https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/](https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/)