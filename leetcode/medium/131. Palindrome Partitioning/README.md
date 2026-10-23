# 131. Palindrome Partitioning

**Разбиение на палиндромы** · 🟡 Medium · Backtracking · [LeetCode](https://leetcode.com/problems/palindrome-partitioning/description/)

## Условие

Дана строка `s`. Разбейте `s` так, чтобы каждая подстрока разбиения была **палиндромом**. Верните *все возможные разбиения `s` на палиндромы*.

## Примеры

**Пример 1:**

```text
Input: s = "aab"
Output: [["a","a","b"],["aa","b"]]
```

**Пример 2:**

```text
Input: s = "a"
Output: [["a"]]
```

## Ограничения

- `1 ≤ s.length ≤ 16`
- `s` содержит только строчные английские буквы.
