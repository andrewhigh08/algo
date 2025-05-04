# 127. [Лестница Слов (Word Ladder)](https://leetcode.com/problems/word-ladder/description/)

`Hard`

**Последовательность преобразований** от слова `beginWord` к слову `endWord` с использованием словаря `wordList` - это последовательность слов `beginWord -> s₁ -> s₂ -> ... -> sk` такая, что:

*   Каждая соседняя пара слов отличается одной буквой.
*   Каждое `si` для `1 <= i <= k` находится в `wordList`. Обратите внимание, что `beginWord` не обязательно должен быть в `wordList`.
*   `sk == endWord`

Даны два слова, `beginWord` и `endWord`, и словарь `wordList`, верните *количество слов в кратчайшей последовательности преобразований* от `beginWord` до `endWord`, или `0`, если такой последовательности не существует.

**Пример 1:**
```
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Explanation: Одна из кратчайших последовательностей преобразований - "hit" -> "hot" -> "dot" -> "dog" -> "cog", которая состоит из 5 слов.
```

**Пример 2:**
```
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
Output: 0
Explanation: Слово "cog" отсутствует в wordList, поэтому допустимой последовательности преобразований не существует.
```

**Ограничения:**

*   1 ≤ `beginWord.length` ≤ 10
*   `endWord.length == beginWord.length`
*   1 ≤ `wordList.length` ≤ 5000
*   `wordList[i].length == beginWord.length`
*   `beginWord`, `endWord` и `wordList[i]` состоят из строчных английских букв.
*   `beginWord != endWord`
*   Все слова в `wordList` уникальны.