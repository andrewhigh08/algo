# 40. [Сумма комбинаций II (Combination Sum II)](https://leetcode.com/problems/combination-sum-ii/description/)

`Medium`

Дана коллекция чисел-кандидатов (`candidates`) и целевое число (`target`), найдите все уникальные комбинации в `candidates`, где числа-кандидаты суммируются в `target`.

Каждое число в `candidates` может быть использовано только **один раз** в комбинации.

**Примечание:** Набор решений не должен содержать дублирующихся комбинаций.

**Пример 1:**
```
Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
```

**Пример 2:**
```
Input: candidates = [2,5,2,1,2], target = 5
Output:
[
[1,2,2],
[5]
]
```

**Ограничения:**

*   1 ≤ `candidates.length` ≤ 100
*   1 ≤ `candidates[i]` ≤ 50
*   1 ≤ `target` ≤ 30