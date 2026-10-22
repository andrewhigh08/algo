# 121. [Лучшее время для покупки и продажи акций (Best Time to Buy and Sell Stock)](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/)

сложность: `Easy`

Дан массив `prices`, где `prices[i]` - это цена акции в `i`-й день.

Вы хотите максимизировать свою прибыль, выбрав **один день** для покупки одной акции и выбрав **другой день** в будущем для продажи этой акции.

Верните максимальную прибыль, которую вы можете получить от этой сделки. Если вы не можете получить никакой прибыли, верните 0.

```
Пример 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

Пример 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
```

**Ограничения:**

*   `1 ≤ prices.length ≤ 10⁵`
*   `0 ≤ prices[i] ≤ 10⁴`