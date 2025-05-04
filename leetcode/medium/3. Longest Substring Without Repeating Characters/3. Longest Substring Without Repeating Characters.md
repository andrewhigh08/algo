# 3. [Самая длинная подстрока без повторяющихся символов (Longest Substring Without Repeating Characters)](https://leetcode.com/problems/longest-substring-without-repeating-characters/description/)

сложность: `Medium`

Дана строка `s`, найдите длину самой длинной подстроки без повторяющихся символов.

```
Пример
Input: s = "abcabcbb"
Output: 3
Объяснение: The answer is "abc", with the length of 3.

Пример
Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Пример
Input: s = "pwwkew"
Output: 3
Объяснение: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

**Ограничения:**

*   `0 ≤ s.length ≤ 5 * 10⁴`
*   `s` состоит из английских букв, цифр, символов и пробелов.