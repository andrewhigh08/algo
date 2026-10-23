# 261. Graph Valid Tree

**Является ли граф деревом** · 🟡 Medium · Graphs · [LeetCode 🔒](https://leetcode.com/problems/graph-valid-tree/description/) · [LintCode 178](https://www.lintcode.com/problem/178/)

## Условие

Дан граф из `n` узлов, пронумерованных от `0` до `n - 1`. Даны целое число `n` и список рёбер `edges`, где `edges[i] = [aᵢ, bᵢ]` означает, что в графе есть **неориентированное** ребро между узлами `aᵢ` и `bᵢ`.

Верните `true`, *если рёбра данного графа образуют допустимое дерево*, и `false` *в противном случае*.

**Примечание:** повторяющихся рёбер в `edges` нет. Поскольку все рёбра неориентированные, `[0, 1]` — то же самое, что `[1, 0]`, поэтому вместе в `edges` они не встретятся.

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2021/03/12/tree1-graph.jpg)

```text
Input: n = 5, edges = [[0,1],[0,2],[0,3],[1,4]]
Output: true
```

**Пример 2:**

![Пример 2](https://assets.leetcode.com/uploads/2021/03/12/tree2-graph.jpg)

```text
Input: n = 5, edges = [[0,1],[1,2],[2,3],[1,3],[1,4]]
Output: false
```

## Ограничения

- `1 ≤ n ≤ 2000`
- `0 ≤ edges.length ≤ 5000`
- `edges[i].length == 2`
- `0 ≤ aᵢ, bᵢ < n`
- `aᵢ ≠ bᵢ`
- Петель и повторяющихся рёбер нет.
