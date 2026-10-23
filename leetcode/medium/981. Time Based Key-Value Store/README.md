# 981. Time Based Key-Value Store

**Хранилище ключ-значение на основе времени** · 🟡 Medium · Binary Search · [LeetCode](https://leetcode.com/problems/time-based-key-value-store/description/)

## Условие

Спроектируйте структуру данных «ключ-значение» на основе времени, которая может хранить несколько значений для одного и того же ключа с разными метками времени и получать значение ключа на определённый момент времени.

Реализуйте класс `TimeMap`:

- `TimeMap()` — инициализирует объект структуры данных.
- `void set(String key, String value, int timestamp)` — сохраняет ключ `key` со значением `value` в момент времени `timestamp`.
- `String get(String key, int timestamp)` — возвращает значение, для которого ранее был вызван `set` с тем же `key` и меткой времени `timestamp_prev ≤ timestamp`. Если таких значений несколько, возвращает значение с наибольшим `timestamp_prev`. Если значений нет, возвращает `""`.

## Примеры

**Пример 1:**

```text
Input
["TimeMap", "set", "get", "get", "set", "get", "get"]
[[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
Output
[null, null, "bar", "bar", null, "bar2", "bar2"]

Explanation
TimeMap timeMap = new TimeMap();
timeMap.set("foo", "bar", 1);  // сохраняем ключ "foo" со значением "bar" и меткой времени timestamp = 1
timeMap.get("foo", 1);         // возвращает "bar"
timeMap.get("foo", 3);         // возвращает "bar": для "foo" нет значений с метками 3 и 2, единственное значение — "bar" с меткой 1
timeMap.set("foo", "bar2", 4); // сохраняем ключ "foo" со значением "bar2" и меткой времени timestamp = 4
timeMap.get("foo", 4);         // возвращает "bar2"
timeMap.get("foo", 5);         // возвращает "bar2"
```

## Ограничения

- `1 ≤ key.length, value.length ≤ 100`
- `key` и `value` состоят из строчных английских букв и цифр.
- `1 ≤ timestamp ≤ 10⁷`
- Все метки времени `timestamp` в вызовах `set` строго возрастают.
- Будет сделано не более `2 * 10⁵` вызовов `set` и `get`.
