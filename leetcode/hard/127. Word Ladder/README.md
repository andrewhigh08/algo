# 127. Word Ladder

**Лестница слов** · 🔴 Hard · Graphs · [LeetCode](https://leetcode.com/problems/word-ladder/description/)

## Условие

**Последовательность преобразований** от слова `beginWord` к слову `endWord` с использованием словаря `wordList` — это последовательность слов `beginWord -> s₁ -> s₂ -> ... -> sₖ`, такая что:

- каждая пара соседних слов отличается ровно одной буквой;
- каждое `sᵢ` для `1 ≤ i ≤ k` есть в `wordList`. Обратите внимание, что `beginWord` не обязательно должно быть в `wordList`;
- `sₖ == endWord`.

Даны два слова, `beginWord` и `endWord`, и словарь `wordList`. Верните **количество слов** в **кратчайшей последовательности преобразований** от `beginWord` до `endWord` или `0`, если такой последовательности не существует.

## Примеры

**Пример 1:**

```text
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Explanation: одна из кратчайших последовательностей преобразований — "hit" -> "hot" -> "dot" -> "dog" -> "cog", она состоит из 5 слов.
```

**Пример 2:**

```text
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log"]
Output: 0
Explanation: слова endWord "cog" нет в wordList, поэтому допустимой последовательности преобразований не существует.
```

## Ограничения

- `1 ≤ beginWord.length ≤ 10`
- `endWord.length == beginWord.length`
- `1 ≤ wordList.length ≤ 5000`
- `wordList[i].length == beginWord.length`
- `beginWord`, `endWord` и `wordList[i]` состоят из строчных английских букв.
- `beginWord ≠ endWord`
- Все слова в `wordList` **уникальны**.
