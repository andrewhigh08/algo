# 72. Edit Distance

**Редакционное расстояние** · 🟡 Medium · 2-D Dynamic Programming · [LeetCode](https://leetcode.com/problems/edit-distance/description/)

## Условие

Даны две строки `word1` и `word2`. Верните *минимальное количество операций, необходимых для преобразования `word1` в `word2`*.

Над словом разрешены следующие три операции:

- Вставить символ
- Удалить символ
- Заменить символ

## Примеры

**Пример 1:**

```text
Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (заменить 'h' на 'r')
rorse -> rose (удалить 'r')
rose -> ros (удалить 'e')
```

**Пример 2:**

```text
Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (удалить 't')
inention -> enention (заменить 'i' на 'e')
enention -> exention (заменить 'n' на 'x')
exention -> exection (заменить 'n' на 'c')
exection -> execution (вставить 'u')
```

## Ограничения

- `0 ≤ word1.length, word2.length ≤ 500`
- `word1` и `word2` состоят из строчных английских букв.
