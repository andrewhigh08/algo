# 125. [Проверка на палиндром (Valid Palindrome)](https://leetcode.com/problems/valid-palindrome/description/)

сложность: `Easy`

Фраза является **палиндромом**, если после преобразования всех заглавных букв в строчные и удаления всех не буквенно-цифровых символов она читается одинаково как вперед, так и назад. Буквенно-цифровые символы включают буквы и цифры.

Дана строка `s`, верните `true`, если она является палиндромом, или `false` в противном случае.

```
Пример 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Пример 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.

Пример 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
```
**Ограничения:**

*   `1 ≤ s.length ≤ 2 * 10⁵`
*   `s` состоит только из печатных символов ASCII.