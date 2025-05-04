# 206. [Развернуть связный список (Reverse Linked List)](https://leetcode.com/problems/reverse-linked-list/description/)

`Easy`

Дана `head` односвязного списка, разверните список и верните развернутый список.

**Пример 1:**\
![Пример 1](https://assets.leetcode.com/uploads/2021/02/19/rev1ex1.jpg)

```
Вход: head = [1,2,3,4,5]
Выход: [5,4,3,2,1]
```

**Пример 2:**\
![Пример 2](https://assets.leetcode.com/uploads/2021/02/19/rev1ex2.jpg)
```
Вход: head = [1,2]
Выход: [2,1]
```

```
Пример 3:

Вход: head = []
Выход: []
```

Ограничения:

*   Количество узлов в списке находится в диапазоне [0, 5000].
*   -5000 ≤ Node.val ≤ 5000