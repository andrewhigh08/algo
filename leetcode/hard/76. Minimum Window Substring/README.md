# 76. [Минимальная подстрока окна (Minimum Window Substring)](https://leetcode.com/problems/minimum-window-substring/description/)

сложность: `Hard`

Даны две строки `s` и `t` длины `m` и `n` соответственно. Верните **минимальную подстроку-окно** строки `s` такую, что каждый символ строки `t` (включая дубликаты) содержится в окне. Если такой подстроки нет, верните пустую строку "".

Тестовые примеры будут сгенерированы таким образом, что ответ будет уникальным.

```
Пример 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

Пример 2:

Input: s = "a", t = "a"
Output: "a"
Explanation: The entire string s is the minimum window.

Пример 3:

Input: s = "a", t = "aa"
Output: ""
Explanation: Both 'a's from t must be included in the window.
Since the largest window of s only has one 'a', return empty string.
```
**Ограничения:**

*   `m == s.length`
*   `n == t.length`
*   `1 ≤ m, n ≤ 10⁵`
*   `s` и `t` состоят из английских букв в верхнем и нижнем регистре.

**Дополнительный вопрос:** Можете ли вы найти алгоритм, работающий за время O(m + n)?