# 295. Find Median from Data Stream

**Найти медиану из потока данных** · 🔴 Hard · Heap / Priority Queue · [LeetCode](https://leetcode.com/problems/find-median-from-data-stream/description/)

## Условие

**Медиана** — это значение посередине упорядоченного списка целых чисел. Если размер списка чётный, единственного среднего элемента нет, и медиана равна среднему арифметическому двух средних элементов.

- Например, для `arr = [2,3,4]` медиана равна `3`.
- Например, для `arr = [2,3]` медиана равна `(2 + 3) / 2 = 2.5`.

Реализуйте класс `MedianFinder`:

- `MedianFinder()` — инициализирует объект `MedianFinder`.
- `void addNum(int num)` — добавляет целое число `num` из потока данных в структуру данных.
- `double findMedian()` — возвращает медиану всех элементов на данный момент. Принимаются ответы, отличающиеся от правильного не более чем на `10⁻⁵`.

## Примеры

**Пример 1:**

```text
Input
["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
[[], [1], [2], [], [3], []]
Output
[null, null, null, 1.5, null, 2.0]

Explanation
MedianFinder medianFinder = new MedianFinder();
medianFinder.addNum(1);    // arr = [1]
medianFinder.addNum(2);    // arr = [1, 2]
medianFinder.findMedian(); // возвращает 1.5 (т. е. (1 + 2) / 2)
medianFinder.addNum(3);    // arr = [1, 2, 3]
medianFinder.findMedian(); // возвращает 2.0
```

## Ограничения

- `-10⁵ ≤ num ≤ 10⁵`
- Перед вызовом `findMedian` в структуре данных будет хотя бы один элемент.
- Будет сделано не более `5 * 10⁴` вызовов `addNum` и `findMedian`.

## Дополнительный вопрос

- Если все целые числа из потока лежат в диапазоне `[0, 100]`, как бы вы оптимизировали своё решение?
- Если `99%` всех целых чисел из потока лежат в диапазоне `[0, 100]`, как бы вы оптимизировали своё решение?
