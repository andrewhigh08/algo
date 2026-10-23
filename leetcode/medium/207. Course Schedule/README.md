# 207. Course Schedule

**Расписание курсов** · 🟡 Medium · Graphs · [LeetCode](https://leetcode.com/problems/course-schedule/description/)

## Условие

Всего вам нужно пройти `numCourses` курсов, пронумерованных от `0` до `numCourses - 1`. Дан массив `prerequisites`, где `prerequisites[i] = [aᵢ, bᵢ]` означает, что вы **обязательно** должны сначала пройти курс `bᵢ`, если хотите пройти курс `aᵢ`.

- Например, пара `[0, 1]` означает, что для прохождения курса `0` нужно сначала пройти курс `1`.

Верните `true`, если можно пройти все курсы. Иначе верните `false`.

## Примеры

**Пример 1:**

```text
Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: всего нужно пройти 2 курса.
Чтобы пройти курс 1, нужно сначала закончить курс 0. Значит, это возможно.
```

**Пример 2:**

```text
Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: всего нужно пройти 2 курса.
Чтобы пройти курс 1, нужно сначала закончить курс 0, а чтобы пройти курс 0, нужно также закончить курс 1. Значит, это невозможно.
```

## Ограничения

- `1 ≤ numCourses ≤ 2000`
- `0 ≤ prerequisites.length ≤ 5000`
- `prerequisites[i].length == 2`
- `0 ≤ aᵢ, bᵢ < numCourses`
- Все пары `prerequisites[i]` **уникальны**.
