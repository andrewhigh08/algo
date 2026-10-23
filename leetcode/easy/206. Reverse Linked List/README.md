# 206. Reverse Linked List

**Развернуть связный список** · 🟢 Easy · Linked List · [LeetCode](https://leetcode.com/problems/reverse-linked-list/description/)

## Условие

Дана голова `head` односвязного списка. Разверните список и верните *развёрнутый список*.

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2021/02/19/rev1ex1.jpg)

```text
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]
```

**Пример 2:**

![Пример 2](https://assets.leetcode.com/uploads/2021/02/19/rev1ex2.jpg)

```text
Input: head = [1,2]
Output: [2,1]
```

**Пример 3:**

```text
Input: head = []
Output: []
```

## Ограничения

- Количество узлов в списке — в диапазоне `[0, 5000]`.
- `-5000 ≤ Node.val ≤ 5000`

## Дополнительный вопрос

Связный список можно развернуть итеративно или рекурсивно. Сможете реализовать оба способа?
