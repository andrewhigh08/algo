# 115. [Различные Подпоследовательности (Distinct Subsequences)](https://leetcode.com/problems/distinct-subsequences/description/)

`Hard`

Даны две строки `s` и `t`, верните *количество **различных подпоследовательностей** `s`, которые равны `t`*.

Тестовые случаи сгенерированы так, что ответ помещается в 32-битное знаковое целое число.

**Пример 1:**
```
Input: s = "rabbbit", t = "rabbit"
Output: 3
Explanation:
Как показано ниже, есть 3 способа сгенерировать "rabbit" из s.
rabbbit
^^^^ ^^
rabbbit
^^ ^^^^
rabbbit
^^^ ^^^
```

**Пример 2:**
```
Input: s = "babgbag", t = "bag"
Output: 5
Explanation:
Как показано ниже, есть 5 способов сгенерировать "bag" из s.
babgbag
^^ ^
babgbag
^^    ^
babgbag
^    ^^
babgbag
  ^  ^^
babgbag
^^^
```

**Ограничения:**

*   1 ≤ `s.length`, `t.length` ≤ 1000
*   `s` и `t` состоят из английских букв.