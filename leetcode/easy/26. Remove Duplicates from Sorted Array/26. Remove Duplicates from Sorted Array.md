# 26. [Удаление дубликатов из отсортированного массива (Remove Duplicates from Sorted Array)](https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/)

**Сложность:** Легкая

## Условие задачи

Дан целочисленный массив `nums`, отсортированный в **неубывающем порядке**. Требуется удалить дубликаты **на месте** так, чтобы каждый уникальный элемент появлялся только **один раз**. **Относительный порядок** элементов должен быть сохранен. Затем вернуть количество уникальных элементов в `nums`.

Рассмотрим количество уникальных элементов в `nums` как `k`. Чтобы решение было принято, необходимо сделать следующее:

* Изменить массив `nums` так, чтобы первые `k` элементов содержали уникальные элементы в том порядке, в котором они изначально присутствовали в `nums`. Оставшиеся элементы не важны, как и размер `nums`.
* Вернуть `k`.

## Пользовательский судья:

Судья будет тестировать ваше решение следующим кодом:

```
int[] nums = [...]; // Входной массив
int[] expectedNums = [...]; // Ожидаемый ответ с правильной длиной

int k = removeDuplicates(nums); // Вызов вашей реализации

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
```

Если все проверки пройдены, то ваше решение будет **принято**.

```
Пример 1:

Input: nums = [1,1,2]
Output: 2, nums = [1,2,_]
Explanation: Your function should return k = 2, with the first two elements of nums being 1 and 2 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

Пример 2:

Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]
Explanation: Your function should return k = 5, with the first five elements of nums being 0, 1, 2, 3, and 4 respectively.
It does not matter what you leave beyond the returned k (hence they are underscores).

```
## Ограничения:

* 1 ≤ nums.length ≤ 3 * 10⁴
* -100 ≤ nums[i] ≤ 100
* Массив `nums` отсортирован в неубывающем порядке.

```mermaid
flowchart TD
    Start([Начало]) --> Check{"Массив пуст?"}
    Check -->|Да| ReturnZero[Вернуть 0]
    Check -->|Нет| InitPointers[Инициализация l = 1]
    InitPointers --> Loop[Проходим по массиву с индекса r = 1]
    Loop --> Compare{"nums[r] не равно nums[r-1]?"}
    Compare -->|Да| Assign["nums[l] = nums[r], l++"]
    Compare -->|Нет| Skip[Пропускаем дублирующийся элемент]
    Assign --> NextIter[r++]
    Skip --> NextIter
    NextIter --> CheckEnd{"r < len(nums)?"}
    CheckEnd -->|Да| Loop
    CheckEnd -->|Нет| ReturnL[Вернуть l]
    ReturnZero --> End([Конец])
    ReturnL --> End
```