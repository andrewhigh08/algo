# 72. [Расстояние Редактирования (Edit Distance)](https://leetcode.com/problems/edit-distance/description/)

`Medium`

Даны две строки `word1` и `word2`, верните *минимальное количество операций, необходимое для преобразования `word1` в `word2`*.

У вас есть следующие три операции, разрешенные для слова:

*   Вставить символ
*   Удалить символ
*   Заменить символ

**Пример 1:**
```
Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (заменить 'h' на 'r')
rorse -> rose (удалить 'r')
rose -> ros (удалить 'e')
```

**Пример 2:**
```
Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (удалить 't')
inention -> enention (заменить 'i' на 'e')
enention -> exention (заменить 'n' на 'x')
exention -> exection (заменить 'n' на 'c')
exection -> execution (вставить 'u')
```

**Ограничения:**

*   0 ≤ `word1.length`, `word2.length` ≤ 500
*   `word1` и `word2` состоят из строчных английских букв.