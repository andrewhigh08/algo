# 981. [Хранилище Ключ-Значение на основе времени (Time Based Key-Value Store)](https://leetcode.com/problems/time-based-key-value-store/description/)

сложность: `Medium`

Разработайте структуру данных ключ-значение на основе времени, которая может хранить несколько значений для одного и того же ключа в разные моменты времени (timestamp) и получать значение ключа в определенный момент времени.

Реализуйте класс `TimeMap`:

*   `TimeMap()` Инициализирует объект структуры данных.
*   `void set(String key, String value, int timestamp)` Сохраняет ключ `key` со значением `value` в заданный момент времени `timestamp`.
*   `String get(String key, int timestamp)` Возвращает значение `value` такое, что ранее был вызван `set` с тем же `key` и моментом времени `timestamp_prev <= timestamp`. Если таких значений несколько, возвращается значение, связанное с наибольшим `timestamp_prev`. Если значений нет, возвращается "".

Обратите внимание, что все моменты времени `timestamp` для вызовов `set` строго возрастают.

```
Пример 1:

Input
["TimeMap", "set", "get", "get", "set", "get", "get"]
[[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
Output
[null, null, "bar", "bar", null, "bar2", "bar2"]

Объяснение
TimeMap timeMap = new TimeMap();
timeMap.set("foo", "bar", 1);  // store the key "foo" and value "bar" along with timestamp = 1.
timeMap.get("foo", 1);         // return "bar"
timeMap.get("foo", 3);         // return "bar", since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
timeMap.set("foo", "bar2", 4); // store the key "foo" and value "bar2" along with timestamp = 4.
timeMap.get("foo", 4);         // return "bar2"
timeMap.get("foo", 5);         // return "bar2"
```

**Ограничения:**

*   `1 ≤ key.length, value.length ≤ 100`
*   `key` и `value` состоят из строчных английских букв и цифр.
*   `1 ≤ timestamp ≤ 10⁷`
*   Все моменты времени `timestamp` вызовов `set` строго возрастают.
*   Будет сделано не более 2 * 10⁵ вызовов методов `set` и `get`.