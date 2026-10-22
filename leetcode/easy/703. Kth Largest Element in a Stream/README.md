# 703. [K-й по величине элемент в потоке (Kth Largest Element in a Stream)](https://leetcode.com/problems/kth-largest-element-in-a-stream/description/)

`Easy`

Вы работаете в приемной комиссии университета и вам нужно отслеживать `k`-й по величине результат теста от абитуриентов в режиме реального времени. Это помогает определить пороговые баллы для собеседований и приема, поскольку новые абитуриенты динамически подают свои баллы.

Вам поручено реализовать класс, который для заданного целого числа `k` поддерживает поток результатов тестов и непрерывно возвращает `k`-й по величине результат теста **после** добавления нового балла. Более конкретно, мы ищем `k`-й по величине балл в отсортированном списке всех баллов.

Реализуйте класс `KthLargest`:

*   `KthLargest(int k, int[] nums)` Инициализирует объект целым числом `k` и потоком результатов тестов `nums`.
*   `int add(int val)` Добавляет новый результат теста `val` в поток и возвращает элемент, представляющий `k`-й по величине элемент в пуле результатов тестов на данный момент.

**Пример 1:**
```
Input:
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
Output:
[null, 4, 5, 5, 8, 8]

Explanation:
KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3);   // return 4
kthLargest.add(5);   // return 5
kthLargest.add(10);  // return 5
kthLargest.add(9);   // return 8
kthLargest.add(4);   // return 8
```

**Пример 2:**
```
Input:
["KthLargest", "add", "add", "add", "add"]
[[1, []], [-3], [-2], [-4], [0], [4]]
Output:
[null, -3, -2, -2, 0, 4]

Explanation:
KthLargest kthLargest = new KthLargest(1, []);
kthLargest.add(-3);  // return -3
kthLargest.add(-2);  // return -2
kthLargest.add(-4);  // return -2
kthLargest.add(0);   // return 0
kthLargest.add(4);   // return 4
```

**Ограничения:**

*   0 ≤ `nums.length` ≤ 10⁴
*   1 ≤ `k` ≤ `nums.length` + 1
*   -10⁴ ≤ `nums[i]` ≤ 10⁴
*   -10⁴ ≤ `val` ≤ 10⁴
*   Будет сделано не более 10⁴ вызовов к `add`.