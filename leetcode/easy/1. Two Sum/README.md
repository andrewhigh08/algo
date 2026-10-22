## 1. [Сумма двух чисел (Two Sum)](https://leetcode.com/problems/two-sum/description/)

**Сложность:** Простая

---

### Описание задачи:

Дан массив целых чисел `nums` и целое число `target`. Верните **индексы** двух чисел из массива `nums` таких, что их сумма равна `target`.

Вы можете предполагать, что каждый входной набор данных имеет **ровно одно решение**, и вы **не можете** использовать один и тот же элемент дважды.

Вы можете вернуть ответ в любом порядке.

---

```
Пример 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Пример 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]

Пример 3:

Input: nums = [3,3], target = 6
Output: [0,1]
```

---

### Ограничения:

*   `2 ≤ nums.length ≤ 10⁴` (Длина массива от 2 до 10,000)
*   `-10⁹ ≤ nums[i] ≤ 10⁹` (Значение каждого элемента от -1,000,000,000 до 1,000,000,000)
*   `-10⁹ ≤ target ≤ 10⁹` (Значение `target` от -1,000,000,000 до 1,000,000,000)
*   **Существует только одно верное решение.**

---

### Дополнительный вопрос (Follow-up):

Можете ли вы придумать алгоритм, временная сложность которого меньше, чем O(n²)?

```mermaid
flowchart TD
    Start([Начало: twoSum]) --> Init["Инициализация: numMap = map[int]int"]
    Init --> LoopStart["Цикл по мапе nums"]
    LoopStart --> CheckMap{"numMap[num] \nсуществует?"}
    CheckMap -->|Да| Return["Вернуть [numMap[num], i]"]
    CheckMap -->|Нет| StoreComplement["numMap[target-num] = i"]
    StoreComplement --> Continue{"Конец \nцикла?"}
    Continue -->|Нет| LoopStart
    Continue -->|Да| ReturnEmpty["Вернуть пустой массив []"]
    Return --> End([Конец])
    ReturnEmpty --> End
```