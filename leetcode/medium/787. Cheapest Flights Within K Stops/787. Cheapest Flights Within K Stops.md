# 787. [Самые Дешевые Рейсы с K Пересадками (Cheapest Flights Within K Stops)](https://leetcode.com/problems/cheapest-flights-within-k-stops/description/)

`Medium`

Есть `n` городов, соединенных некоторым количеством рейсов. Вам дан массив `flights`, где `flights[i] = [fromᵢ, toᵢ, priceᵢ]` указывает, что есть рейс из города `fromᵢ` в город `toᵢ` со стоимостью `priceᵢ`.

Вам также даны три целых числа `src`, `dst` и `k`. Верните *самую дешевую цену* от `src` до `dst` с **не более чем** `k` пересадками. Если такого маршрута не существует, верните `-1`.

**Пример 1:**\
![](https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-3drawio.png)
```
Input: n = 4, flights = [[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], src = 0, dst = 3, k = 1
Output: 700
Explanation:
Граф показан выше.
Оптимальный путь с не более чем 1 пересадкой из города 0 в 3 отмечен красным и имеет стоимость 100 + 600 = 700.
Обратите внимание, что путь через города [0,1,2,3] дешевле, но он недопустим, так как использует 2 пересадки.
```

**Пример 2:**\
![](https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-1drawio.png)
```
Input: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 1
Output: 200
Explanation:
Граф показан выше.
Оптимальный путь с не более чем 1 пересадкой из города 0 в 2 отмечен красным и имеет стоимость 100 + 100 = 200.
```

**Пример 3:**\
![](https://assets.leetcode.com/uploads/2022/03/18/cheapest-flights-within-k-stops-2drawio.png)
```
Input: n = 3, flights = [[0,1,100],[1,2,100],[0,2,500]], src = 0, dst = 2, k = 0
Output: 500
Explanation:
Граф показан выше.
Оптимальный путь без пересадок из города 0 в 2 отмечен красным и имеет стоимость 500.
```

**Ограничения:**

*   1 ≤ `n` ≤ 100
*   0 ≤ `flights.length` ≤ (n \* (n - 1) / 2)
*   `flights[i].length == 3`
*   0 ≤ `fromᵢ`, `toᵢ` < n
*   `fromᵢ != toᵢ`
*   1 ≤ `priceᵢ` ≤ 10⁴
*   Между двумя городами не будет кратных рейсов.
*   0 ≤ `src`, `dst`, `k` < n
*   `src != dst`