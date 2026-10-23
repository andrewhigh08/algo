# 332. Reconstruct Itinerary

**Восстановить маршрут** · 🔴 Hard · Advanced Graphs · [LeetCode](https://leetcode.com/problems/reconstruct-itinerary/description/)

## Условие

Дан список авиабилетов `tickets`, где `tickets[i] = [fromᵢ, toᵢ]` — аэропорты отправления и прибытия одного рейса. Восстановите маршрут по порядку и верните его.

Все билеты принадлежат человеку, который вылетает из `"JFK"`, поэтому маршрут должен начинаться с `"JFK"`. Если допустимых маршрутов несколько, верните тот, который имеет наименьший лексикографический порядок при чтении как одна строка.

- Например, маршрут `["JFK", "LGA"]` имеет меньший лексикографический порядок, чем `["JFK", "LGB"]`.

Можно считать, что все билеты образуют хотя бы один допустимый маршрут. Использовать нужно все билеты, каждый — ровно один раз.

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2021/03/14/itinerary1-graph.jpg)

```text
Input: tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
Output: ["JFK","MUC","LHR","SFO","SJC"]
```

**Пример 2:**

![Пример 2](https://assets.leetcode.com/uploads/2021/03/14/itinerary2-graph.jpg)

```text
Input: tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
Output: ["JFK","ATL","JFK","SFO","ATL","SFO"]
Explanation: другой возможный маршрут — ["JFK","SFO","ATL","JFK","ATL","SFO"], но он больше в лексикографическом порядке.
```

## Ограничения

- `1 ≤ tickets.length ≤ 300`
- `tickets[i].length == 2`
- `fromᵢ.length == 3`
- `toᵢ.length == 3`
- `fromᵢ` и `toᵢ` состоят из заглавных английских букв.
- `fromᵢ ≠ toᵢ`
