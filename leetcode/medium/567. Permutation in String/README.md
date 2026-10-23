# 567. Permutation in String

**Перестановка в строке** · 🟡 Medium · Sliding Window · [LeetCode](https://leetcode.com/problems/permutation-in-string/description/)

## Условие

Даны две строки `s1` и `s2`. Верните `true`, если `s2` содержит перестановку `s1`, или `false` в противном случае.

Другими словами, верните `true`, если одна из перестановок `s1` является подстрокой `s2`.

## Примеры

**Пример 1:**

```text
Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 содержит одну из перестановок s1 ("ba").
```

**Пример 2:**

```text
Input: s1 = "ab", s2 = "eidboaoo"
Output: false
```

## Ограничения

- `1 ≤ s1.length, s2.length ≤ 10⁴`
- `s1` и `s2` состоят из строчных английских букв.
