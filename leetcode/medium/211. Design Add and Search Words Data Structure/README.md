# 211. Design Add and Search Words Data Structure

**Структура данных для добавления и поиска слов** · 🟡 Medium · Tries · [LeetCode](https://leetcode.com/problems/design-add-and-search-words-data-structure/description/)

## Условие

Разработайте структуру данных, которая поддерживает добавление новых слов и проверку, соответствует ли строка какой-либо ранее добавленной строке.

Реализуйте класс `WordDictionary`:

- `WordDictionary()` — инициализирует объект.
- `void addWord(word)` — добавляет `word` в структуру данных, чтобы позже его можно было сопоставить.
- `bool search(word)` — возвращает `true`, если в структуре данных есть какая-либо строка, которая соответствует `word`, или `false` в противном случае. `word` может содержать точки `'.'`, и точка может соответствовать любой букве.

## Примеры

**Пример 1:**

```text
Input
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
Output
[null,null,null,null,false,true,true,true]

Explanation
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // возвращает False
wordDictionary.search("bad"); // возвращает True
wordDictionary.search(".ad"); // возвращает True
wordDictionary.search("b.."); // возвращает True
```

## Ограничения

- `1 ≤ word.length ≤ 25`
- `word` в `addWord` состоит из строчных английских букв.
- `word` в `search` состоит из `'.'` или строчных английских букв.
- В `word` для запросов `search` будет не более `2` точек.
- Будет сделано не более `10⁴` вызовов `addWord` и `search`.
