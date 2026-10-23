# 19. Remove Nth Node From End of List

**Удалить N-й узел с конца списка** · 🟡 Medium · Linked List · [LeetCode](https://leetcode.com/problems/remove-nth-node-from-end-of-list/description/)

## Условие

Дана голова `head` связного списка. Удалите `n`-й узел с конца списка и верните голову списка.

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2020/10/03/remove_ex1.jpg)

```text
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
```

**Пример 2:**

```text
Input: head = [1], n = 1
Output: []
```

**Пример 3:**

```text
Input: head = [1,2], n = 1
Output: [1]
```

## Ограничения

- Количество узлов в списке равно `sz`.
- `1 ≤ sz ≤ 30`
- `0 ≤ Node.val ≤ 100`
- `1 ≤ n ≤ sz`

## Дополнительный вопрос

Сможете ли вы сделать это за один проход?
