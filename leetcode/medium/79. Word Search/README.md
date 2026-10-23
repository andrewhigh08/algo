# 79. Word Search

**Поиск слова** · 🟡 Medium · Backtracking · [LeetCode](https://leetcode.com/problems/word-search/description/)

## Условие

Дана сетка символов `board` размером `m x n` и строка `word`. Верните `true`, *если `word` существует в сетке*.

Слово может быть составлено из букв последовательно соседних ячеек, где соседние ячейки — это ячейки, расположенные рядом по горизонтали или по вертикали. Одна и та же ячейка с буквой не может быть использована более одного раза.

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2020/11/04/word2.jpg)

```text
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true
```

**Пример 2:**

![Пример 2](https://assets.leetcode.com/uploads/2020/11/04/word-1.jpg)

```text
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
Output: true
```

**Пример 3:**

![Пример 3](https://assets.leetcode.com/uploads/2020/10/15/word3.jpg)

```text
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
Output: false
```

## Ограничения

- `m == board.length`
- `n == board[i].length`
- `1 ≤ m, n ≤ 6`
- `1 ≤ word.length ≤ 15`
- `board` и `word` состоят только из строчных и прописных английских букв.

## Дополнительный вопрос

Можете ли вы использовать отсечение перебора, чтобы ваше решение работало быстрее на доске `board` большего размера?
