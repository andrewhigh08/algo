# 424. [Самая длинная подстрока с повторяющимися символами после замены (Longest Repeating Character Replacement)](https://leetcode.com/problems/longest-substring-without-repeating-characters/description/)

сложность: `Medium`

Вам дана строка `s` и целое число `k`. Вы можете выбрать любой символ строки и изменить его на любой другой символ английского алфавита в верхнем регистре. Вы можете выполнить эту операцию не более `k` раз.

Верните длину самой длинной подстроки, содержащей один и тот же символ, которую вы можете получить после выполнения вышеуказанных операций.

```
Пример 1:

Input: s = "abcabcbb"
Output: 3
Объяснение: The answer is "abc", with the length of 3.

Пример 2:

Input: s = "bbbbb"
Output: 1
Объяснение: The answer is "b", with the length of 1.

Пример 3:

Input: s = "pwwkew"
Output: 3
Объяснение: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
```
**Ограничения:**

*   `1 ≤ s.length ≤ 10⁵`
*   `s` состоит только из английских букв в верхнем регистре.
*   `0 ≤ k ≤ s.length`