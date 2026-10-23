# 212. Word Search II

**Поиск слов II** · 🔴 Hard · Tries · [LeetCode](https://leetcode.com/problems/word-search-ii/description/)

## Условие

Дана доска `board` размером `m x n`, состоящая из символов, и список строк `words`. Верните *все слова из `words`, которые есть на доске*.

Каждое слово должно быть составлено из букв последовательно соседних ячеек, где **соседние ячейки** — это ячейки, расположенные рядом по горизонтали или по вертикали. Одну и ту же ячейку нельзя использовать в слове более одного раза.

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2020/11/07/search1.jpg)

```text
Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
Output: ["eat","oath"]
```

**Пример 2:**

![Пример 2](https://assets.leetcode.com/uploads/2020/11/07/search2.jpg)

```text
Input: board = [["a","b"],["c","d"]], words = ["abcb"]
Output: []
```

## Ограничения

- `m == board.length`
- `n == board[i].length`
- `1 ≤ m, n ≤ 12`
- `board[i][j]` — строчная английская буква.
- `1 ≤ words.length ≤ 3 * 10⁴`
- `1 ≤ words[i].length ≤ 10`
- `words[i]` состоит из строчных английских букв.
- Все строки в `words` уникальны.
