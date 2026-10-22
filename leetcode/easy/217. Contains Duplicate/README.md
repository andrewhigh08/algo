## 217. [Содержит дубликат(Contains Duplicate)](https://leetcode.com/problems/contains-duplicate/)

**Сложность:** Простая

---

### Описание задачи:

Дан целочисленный массив `nums`. Верните `true`, если какое-либо значение встречается в массиве **хотя бы два раза**, и `false`, если все элементы массива уникальны.

---

```
Пример 1:
Input: nums = [1,2,3,1]
Output: true

Объяснение:
The element 1 occurs at the indices 0 and 3.

Пример 2:
Input: nums = [1,2,3,4]
Output: false

Explanation:
All elements are distinct.

Пример 3:
Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true
```

---

### Ограничения:

*   `1 ≤ nums.length ≤ 10⁵` (Длина массива от 1 до 100,000)
*   `-10⁹ ≤ nums[i] ≤ 10⁹` (Значение каждого элемента от -1,000,000,000 до 1,000,000,000)

```mermaid

flowchart TD
    Start([Начало: containsDuplicate]) --> Init["Инициализация: seen = map[int]bool"]
    Init --> LoopStart["Цикл по мапе nums"]
    LoopStart --> CheckExists{"Проверка: 
    seen[num] существует?"}
    CheckExists -->|Да| ReturnTrue["Вернуть true
    (найден дубликат)"]
    CheckExists -->|Нет| MarkSeen["seen[num] = true
    (отметить как увиденное)"]
    MarkSeen --> LoopContinue{"Конец 
    цикла?"}
    LoopContinue -->|Нет| LoopStart
    LoopContinue -->|Да| ReturnFalse["Вернуть false
    (дубликатов нет)"]
    ReturnTrue --> End([Конец])
    ReturnFalse --> End
```