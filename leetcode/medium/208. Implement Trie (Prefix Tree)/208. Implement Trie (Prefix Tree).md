# 208. [Реализовать Trie (Префиксное Дерево) (Implement Trie (Prefix Tree))]()

`Medium`

**Trie** (произносится как "трай") или **префиксное дерево** - это древовидная структура данных, используемая для эффективного хранения и извлечения ключей из набора строк. Существуют различные применения этой структуры данных, такие как автодополнение и проверка орфографии.

Реализуйте класс `Trie`:

*   `Trie()` Инициализирует объект trie.
*   `void insert(String word)` Вставляет строку `word` в trie.
*   `boolean search(String word)` Возвращает `true`, если строка `word` находится в trie (т.е. была вставлена ранее), и `false` в противном случае.
*   `boolean startsWith(String prefix)` Возвращает `true`, если есть ранее вставленная строка `word`, которая имеет префикс `prefix`, и `false` в противном случае.

**Пример 1:**
```
Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
Output
[null, null, true, false, true, null, true]

Explanation
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // return True
trie.search("app");     // return False
trie.startsWith("app"); // return True
trie.insert("app");
trie.search("app");     // return True
```

**Ограничения:**

*   1 ≤ `word.length`, `prefix.length` ≤ 2000
*   `word` и `prefix` состоят только из строчных английских букв.
*   Будет сделано не более 3 * 10⁴ вызовов к `insert`, `search` и `startsWith` в общей сложности.
```