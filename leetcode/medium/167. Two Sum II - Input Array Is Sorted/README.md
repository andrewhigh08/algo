# 167. [Сумма двух II - Массив отсортирован (Two Sum II - Input Array Is Sorted)](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/)

сложность: `Medium`

Дан 1-индексированный массив целых чисел `numbers`, который уже **отсортирован** в неубывающем порядке. Найдите два числа, сумма которых равна определенному числу `target`. Пусть эти два числа будут `numbers[index₁]` и `numbers[index₂]`, где `1 ≤ index₁ < index₂ ≤ numbers.length`. Верните индексы двух чисел, `index₁` и `index₂`, **увеличенные на единицу**, в виде целочисленного массива `[index₁, index₂]` длины 2.

Тесты сгенерированы таким образом, что существует ровно одно решение. Вы не можете использовать один и тот же элемент дважды.

Ваше решение должно использовать только константное дополнительное пространство.

```
Пример 1:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Объяснение: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].

Пример 2:

Input: numbers = [2,3,4], target = 6
Output: [1,3]
Объяснение: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].

Пример 3:

Input: numbers = [-1,0], target = -1
Output: [1,2]
Объяснение: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].
```

**Ограничения:**

*   `2 ≤ numbers.length ≤ 3 * 10⁴`
*   `-1000 ≤ numbers[i] ≤ 1000`
*   `numbers` отсортирован в неубывающем порядке.
*   `-1000 ≤ target ≤ 1000`
*   Тесты сгенерированы таким образом, что существует ровно одно решение.