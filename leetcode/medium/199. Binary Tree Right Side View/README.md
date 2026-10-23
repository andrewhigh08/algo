# 199. Binary Tree Right Side View

**Вид бинарного дерева справа** · 🟡 Medium · Trees · [LeetCode](https://leetcode.com/problems/binary-tree-right-side-view/description/)

## Условие

Дан корень бинарного дерева `root`. Представьте, что вы стоите с **правой стороны** от него, и верните *значения узлов, которые вы можете видеть, упорядоченные сверху вниз*.

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2024/11/24/tmpd5jn43fs-1.png)

```text
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]
```

**Пример 2:**

![Пример 2](https://assets.leetcode.com/uploads/2024/11/24/tmpkpe40xeh-1.png)

```text
Input: root = [1,2,3,4,null,null,null,5]
Output: [1,3,4,5]
```

**Пример 3:**

```text
Input: root = [1,null,3]
Output: [1,3]
```

**Пример 4:**

```text
Input: root = []
Output: []
```

## Ограничения

- Количество узлов в дереве — в диапазоне `[0, 100]`.
- `-100 ≤ Node.val ≤ 100`
