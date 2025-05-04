# 295. [Найти медиану из потока данных (Find Median from Data Stream)](https://leetcode.com/problems/find-median-from-data-stream/description/)

`Hard`

**Медиана** — это среднее значение в упорядоченном списке целых чисел. Если размер списка четный, среднего значения нет, и медиана — это среднее арифметическое двух средних значений.

*   Например, для `arr = [2,3,4]`, медиана равна `3`.
*   Например, для `arr = [2,3]`, медиана равна `(2 + 3) / 2 = 2.5`.

Реализуйте класс `MedianFinder`:

*   `MedianFinder()` инициализирует объект `MedianFinder`.
*   `void addNum(int num)` добавляет целое число `num` из потока данных в структуру данных.
*   `double findMedian()` возвращает медиану всех элементов на данный момент. Ответы в пределах 10⁻⁵ от фактического ответа будут приняты.

**Пример 1:**
```
Input
["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
[[], [1], [2], [], [3], []]
Output
[null, null, null, 1.5, null, 2.0]

Explanation
MedianFinder medianFinder = new MedianFinder();
medianFinder.addNum(1);    // arr = [1]
medianFinder.addNum(2);    // arr = [1, 2]
medianFinder.findMedian(); // return 1.5 (i.e., (1 + 2) / 2)
medianFinder.addNum(3);    // arr[1, 2, 3]
medianFinder.findMedian(); // return 2.0
```

**Ограничения:**

*   -10⁵ ≤ `num` ≤ 10⁵
*   В структуре данных будет хотя бы один элемент перед вызовом `findMedian`.
*   Будет сделано не более 5 \* 10⁴ вызовов к `addNum` и `findMedian`.

**Дополнительные вопросы:**

*   Если все целые числа из потока находятся в диапазоне [0, 100], как бы вы оптимизировали свое решение?
*   Если 99% всех целых чисел из потока находятся в диапазоне [0, 100], как бы вы оптимизировали свое решение?