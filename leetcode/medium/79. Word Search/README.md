# 79. [Поиск слова (Word Search)](https://leetcode.com/problems/word-search/description/)

`Medium`

Дана сетка символов `board` размером `m x n` и строка `word`, верните `true`, если `word` существует в сетке.

Слово может быть составлено из букв последовательно соседних ячеек, где соседние ячейки - это ячейки, расположенные горизонтально или вертикально рядом. Одна и та же ячейка с буквой не может быть использована более одного раза.

**Пример 1:**\
![](https://assets.leetcode.com/uploads/2020/11/04/word2.jpg)
```
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true
```

**Пример 2:**\
![](https://assets.leetcode.com/uploads/2020/11/04/word-1.jpg)
```
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
Output: true
```

**Пример 3:**\
![](https://assets.leetcode.com/uploads/2020/10/15/word3.jpg)
```
Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
Output: false
```

**Ограничения:**

*   `m == board.length`
*   `n == board[i].length`
*   1 ≤ m, n ≤ 6
*   1 ≤ `word.length` ≤ 15
*   `board` и `word` состоят только из строчных и прописных английских букв.

**Дополнительный вопрос:** Можете ли вы использовать отсечение поиска, чтобы сделать ваше решение быстрее с большей `board`?
```