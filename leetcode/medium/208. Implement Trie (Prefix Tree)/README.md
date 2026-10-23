# 208. Implement Trie (Prefix Tree)

**Реализовать Trie (префиксное дерево)** · 🟡 Medium · Tries · [LeetCode](https://leetcode.com/problems/implement-trie-prefix-tree/description/)

## Условие

[**Trie**](https://en.wikipedia.org/wiki/Trie) (произносится как «трай»), или **префиксное дерево**, — это древовидная структура данных, используемая для эффективного хранения и извлечения ключей из набора строк. У этой структуры данных есть различные применения, например автодополнение и проверка орфографии.

Реализуйте класс `Trie`:

- `Trie()` — инициализирует объект trie.
- `void insert(String word)` — вставляет строку `word` в trie.
- `boolean search(String word)` — возвращает `true`, если строка `word` есть в trie (то есть была вставлена ранее), и `false` в противном случае.
- `boolean startsWith(String prefix)` — возвращает `true`, если есть ранее вставленная строка `word`, которая имеет префикс `prefix`, и `false` в противном случае.

## Примеры

**Пример 1:**

```text
Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
Output
[null, null, true, false, true, null, true]

Explanation
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // возвращает True
trie.search("app");     // возвращает False
trie.startsWith("app"); // возвращает True
trie.insert("app");
trie.search("app");     // возвращает True
```

## Ограничения

- `1 ≤ word.length, prefix.length ≤ 2000`
- `word` и `prefix` состоят только из строчных английских букв.
- Будет сделано не более `3 * 10⁴` вызовов `insert`, `search` и `startsWith` **в сумме**.
