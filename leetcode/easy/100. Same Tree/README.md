# 100. Same Tree

**Одинаковые деревья** · 🟢 Easy · Trees · [LeetCode](https://leetcode.com/problems/same-tree/description/)

## Условие

Даны корни двух бинарных деревьев `p` и `q`. Напишите функцию, которая проверяет, одинаковы они или нет.

Два бинарных дерева считаются одинаковыми, если они структурно идентичны и их узлы имеют одинаковые значения.

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2020/12/20/ex1.jpg)

```text
Input: p = [1,2,3], q = [1,2,3]
Output: true
```

**Пример 2:**

![Пример 2](https://assets.leetcode.com/uploads/2020/12/20/ex2.jpg)

```text
Input: p = [1,2], q = [1,null,2]
Output: false
```

**Пример 3:**

![Пример 3](https://assets.leetcode.com/uploads/2020/12/20/ex3.jpg)

```text
Input: p = [1,2,1], q = [1,1,2]
Output: false
```

## Ограничения

- Количество узлов в обоих деревьях — в диапазоне `[0, 100]`.
- `-10⁴ ≤ Node.val ≤ 10⁴`
