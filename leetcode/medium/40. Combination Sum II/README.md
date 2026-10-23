# 40. Combination Sum II

**Сумма комбинаций II** · 🟡 Medium · Backtracking · [LeetCode](https://leetcode.com/problems/combination-sum-ii/description/)

## Условие

Дана коллекция чисел-кандидатов (`candidates`) и целевое число (`target`). Найдите все уникальные комбинации чисел из `candidates`, в которых числа-кандидаты в сумме дают `target`.

Каждое число из `candidates` может быть использовано в комбинации только **один раз**.

**Примечание:** набор решений не должен содержать дублирующихся комбинаций.

## Примеры

**Пример 1:**

```text
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

```text
Input: candidates = [2,5,2,1,2], target = 5
Output:
[
[1,2,2],
[5]
]
```

## Ограничения

- `1 ≤ candidates.length ≤ 100`
- `1 ≤ candidates[i] ≤ 50`
- `1 ≤ target ≤ 30`
