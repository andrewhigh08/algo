# 212. [Поиск слов II (Word Search II)](https://leetcode.com/problems/word-search-ii/description/)

`Hard`

Дана доска `board` размером `m x n`, состоящая из символов, и список строк `words`, верните все слова из `words`, которые есть на доске.

Каждое слово должно быть составлено из букв последовательно соседних ячеек, где **соседние ячейки** - это ячейки, расположенные горизонтально или вертикально рядом. Одна и та же буква ячейки не может быть использована в слове более одного раза.

**Пример 1:**\
![](https://assets.leetcode.com/uploads/2020/11/07/search1.jpg)
```
Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
Output: ["eat","oath"]
```

**Пример 2:**\
![](https://assets.leetcode.com/uploads/2020/11/07/search2.jpg)
```
Input: board = [["a","b"],["c","d"]], words = ["abcb"]
Output: []
```

**Ограничения:**

*   `m == board.length`
*   `n == board[i].length`
*   1 ≤ m, n ≤ 12
*   `board[i][j]` - строчная английская буква.
*   1 ≤ `words.length` ≤ 3 * 10⁴
*   1 ≤ `words[i].length` ≤ 10
*   `words[i]` состоит из строчных английских букв.
*   Все строки в `words` уникальны.