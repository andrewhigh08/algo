# 211. [Дизайн структуры данных для добавления и поиска слов (Design Add and Search Words Data Structure)](https://leetcode.com/problems/design-add-and-search-words-data-structure/description/)

`Medium`

Разработайте структуру данных, которая поддерживает добавление новых слов и проверку, соответствует ли строка какой-либо ранее добавленной строке.

Реализуйте класс `WordDictionary`:

*   `WordDictionary()` Инициализирует объект.
*   `void addWord(word)` Добавляет `word` в структуру данных, чтобы его можно было сопоставить позже.
*   `bool search(word)` Возвращает `true`, если в структуре данных есть какая-либо строка, которая соответствует `word`, или `false` в противном случае. `word` может содержать точки `.`, где точки могут быть сопоставлены с любой буквой.

**Пример:**
```
Input:
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
Output:
[null,null,null,null,false,true,true,true]

Explanation:
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // return False
wordDictionary.search("bad"); // return True
wordDictionary.search(".ad"); // return True
wordDictionary.search("b.."); // return True
```

**Ограничения:**

*   1 ≤ `word.length` ≤ 25
*   `word` в `addWord` состоит из строчных английских букв.
*   `word` в `search` состоит из `.` или строчных английских букв.
*   В `word` для поисковых запросов будет не более 2 точек.
*   Будет сделано не более 10⁴ вызовов к `addWord` и `search`.