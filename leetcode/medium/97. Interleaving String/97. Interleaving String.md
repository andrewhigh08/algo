# 97. [Чередующаяся Строка (Interleaving String)]()

`Medium`

Даны строки `s1`, `s2`, `s3`. Найдите, образована ли `s3` **чередованием** `s1` и `s2`.

**Чередование** двух строк `s` и `t` - это конфигурация, где `s` и `t` разделены на `n` и `m` **подстрок** соответственно, так что:

*   `s = s₁ + s₂ + ... + sₙ`
*   `t = t₁ + t₂ + ... + tₘ`
*   `|n - m| ≤ 1`
*   **Чередование** равно `s₁ + t₁ + s₂ + t₂ + s₃ + t₃ + ...` или `t₁ + s₁ + t₂ + s₂ + t₃ + s₃ + ...`

**Примечание:** `a + b` - это конкатенация строк `a` и `b`.

Example 1:\
![](https://assets.leetcode.com/uploads/2020/09/02/interleave.jpg)
```
Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
Output: true
Explanation: One way to obtain s3 is:
Split s1 into s1 = "aa" + "bc" + "c", and s2 into s2 = "dbbc" + "a".
Interleaving the two splits, we get "aa" + "dbbc" + "bc" + "a" + "c" = "aadbbcbcac".
Since s3 can be obtained by interleaving s1 and s2, we return true.

Example 2:

Input: s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
Output: false
Explanation: Notice how it is impossible to interleave s2 with any other string to obtain s3.

Example 3:

Input: s1 = "", s2 = "", s3 = ""
Output: true
```
**Ограничения:**

*   0 ≤ `s1.length`, `s2.length` ≤ 100
*   0 ≤ `s3.length` ≤ 200
*   `s1`, `s2` и `s3` состоят из строчных английских букв.

**Дополнительный вопрос:** Можете ли вы решить эту задачу, используя только `O(s2.length)` дополнительного пространства памяти?
