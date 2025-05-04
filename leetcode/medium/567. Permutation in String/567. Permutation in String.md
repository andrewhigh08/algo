# 567. [Перестановка в строке (Permutation in String)](https://leetcode.com/problems/permutation-in-string/description/)

сложность: `Medium`

Даны две строки `s1` и `s2`, верните `true`, если `s2` содержит перестановку `s1`, иначе верните `false`. Другими словами, верните `true`, если одна из перестановок `s1` является подстрокой `s2`.

```
Пример 1:

Input: s1 = "ab", s2 = "eidbaooo"
Output: true
Explanation: s2 contains one permutation of s1 ("ba").

Пример 2:

Input: s1 = "ab", s2 = "eidboaoo"
Output: false
```

**Ограничения:**

*   `1 ≤ s1.length, s2.length ≤ 10⁴`
*   `s1` и `s2` состоят из строчных английских букв.