# 3. Longest Substring Without Repeating Characters

**Самая длинная подстрока без повторяющихся символов** · 🟡 Medium · Sliding Window · [LeetCode](https://leetcode.com/problems/longest-substring-without-repeating-characters/description/)

## Условие

Дана строка `s`. Найдите длину **самой длинной подстроки** без повторяющихся символов.

## Примеры

**Пример 1:**

```text
Input: s = "abcabcbb"
Output: 3
Explanation: ответ — "abc", его длина равна 3. Заметьте, что "bca" и "cab" — тоже правильные ответы.
```

**Пример 2:**

```text
Input: s = "bbbbb"
Output: 1
Explanation: ответ — "b", его длина равна 1.
```

**Пример 3:**

```text
Input: s = "pwwkew"
Output: 3
Explanation: ответ — "wke", его длина равна 3.
Обратите внимание, что ответ должен быть подстрокой: "pwke" — это подпоследовательность, а не подстрока.
```

## Ограничения

- `0 ≤ s.length ≤ 10⁵`
- `s` состоит из английских букв, цифр, символов и пробелов.
