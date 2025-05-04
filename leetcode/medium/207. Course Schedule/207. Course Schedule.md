# 207. [Расписание Курсов (Course Schedule)](https://leetcode.com/problems/course-schedule/description/)

`Medium`

Всего вам нужно пройти `numCourses` курсов, помеченных от `0` до `numCourses - 1`. Вам дан массив `prerequisites`, где `prerequisites[i] = [ai, bi]` указывает, что вы **должны** сначала пройти курс `bi`, если хотите пройти курс `ai`.

*   Например, пара `[0, 1]` указывает, что для прохождения курса `0` вы должны сначала пройти курс `1`.

Верните `true`, если вы можете закончить все курсы. В противном случае верните `false`.

**Пример 1:**
```
Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: Всего нужно пройти 2 курса.
Чтобы пройти курс 1, вы должны закончить курс 0. Так что это возможно.
```

**Пример 2:**
```
Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: Всего нужно пройти 2 курса.
Чтобы пройти курс 1, вы должны закончить курс 0, а чтобы пройти курс 0, вы также должны закончить курс 1. Так что это невозможно.
```

**Ограничения:**

*   1 ≤ `numCourses` ≤ 2000
*   0 ≤ `prerequisites.length` ≤ 5000
*   `prerequisites[i].length == 2`
*   0 ≤ `ai`, `bi` < `numCourses`
*   Все пары `prerequisites[i]` уникальны.