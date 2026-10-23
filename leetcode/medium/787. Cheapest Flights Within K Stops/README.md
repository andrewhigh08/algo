# 787. Cheapest Flights Within K Stops

**Самые дешёвые рейсы не более чем с K пересадками** · 🟡 Medium · Advanced Graphs · [LeetCode](https://leetcode.com/problems/cheapest-flights-within-k-stops/description/)

## Условие

Есть `n` городов, соединённых некоторым количеством рейсов. Дан массив `flights`, где `flights[i] = [fromᵢ, toᵢ, priceᵢ]` означает, что есть рейс из города `fromᵢ` в город `toᵢ` стоимостью `priceᵢ`.

Также даны три целых числа `src`, `dst` и `k`. Верните ***самую низкую цену*** *перелёта из* `src` *в* `dst` *не более чем с* `k` *пересадками*. Если такого маршрута нет, верните `-1`.

## Примеры

**Пример 1:**

![Пример 1](https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-3drawio.png)

```text
Input: n = 4, flights = [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], src = 0, dst = 3, k = 1
Output: 700
Explanation:
Граф показан выше.
Оптимальный путь из города 0 в 3 не более чем с 1 пересадкой отмечен красным, его стоимость 100 + 600 = 700.
Обратите внимание: путь через города [0,1,2,3] дешевле, но недопустим, потому что в нём 2 пересадки.
```

**Пример 2:**

![Пример 2](https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-1drawio.png)

```text
Input: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 1
Output: 200
Explanation:
Граф показан выше.
Оптимальный путь из города 0 в 2 не более чем с 1 пересадкой отмечен красным, его стоимость 100 + 100 = 200.
```

**Пример 3:**

![Пример 3](https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-2drawio.png)

```text
Input: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 0
Output: 500
Explanation:
Граф показан выше.
Оптимальный путь из города 0 в 2 без пересадок отмечен красным, его стоимость 500.
```

## Ограничения

- `2 ≤ n ≤ 100`
- `0 ≤ flights.length ≤ (n * (n - 1) / 2)`
- `flights[i].length == 3`
- `0 ≤ fromᵢ, toᵢ < n`
- `fromᵢ ≠ toᵢ`
- `1 ≤ priceᵢ ≤ 10⁴`
- Между двумя городами нет кратных рейсов.
- `0 ≤ src, dst, k < n`
- `src ≠ dst`
