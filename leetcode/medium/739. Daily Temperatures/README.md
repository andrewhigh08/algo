# 739. Daily Temperatures

**Дневные температуры** · 🟡 Medium · Stack · [LeetCode](https://leetcode.com/problems/daily-temperatures/description/)

## Условие

Дан массив целых чисел `temperatures`, представляющий дневные температуры. Верните *массив `answer`, в котором `answer[i]` — это количество дней, которое нужно подождать после `i`-го дня до более тёплой температуры*. Если в будущем нет дня с более тёплой температурой, оставьте `answer[i] == 0`.

## Примеры

**Пример 1:**

```text
Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]
```

**Пример 2:**

```text
Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]
```

**Пример 3:**

```text
Input: temperatures = [30,60,90]
Output: [1,1,0]
```

## Ограничения

- `1 ≤ temperatures.length ≤ 10⁵`
- `30 ≤ temperatures[i] ≤ 100`
