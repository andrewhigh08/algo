# 704. Binary Search

**Бинарный поиск** · 🟢 Easy · Binary Search · [LeetCode](https://leetcode.com/problems/binary-search/description/)

## Условие

Дан массив целых чисел `nums`, отсортированный по возрастанию, и целое число `target`. Напишите функцию, которая ищет `target` в `nums`. Если `target` есть в массиве, верните его индекс. В противном случае верните `-1`.

Алгоритм должен работать за время `O(log n)`.

## Примеры

**Пример 1:**

```text
Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 есть в nums, его индекс равен 4.
```

**Пример 2:**

```text
Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 нет в nums, поэтому возвращаем -1.
```

## Ограничения

- `1 ≤ nums.length ≤ 10⁴`
- `-10⁴ < nums[i], target < 10⁴`
- Все целые числа в `nums` **уникальны**.
- `nums` отсортирован по возрастанию.
