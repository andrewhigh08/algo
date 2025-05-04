# 739. [Дневные температуры (Daily Temperatures)](https://leetcode.com/problems/daily-temperatures/description/)

сложность: `Medium`

Дан массив целых чисел `temperatures`, представляющий дневные температуры. Верните массив `answer` такой, что `answer[i]` - это количество дней, которое нужно подождать после `i`-го дня, чтобы получить более теплую температуру. Если в будущем нет дня с более теплой температурой, оставьте `answer[i] == 0`.

```
Пример 1:

Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]

Пример 2:

Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]

Пример 3:

Input: temperatures = [30,60,90]
Output: [1,1,0]

```

**Ограничения:**

*   `1 ≤ temperatures.length ≤ 10⁵`
*   `30 ≤ temperatures[i] ≤ 100`