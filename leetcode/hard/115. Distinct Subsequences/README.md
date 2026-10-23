# 115. Distinct Subsequences

**Различные подпоследовательности** · 🔴 Hard · 2-D Dynamic Programming · [LeetCode](https://leetcode.com/problems/distinct-subsequences/description/)

## Условие

Даны две строки `s` и `t`. Верните *количество различных **подпоследовательностей** строки `s`, равных `t`*.

Тесты составлены так, что ответ помещается в 32-битное знаковое целое число.

## Примеры

**Пример 1:**

```text
Input: s = "rabbbit", t = "rabbit"
Output: 3
Explanation:
Как показано ниже, получить "rabbit" из s можно 3 способами.
rabbbit
^^^^ ^^
rabbbit
^^ ^^^^
rabbbit
^^^ ^^^
```

**Пример 2:**

```text
Input: s = "babgbag", t = "bag"
Output: 5
Explanation:
Как показано ниже, получить "bag" из s можно 5 способами.
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

## Ограничения

- `1 ≤ s.length, t.length ≤ 1000`
- `s` и `t` состоят из английских букв.
