# 17. Letter Combinations of a Phone Number

**Буквенные комбинации телефонного номера** · 🟡 Medium · Backtracking · [LeetCode](https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/)

## Условие

Дана строка, содержащая цифры от `2` до `9` включительно. Верните все возможные буквенные комбинации, которые может представлять это число. Ответ можно вернуть в **любом порядке**.

Сопоставление цифр буквам (как на кнопках телефона) приведено ниже. Обратите внимание, что цифра `1` не сопоставляется ни с какими буквами.

![Иллюстрация](https://assets.leetcode.com/uploads/2022/03/15/1200px-telephone-keypad2svg.png)

## Примеры

**Пример 1:**

```text
Input: digits = "23"
Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
```

**Пример 2:**

```text
Input: digits = "2"
Output: ["a","b","c"]
```

## Ограничения

- `1 ≤ digits.length ≤ 4`
- `digits[i]` — цифра в диапазоне `['2', '9']`.
