# 146. [LRU кэш (LRU Cache)](https://leetcode.com/problems/lru-cache/description/)

`Medium`

Разработайте структуру данных, которая соответствует ограничениям кэша **Least Recently Used (LRU)**.

Реализуйте класс `LRUCache`:

*   `LRUCache(int capacity)` Инициализирует LRU кэш с **положительным** размером `capacity`.
*   `int get(int key)` Возвращает значение `key`, если ключ существует, иначе возвращает -1.
*   `void put(int key, int value)` Обновляет значение `key`, если `key` существует. В противном случае добавьте пару ключ-значение в кэш. Если количество ключей превышает `capacity` в результате этой операции, **вытесните** наименее недавно использованный ключ.

Функции `get` и `put` должны выполняться со средней временной сложностью O(1).

**Пример 1:**
```
Input
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
Output
[null, null, null, 1, null, -1, null, -1, 3, 4]

Explanation
LRUCache lRUCache = new LRUCache(2);
lRUCache.put(1, 1); // cache is {1=1}
lRUCache.put(2, 2); // cache is {1=1, 2=2}
lRUCache.get(1);    // return 1
lRUCache.put(3, 3); // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
lRUCache.get(2);    // returns -1 (not found)
lRUCache.put(4, 4); // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
lRUCache.get(1);    // return -1 (not found)
lRUCache.get(3);    // return 3
lRUCache.get(4);    // return 4
```

**Ограничения:**

*   1 ≤ `capacity` ≤ 3000
*   0 ≤ `key` ≤ 10⁴
*   0 ≤ `value` ≤ 10⁵
*   Будет сделано не более 2 * 10⁵ вызовов к `get` и `put`.