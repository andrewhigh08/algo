# 146. LRU Cache

**LRU-кэш** · 🟡 Medium · Linked List · [LeetCode](https://leetcode.com/problems/lru-cache/description/)

## Условие

Спроектируйте структуру данных, которая соблюдает ограничения [кэша **LRU** (Least Recently Used)](https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU).

Реализуйте класс `LRUCache`:

- `LRUCache(int capacity)` — инициализирует LRU-кэш с **положительным** размером `capacity`.
- `int get(int key)` — возвращает значение по ключу `key`, если ключ существует, иначе возвращает `-1`.
- `void put(int key, int value)` — обновляет значение `key`, если ключ существует. Иначе добавляет пару `key-value` в кэш. Если после этой операции число ключей превышает `capacity`, **вытесняет** ключ, который использовался давнее всех.

Функции `get` и `put` должны работать в среднем за `O(1)`.

## Примеры

**Пример 1:**

```text
Input
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, null, -1, 3, 4]

Explanation
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // кэш: {1=1}
lRUCache.put(2, 2); // кэш: {1=1, 2=2}
lRUCache.get(1);    // возвращает 1
lRUCache.put(3, 3); // давнее всех использовался ключ 2, вытесняем его, кэш: {1=1, 3=3}
lRUCache.get(2);    // возвращает -1 (не найден)
lRUCache.put(4, 4); // давнее всех использовался ключ 1, вытесняем его, кэш: {4=4, 3=3}
lRUCache.get(1);    // возвращает -1 (не найден)
lRUCache.get(3);    // возвращает 3
lRUCache.get(4);    // возвращает 4
```

## Ограничения

- `1 ≤ capacity ≤ 3000`
- `0 ≤ key ≤ 10⁴`
- `0 ≤ value ≤ 10⁵`
- Будет сделано не более `2 * 10⁵` вызовов `get` и `put`.
