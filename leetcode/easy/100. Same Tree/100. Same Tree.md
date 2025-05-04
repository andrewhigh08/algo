# 100. [Одинаковое дерево (Same Tree)](https://leetcode.com/problems/same-tree/description/)

`Easy`

Даны корни двух бинарных деревьев `p` и `q`, напишите функцию, чтобы проверить, одинаковы они или нет.

Два бинарных дерева считаются одинаковыми, если они структурно идентичны, и узлы имеют одинаковое значение.

**Пример 1:**\
![](https://assets.leetcode.com/uploads/2020/12/20/ex1.jpg)
```
Input: p = [1,2,3], q = [1,2,3]
Output: true
```

**Пример 2:**\
![](https://assets.leetcode.com/uploads/2020/12/20/ex2.jpg)
```
Input: p = [1,2], q = [1,null,2]
Output: false
```

**Пример 3:**\
![](https://assets.leetcode.com/uploads/2020/12/20/ex3.jpg)
```
Input: p = [1,2,1], q = [1,1,2]
Output: false
```

**Ограничения:**

*   Количество узлов в обоих деревьях находится в диапазоне [0, 100].
*   -10⁴ ≤ `Node.val` ≤ 10⁴