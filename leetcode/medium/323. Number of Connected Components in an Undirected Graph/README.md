# 323. Number of Connected Components in an Undirected Graph

**Количество компонент связности в неориентированном графе** · 🟡 Medium · Graphs · [LeetCode 🔒](https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/description/) · [LintCode 3651](https://www.lintcode.com/problem/3651/)

## Условие

Дан граф из `n` узлов, пронумерованных от `0` до `n - 1`. Даны целое число `n` и массив `edges`, где `edges[i] = [aᵢ, bᵢ]` означает, что в графе есть ребро между `aᵢ` и `bᵢ`.

Верните *количество компонент связности в графе*.

**Примечание:** повторяющихся рёбер в `edges` нет. Поскольку все рёбра неориентированные, `[0, 1]` — то же самое, что `[1, 0]`, поэтому вместе в `edges` они не встретятся.

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2021/03/14/conn1-graph.jpg)

```text
Input: n = 5, edges = [[0,1],[1,2],[3,4]]
Output: 2
```

**Пример 2:**

![Пример 2](https://assets.leetcode.com/uploads/2021/03/14/conn2-graph.jpg)

```text
Input: n = 5, edges = [[0,1],[1,2],[2,3],[3,4]]
Output: 1
```

## Ограничения

- `1 ≤ n ≤ 2000`
- `1 ≤ edges.length ≤ 5000`
- `edges[i] = [aᵢ, bᵢ]`
- `aᵢ ≠ bᵢ`
- Повторяющихся рёбер нет.
