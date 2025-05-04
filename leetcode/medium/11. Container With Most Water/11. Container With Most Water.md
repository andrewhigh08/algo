# 11. [Контейнер с наибольшим количеством воды (Container With Most Water)](https://leetcode.com/problems/container-with-most-water/description/)

сложность: `Medium`

Дан массив целых чисел `height` длины `n`. Нарисованы `n` вертикальных линий так, что две конечные точки `i`-й линии находятся в координатах `(i, 0)` и `(i, height[i])`. Найдите две линии, которые вместе с осью X образуют контейнер, так чтобы этот контейнер мог вместить наибольшее количество воды. Верните максимально возможное количество воды, которое может вместить контейнер.

Обратите внимание, что контейнер нельзя наклонять.

**Пример 1:**

![Визуализация контейнера для Примера 1](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/17/question_11.jpg)

```
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Объяснение: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.

Пример 2:

Input: height = [1,1]
Output: 1
```

**Ограничения:**

*   `n == height.length`
*   `2 ≤ n ≤ 10⁵`
*   `0 ≤ height[i] ≤ 10⁴`