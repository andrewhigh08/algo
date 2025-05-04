143. [Изменить порядок списка (Reorder List)](https://leetcode.com/problems/reorder-list/description/)

`Medium`

Вам дан `head` односвязного списка. Список может быть представлен как:
L₀ → L₁ → ... → Lₙ₋₁ → Lₙ

Измените порядок списка так, чтобы он имел следующий вид:
L₀ → Lₙ → L₁ → Lₙ₋₁ → L₂ → Lₙ₋₂ → ...

Вы не можете изменять значения в узлах списка. Могут быть изменены только сами узлы.

**Пример 1:**\
![Пример 1](https://assets.leetcode.com/uploads/2021/03/04/reorder1linked-list.jpg)
```
Input: head = [1,2,3,4]
Output: [1,4,2,3]
```

**Пример 2:**\
![Пример 2](https://assets.leetcode.com/uploads/2021/03/09/reorder2-linked-list.jpg)
```
Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]
```

**Ограничения:**

*   Количество узлов в списке находится в диапазоне [1, 5 * 10⁴].
*   1 ≤ `Node.val` ≤ 1000