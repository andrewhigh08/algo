# 143. Reorder List

**Изменить порядок списка** · 🟡 Medium · Linked List · [LeetCode](https://leetcode.com/problems/reorder-list/description/)

## Условие

Дана голова `head` односвязного списка. Список можно представить так:

`L₀ → L₁ → … → Lₙ₋₁ → Lₙ`

*Измените порядок списка так, чтобы он имел следующий вид:*

`L₀ → Lₙ → L₁ → Lₙ₋₁ → L₂ → Lₙ₋₂ → …`

Изменять значения в узлах списка нельзя. Менять можно только сами узлы.

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2021/03/04/reorder1linked-list.jpg)

```text
Input: head = [1,2,3,4]
Output: [1,4,2,3]
```

**Пример 2:**

![Пример 2](https://assets.leetcode.com/uploads/2021/03/09/reorder2-linked-list.jpg)

```text
Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]
```

## Ограничения

- Количество узлов в списке — в диапазоне `[1, 5 * 10⁴]`.
- `1 ≤ Node.val ≤ 1000`
