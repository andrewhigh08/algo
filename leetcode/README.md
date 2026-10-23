# LeetCode

Условия задач на русском и решения на Go. Основа — список [NeetCode 150](https://neetcode.io/practice?tab=neetcode150): 150 задач, разбитых по темам. Сверх него — ещё 4 задачи.

**Прогресс: 17 / 154** · 🟢 Easy 8 / 32 · 🟡 Medium 9 / 101 · 🔴 Hard 0 / 21

## Как устроена задача

```text
<сложность>/<номер>. <название>/
├── README.md   условие на русском: примеры, ограничения, разбор решения
└── main.go     решение или заготовка: сигнатура с LeetCode, примеры в main() и // TODO: решить
```

Запуск — из папки задачи:

```bash
cd "easy/1. Two Sum"
go run main.go
```

✅ — решено · ⬜ — заготовка · 🔒 — премиум-задача LeetCode; бесплатная версия на LintCode указана в её README.

## NeetCode 150

| Тема | Решено |
|---|:-:|
| [Arrays & Hashing](#arrays--hashing) — массивы и хеширование | 9 / 9 |
| [Two Pointers](#two-pointers) — два указателя | 0 / 5 |
| [Sliding Window](#sliding-window) — скользящее окно | 0 / 6 |
| [Stack](#stack) — стек | 4 / 7 |
| [Binary Search](#binary-search) — бинарный поиск | 0 / 7 |
| [Linked List](#linked-list) — связные списки | 0 / 11 |
| [Trees](#trees) — деревья | 0 / 15 |
| [Tries](#tries) — префиксные деревья | 0 / 3 |
| [Heap / Priority Queue](#heap--priority-queue) — куча и очередь с приоритетом | 0 / 7 |
| [Backtracking](#backtracking) — перебор с возвратом | 0 / 9 |
| [Graphs](#graphs) — графы | 0 / 13 |
| [Advanced Graphs](#advanced-graphs) — продвинутые графы | 0 / 6 |
| [1-D Dynamic Programming](#1-d-dynamic-programming) — одномерное ДП | 0 / 12 |
| [2-D Dynamic Programming](#2-d-dynamic-programming) — двумерное ДП | 0 / 11 |
| [Greedy](#greedy) — жадные алгоритмы | 0 / 8 |
| [Intervals](#intervals) — интервалы | 0 / 6 |
| [Math & Geometry](#math--geometry) — математика и геометрия | 0 / 8 |
| [Bit Manipulation](#bit-manipulation) — битовые операции | 0 / 7 |

### Arrays & Hashing

| # | Задача | По-русски | Сложность | Статус |
|--:|---|---|:-:|:-:|
| 217 | [Contains Duplicate](easy/217.%20Contains%20Duplicate/) | Содержит дубликат | 🟢 Easy | [✅](easy/217.%20Contains%20Duplicate/main.go) |
| 242 | [Valid Anagram](easy/242.%20Valid%20Anagram/) | Корректная анаграмма | 🟢 Easy | [✅](easy/242.%20Valid%20Anagram/main.go) |
| 1 | [Two Sum](easy/1.%20Two%20Sum/) | Сумма двух чисел | 🟢 Easy | [✅](easy/1.%20Two%20Sum/main.go) |
| 49 | [Group Anagrams](medium/49.%20Group%20Anagrams/) | Группировка анаграмм | 🟡 Medium | [✅](medium/49.%20Group%20Anagrams/main.go) |
| 347 | [Top K Frequent Elements](medium/347.%20Top%20K%20Frequent%20Elements/) | Топ K частых элементов | 🟡 Medium | [✅](medium/347.%20Top%20K%20Frequent%20Elements/main.go) |
| 271 | [Encode and Decode Strings](medium/271.%20Encode%20and%20Decode%20Strings/) 🔒 | Кодирование и декодирование строк | 🟡 Medium | [✅](medium/271.%20Encode%20and%20Decode%20Strings/main.go) |
| 238 | [Product of Array Except Self](medium/238.%20Product%20of%20Array%20Except%20Self/) | Произведение элементов массива, кроме текущего | 🟡 Medium | [✅](medium/238.%20Product%20of%20Array%20Except%20Self/main.go) |
| 36 | [Valid Sudoku](medium/36.%20Valid%20Sudoku/) | Корректный судоку | 🟡 Medium | [✅](medium/36.%20Valid%20Sudoku/main.go) |
| 128 | [Longest Consecutive Sequence](medium/128.%20Longest%20Consecutive%20Sequence/) | Самая длинная последовательность подряд идущих чисел | 🟡 Medium | [✅](medium/128.%20Longest%20Consecutive%20Sequence/main.go) |

### Two Pointers

| # | Задача | По-русски | Сложность | Статус |
|--:|---|---|:-:|:-:|
| 125 | [Valid Palindrome](easy/125.%20Valid%20Palindrome/) | Проверка на палиндром | 🟢 Easy | ⬜ |
| 167 | [Two Sum II - Input Array Is Sorted](medium/167.%20Two%20Sum%20II%20-%20Input%20Array%20Is%20Sorted/) | Сумма двух II — массив отсортирован | 🟡 Medium | ⬜ |
| 15 | [3Sum](medium/15.%203Sum/) | Сумма трёх чисел | 🟡 Medium | ⬜ |
| 11 | [Container With Most Water](medium/11.%20Container%20With%20Most%20Water/) | Контейнер с наибольшим количеством воды | 🟡 Medium | ⬜ |
| 42 | [Trapping Rain Water](hard/42.%20Trapping%20Rain%20Water/) | Сбор дождевой воды | 🔴 Hard | ⬜ |

### Sliding Window

| # | Задача | По-русски | Сложность | Статус |
|--:|---|---|:-:|:-:|
| 121 | [Best Time to Buy and Sell Stock](easy/121.%20Best%20Time%20to%20Buy%20and%20Sell%20Stock/) | Лучшее время для покупки и продажи акций | 🟢 Easy | ⬜ |
| 3 | [Longest Substring Without Repeating Characters](medium/3.%20Longest%20Substring%20Without%20Repeating%20Characters/) | Самая длинная подстрока без повторяющихся символов | 🟡 Medium | ⬜ |
| 424 | [Longest Repeating Character Replacement](medium/424.%20Longest%20Repeating%20Character%20Replacement/) | Самая длинная подстрока с повторяющимися символами после замены | 🟡 Medium | ⬜ |
| 567 | [Permutation in String](medium/567.%20Permutation%20in%20String/) | Перестановка в строке | 🟡 Medium | ⬜ |
| 76 | [Minimum Window Substring](hard/76.%20Minimum%20Window%20Substring/) | Минимальная подстрока-окно | 🔴 Hard | ⬜ |
| 239 | [Sliding Window Maximum](hard/239.%20Sliding%20Window%20Maximum/) | Максимум скользящего окна | 🔴 Hard | ⬜ |

### Stack

| # | Задача | По-русски | Сложность | Статус |
|--:|---|---|:-:|:-:|
| 20 | [Valid Parentheses](easy/20.%20Valid%20Parentheses/) | Валидные скобки | 🟢 Easy | [✅](easy/20.%20Valid%20Parentheses/main.go) |
| 155 | [Min Stack](medium/155.%20Min%20Stack/) | Минимальный стек | 🟡 Medium | [✅](medium/155.%20Min%20Stack/main.go) |
| 150 | [Evaluate Reverse Polish Notation](medium/150.%20Evaluate%20Reverse%20Polish%20Notation/) | Вычисление обратной польской нотации | 🟡 Medium | [✅](medium/150.%20Evaluate%20Reverse%20Polish%20Notation/main.go) |
| 22 | [Generate Parentheses](medium/22.%20Generate%20Parentheses/) | Сгенерировать скобки | 🟡 Medium | [✅](medium/22.%20Generate%20Parentheses/main.go) |
| 739 | [Daily Temperatures](medium/739.%20Daily%20Temperatures/) | Дневные температуры | 🟡 Medium | ⬜ |
| 853 | [Car Fleet](medium/853.%20Car%20Fleet/) | Автоколонна | 🟡 Medium | ⬜ |
| 84 | [Largest Rectangle in Histogram](hard/84.%20Largest%20Rectangle%20in%20Histogram/) | Самый большой прямоугольник в гистограмме | 🔴 Hard | ⬜ |

### Binary Search

| # | Задача | По-русски | Сложность | Статус |
|--:|---|---|:-:|:-:|
| 704 | [Binary Search](easy/704.%20Binary%20Search/) | Бинарный поиск | 🟢 Easy | ⬜ |
| 74 | [Search a 2D Matrix](medium/74.%20Search%20a%202D%20Matrix/) | Поиск в 2D-матрице | 🟡 Medium | ⬜ |
| 875 | [Koko Eating Bananas](medium/875.%20Koko%20Eating%20Bananas/) | Коко ест бананы | 🟡 Medium | ⬜ |
| 153 | [Find Minimum in Rotated Sorted Array](medium/153.%20Find%20Minimum%20in%20Rotated%20Sorted%20Array/) | Найти минимум в повёрнутом отсортированном массиве | 🟡 Medium | ⬜ |
| 33 | [Search in Rotated Sorted Array](medium/33.%20Search%20in%20Rotated%20Sorted%20Array/) | Поиск в повёрнутом отсортированном массиве | 🟡 Medium | ⬜ |
| 981 | [Time Based Key-Value Store](medium/981.%20Time%20Based%20Key-Value%20Store/) | Хранилище ключ-значение на основе времени | 🟡 Medium | ⬜ |
| 4 | [Median of Two Sorted Arrays](hard/4.%20Median%20of%20Two%20Sorted%20Arrays/) | Медиана двух отсортированных массивов | 🔴 Hard | ⬜ |

### Linked List

| # | Задача | По-русски | Сложность | Статус |
|--:|---|---|:-:|:-:|
| 206 | [Reverse Linked List](easy/206.%20Reverse%20Linked%20List/) | Развернуть связный список | 🟢 Easy | ⬜ |
| 21 | [Merge Two Sorted Lists](easy/21.%20Merge%20Two%20Sorted%20Lists/) | Объединить два отсортированных списка | 🟢 Easy | ⬜ |
| 143 | [Reorder List](medium/143.%20Reorder%20List/) | Изменить порядок списка | 🟡 Medium | ⬜ |
| 19 | [Remove Nth Node From End of List](medium/19.%20Remove%20Nth%20Node%20From%20End%20of%20List/) | Удалить N-й узел с конца списка | 🟡 Medium | ⬜ |
| 138 | [Copy List with Random Pointer](medium/138.%20Copy%20List%20with%20Random%20Pointer/) | Копирование списка со случайным указателем | 🟡 Medium | ⬜ |
| 2 | [Add Two Numbers](medium/2.%20Add%20Two%20Numbers/) | Сложение двух чисел | 🟡 Medium | ⬜ |
| 141 | [Linked List Cycle](easy/141.%20Linked%20List%20Cycle/) | Цикл в связном списке | 🟢 Easy | ⬜ |
| 287 | [Find the Duplicate Number](medium/287.%20Find%20the%20Duplicate%20Number/) | Найти дублирующееся число | 🟡 Medium | ⬜ |
| 146 | [LRU Cache](medium/146.%20LRU%20Cache/) | LRU-кэш | 🟡 Medium | ⬜ |
| 23 | [Merge k Sorted Lists](hard/23.%20Merge%20k%20Sorted%20Lists/) | Объединить k отсортированных списков | 🔴 Hard | ⬜ |
| 25 | [Reverse Nodes in k-Group](hard/25.%20Reverse%20Nodes%20in%20k-Group/) | Перевернуть узлы в k-группах | 🔴 Hard | ⬜ |

### Trees

| # | Задача | По-русски | Сложность | Статус |
|--:|---|---|:-:|:-:|
| 226 | [Invert Binary Tree](easy/226.%20Invert%20Binary%20Tree/) | Инвертировать бинарное дерево | 🟢 Easy | ⬜ |
| 104 | [Maximum Depth of Binary Tree](easy/104.%20Maximum%20Depth%20of%20Binary%20Tree/) | Максимальная глубина бинарного дерева | 🟢 Easy | ⬜ |
| 543 | [Diameter of Binary Tree](easy/543.%20Diameter%20of%20Binary%20Tree/) | Диаметр бинарного дерева | 🟢 Easy | ⬜ |
| 110 | [Balanced Binary Tree](easy/110.%20Balanced%20Binary%20Tree/) | Сбалансированное бинарное дерево | 🟢 Easy | ⬜ |
| 100 | [Same Tree](easy/100.%20Same%20Tree/) | Одинаковые деревья | 🟢 Easy | ⬜ |
| 572 | [Subtree of Another Tree](easy/572.%20Subtree%20of%20Another%20Tree/) | Поддерево другого дерева | 🟢 Easy | ⬜ |
| 235 | [Lowest Common Ancestor of a Binary Search Tree](medium/235.%20Lowest%20Common%20Ancestor%20of%20a%20Binary%20Search%20Tree/) | Наименьший общий предок бинарного дерева поиска | 🟡 Medium | ⬜ |
| 102 | [Binary Tree Level Order Traversal](medium/102.%20Binary%20Tree%20Level%20Order%20Traversal/) | Обход бинарного дерева по уровням | 🟡 Medium | ⬜ |
| 199 | [Binary Tree Right Side View](medium/199.%20Binary%20Tree%20Right%20Side%20View/) | Вид бинарного дерева справа | 🟡 Medium | ⬜ |
| 1448 | [Count Good Nodes in Binary Tree](medium/1448.%20Count%20Good%20Nodes%20in%20Binary%20Tree/) | Подсчет хороших узлов в бинарном дереве | 🟡 Medium | ⬜ |
| 98 | [Validate Binary Search Tree](medium/98.%20Validate%20Binary%20Search%20Tree/) | Проверка бинарного дерева поиска | 🟡 Medium | ⬜ |
| 230 | [Kth Smallest Element in a BST](medium/230.%20Kth%20Smallest%20Element%20in%20a%20BST/) | K-й наименьший элемент в BST | 🟡 Medium | ⬜ |
| 105 | [Construct Binary Tree from Preorder and Inorder Traversal](medium/105.%20Construct%20Binary%20Tree%20from%20Preorder%20and%20Inorder%20Traversal/) | Построить бинарное дерево из обходов preorder и inorder | 🟡 Medium | ⬜ |
| 124 | [Binary Tree Maximum Path Sum](hard/124.%20Binary%20Tree%20Maximum%20Path%20Sum/) | Максимальная сумма пути в бинарном дереве | 🔴 Hard | ⬜ |
| 297 | [Serialize and Deserialize Binary Tree](hard/297.%20Serialize%20and%20Deserialize%20Binary%20Tree/) | Сериализация и десериализация бинарного дерева | 🔴 Hard | ⬜ |

### Tries

| # | Задача | По-русски | Сложность | Статус |
|--:|---|---|:-:|:-:|
| 208 | [Implement Trie (Prefix Tree)](medium/208.%20Implement%20Trie%20%28Prefix%20Tree%29/) | Реализовать Trie (префиксное дерево) | 🟡 Medium | ⬜ |
| 211 | [Design Add and Search Words Data Structure](medium/211.%20Design%20Add%20and%20Search%20Words%20Data%20Structure/) | Структура данных для добавления и поиска слов | 🟡 Medium | ⬜ |
| 212 | [Word Search II](hard/212.%20Word%20Search%20II/) | Поиск слов II | 🔴 Hard | ⬜ |

### Heap / Priority Queue

| # | Задача | По-русски | Сложность | Статус |
|--:|---|---|:-:|:-:|
| 703 | [Kth Largest Element in a Stream](easy/703.%20Kth%20Largest%20Element%20in%20a%20Stream/) | K-й по величине элемент в потоке | 🟢 Easy | ⬜ |
| 1046 | [Last Stone Weight](easy/1046.%20Last%20Stone%20Weight/) | Вес последнего камня | 🟢 Easy | ⬜ |
| 973 | [K Closest Points to Origin](medium/973.%20K%20Closest%20Points%20to%20Origin/) | K ближайших точек к началу координат | 🟡 Medium | ⬜ |
| 215 | [Kth Largest Element in an Array](medium/215.%20Kth%20Largest%20Element%20in%20an%20Array/) | K-й по величине элемент в массиве | 🟡 Medium | ⬜ |
| 621 | [Task Scheduler](medium/621.%20Task%20Scheduler/) | Планировщик задач | 🟡 Medium | ⬜ |
| 355 | [Design Twitter](medium/355.%20Design%20Twitter/) | Разработка Твиттера | 🟡 Medium | ⬜ |
| 295 | [Find Median from Data Stream](hard/295.%20Find%20Median%20from%20Data%20Stream/) | Найти медиану из потока данных | 🔴 Hard | ⬜ |

### Backtracking

| # | Задача | По-русски | Сложность | Статус |
|--:|---|---|:-:|:-:|
| 78 | [Subsets](medium/78.%20Subsets/) | Подмножества | 🟡 Medium | ⬜ |
| 39 | [Combination Sum](medium/39.%20Combination%20Sum/) | Сумма комбинаций | 🟡 Medium | ⬜ |
| 40 | [Combination Sum II](medium/40.%20Combination%20Sum%20II/) | Сумма комбинаций II | 🟡 Medium | ⬜ |
| 46 | [Permutations](medium/46.%20Permutations/) | Перестановки | 🟡 Medium | ⬜ |
| 90 | [Subsets II](medium/90.%20Subsets%20II/) | Подмножества II | 🟡 Medium | ⬜ |
| 79 | [Word Search](medium/79.%20Word%20Search/) | Поиск слова | 🟡 Medium | ⬜ |
| 131 | [Palindrome Partitioning](medium/131.%20Palindrome%20Partitioning/) | Разбиение на палиндромы | 🟡 Medium | ⬜ |
| 17 | [Letter Combinations of a Phone Number](medium/17.%20Letter%20Combinations%20of%20a%20Phone%20Number/) | Буквенные комбинации телефонного номера | 🟡 Medium | ⬜ |
| 51 | [N-Queens](hard/51.%20N-Queens/) | N ферзей | 🔴 Hard | ⬜ |

### Graphs

| # | Задача | По-русски | Сложность | Статус |
|--:|---|---|:-:|:-:|
| 200 | [Number of Islands](medium/200.%20Number%20of%20Islands/) | Количество островов | 🟡 Medium | ⬜ |
| 695 | [Max Area of Island](medium/695.%20Max%20Area%20of%20Island/) | Максимальная площадь острова | 🟡 Medium | ⬜ |
| 133 | [Clone Graph](medium/133.%20Clone%20Graph/) | Клонировать граф | 🟡 Medium | ⬜ |
| 286 | [Walls and Gates](medium/286.%20Walls%20and%20Gates/) 🔒 | Стены и ворота | 🟡 Medium | ⬜ |
| 994 | [Rotting Oranges](medium/994.%20Rotting%20Oranges/) | Гниющие апельсины | 🟡 Medium | ⬜ |
| 417 | [Pacific Atlantic Water Flow](medium/417.%20Pacific%20Atlantic%20Water%20Flow/) | Поток воды Тихого и Атлантического океанов | 🟡 Medium | ⬜ |
| 130 | [Surrounded Regions](medium/130.%20Surrounded%20Regions/) | Окружённые регионы | 🟡 Medium | ⬜ |
| 207 | [Course Schedule](medium/207.%20Course%20Schedule/) | Расписание курсов | 🟡 Medium | ⬜ |
| 210 | [Course Schedule II](medium/210.%20Course%20Schedule%20II/) | Расписание курсов II | 🟡 Medium | ⬜ |
| 261 | [Graph Valid Tree](medium/261.%20Graph%20Valid%20Tree/) 🔒 | Является ли граф деревом | 🟡 Medium | ⬜ |
| 323 | [Number of Connected Components in an Undirected Graph](medium/323.%20Number%20of%20Connected%20Components%20in%20an%20Undirected%20Graph/) 🔒 | Количество компонент связности в неориентированном графе | 🟡 Medium | ⬜ |
| 684 | [Redundant Connection](medium/684.%20Redundant%20Connection/) | Избыточное соединение | 🟡 Medium | ⬜ |
| 127 | [Word Ladder](hard/127.%20Word%20Ladder/) | Лестница слов | 🔴 Hard | ⬜ |

### Advanced Graphs

| # | Задача | По-русски | Сложность | Статус |
|--:|---|---|:-:|:-:|
| 743 | [Network Delay Time](medium/743.%20Network%20Delay%20Time/) | Время задержки сети | 🟡 Medium | ⬜ |
| 332 | [Reconstruct Itinerary](hard/332.%20Reconstruct%20Itinerary/) | Восстановить маршрут | 🔴 Hard | ⬜ |
| 1584 | [Min Cost to Connect All Points](medium/1584.%20Min%20Cost%20to%20Connect%20All%20Points/) | Минимальная стоимость соединения всех точек | 🟡 Medium | ⬜ |
| 778 | [Swim in Rising Water](hard/778.%20Swim%20in%20Rising%20Water/) | Плавание в поднимающейся воде | 🔴 Hard | ⬜ |
| 269 | [Alien Dictionary](hard/269.%20Alien%20Dictionary/) 🔒 | Инопланетный словарь | 🔴 Hard | ⬜ |
| 787 | [Cheapest Flights Within K Stops](medium/787.%20Cheapest%20Flights%20Within%20K%20Stops/) | Самые дешёвые рейсы не более чем с K пересадками | 🟡 Medium | ⬜ |

### 1-D Dynamic Programming

| # | Задача | По-русски | Сложность | Статус |
|--:|---|---|:-:|:-:|
| 70 | [Climbing Stairs](easy/70.%20Climbing%20Stairs/) | Подъем по лестнице | 🟢 Easy | ⬜ |
| 746 | [Min Cost Climbing Stairs](easy/746.%20Min%20Cost%20Climbing%20Stairs/) | Минимальная стоимость подъема по лестнице | 🟢 Easy | ⬜ |
| 198 | [House Robber](medium/198.%20House%20Robber/) | Ограбление домов | 🟡 Medium | ⬜ |
| 213 | [House Robber II](medium/213.%20House%20Robber%20II/) | Ограбление домов II | 🟡 Medium | ⬜ |
| 5 | [Longest Palindromic Substring](medium/5.%20Longest%20Palindromic%20Substring/) | Самая длинная палиндромная подстрока | 🟡 Medium | ⬜ |
| 647 | [Palindromic Substrings](medium/647.%20Palindromic%20Substrings/) | Палиндромные подстроки | 🟡 Medium | ⬜ |
| 91 | [Decode Ways](medium/91.%20Decode%20Ways/) | Способы декодирования | 🟡 Medium | ⬜ |
| 322 | [Coin Change](medium/322.%20Coin%20Change/) | Размен монет | 🟡 Medium | ⬜ |
| 152 | [Maximum Product Subarray](medium/152.%20Maximum%20Product%20Subarray/) | Максимальное произведение подмассива | 🟡 Medium | ⬜ |
| 139 | [Word Break](medium/139.%20Word%20Break/) | Разбиение на слова | 🟡 Medium | ⬜ |
| 300 | [Longest Increasing Subsequence](medium/300.%20Longest%20Increasing%20Subsequence/) | Самая длинная возрастающая подпоследовательность | 🟡 Medium | ⬜ |
| 416 | [Partition Equal Subset Sum](medium/416.%20Partition%20Equal%20Subset%20Sum/) | Разбиение на равные подмножества по сумме | 🟡 Medium | ⬜ |

### 2-D Dynamic Programming

| # | Задача | По-русски | Сложность | Статус |
|--:|---|---|:-:|:-:|
| 62 | [Unique Paths](medium/62.%20Unique%20Paths/) | Уникальные пути | 🟡 Medium | ⬜ |
| 1143 | [Longest Common Subsequence](medium/1143.%20Longest%20Common%20Subsequence/) | Самая длинная общая подпоследовательность | 🟡 Medium | ⬜ |
| 309 | [Best Time to Buy and Sell Stock with Cooldown](medium/309.%20Best%20Time%20to%20Buy%20and%20Sell%20Stock%20with%20Cooldown/) | Лучшее время для покупки и продажи акций с охлаждением | 🟡 Medium | ⬜ |
| 518 | [Coin Change II](medium/518.%20Coin%20Change%20II/) | Размен монет II | 🟡 Medium | ⬜ |
| 494 | [Target Sum](medium/494.%20Target%20Sum/) | Целевая сумма | 🟡 Medium | ⬜ |
| 97 | [Interleaving String](medium/97.%20Interleaving%20String/) | Чередующаяся строка | 🟡 Medium | ⬜ |
| 329 | [Longest Increasing Path in a Matrix](hard/329.%20Longest%20Increasing%20Path%20in%20a%20Matrix/) | Самый длинный возрастающий путь в матрице | 🔴 Hard | ⬜ |
| 115 | [Distinct Subsequences](hard/115.%20Distinct%20Subsequences/) | Различные подпоследовательности | 🔴 Hard | ⬜ |
| 72 | [Edit Distance](medium/72.%20Edit%20Distance/) | Редакционное расстояние | 🟡 Medium | ⬜ |
| 312 | [Burst Balloons](hard/312.%20Burst%20Balloons/) | Лопнуть шары | 🔴 Hard | ⬜ |
| 10 | [Regular Expression Matching](hard/10.%20Regular%20Expression%20Matching/) | Сопоставление регулярных выражений | 🔴 Hard | ⬜ |

### Greedy

| # | Задача | По-русски | Сложность | Статус |
|--:|---|---|:-:|:-:|
| 53 | [Maximum Subarray](medium/53.%20Maximum%20Subarray/) | Максимальный подмассив | 🟡 Medium | ⬜ |
| 55 | [Jump Game](medium/55.%20Jump%20Game/) | Игра в прыжки | 🟡 Medium | ⬜ |
| 45 | [Jump Game II](medium/45.%20Jump%20Game%20II/) | Игра в прыжки II | 🟡 Medium | ⬜ |
| 134 | [Gas Station](medium/134.%20Gas%20Station/) | Заправка | 🟡 Medium | ⬜ |
| 846 | [Hand of Straights](medium/846.%20Hand%20of%20Straights/) | Разбиение карт на стриты | 🟡 Medium | ⬜ |
| 1899 | [Merge Triplets to Form Target Triplet](medium/1899.%20Merge%20Triplets%20to%20Form%20Target%20Triplet/) | Объединение триплетов для формирования целевого триплета | 🟡 Medium | ⬜ |
| 763 | [Partition Labels](medium/763.%20Partition%20Labels/) | Разделение меток | 🟡 Medium | ⬜ |
| 678 | [Valid Parenthesis String](medium/678.%20Valid%20Parenthesis%20String/) | Валидная строка со скобками | 🟡 Medium | ⬜ |

### Intervals

| # | Задача | По-русски | Сложность | Статус |
|--:|---|---|:-:|:-:|
| 57 | [Insert Interval](medium/57.%20Insert%20Interval/) | Вставка интервала | 🟡 Medium | ⬜ |
| 56 | [Merge Intervals](medium/56.%20Merge%20Intervals/) | Объединение интервалов | 🟡 Medium | ⬜ |
| 435 | [Non-overlapping Intervals](medium/435.%20Non-overlapping%20Intervals/) | Непересекающиеся интервалы | 🟡 Medium | ⬜ |
| 252 | [Meeting Rooms](easy/252.%20Meeting%20Rooms/) 🔒 | Комнаты для переговоров | 🟢 Easy | ⬜ |
| 253 | [Meeting Rooms II](medium/253.%20Meeting%20Rooms%20II/) 🔒 | Комнаты для переговоров II | 🟡 Medium | ⬜ |
| 1851 | [Minimum Interval to Include Each Query](hard/1851.%20Minimum%20Interval%20to%20Include%20Each%20Query/) | Минимальный интервал для включения каждого запроса | 🔴 Hard | ⬜ |

### Math & Geometry

| # | Задача | По-русски | Сложность | Статус |
|--:|---|---|:-:|:-:|
| 48 | [Rotate Image](medium/48.%20Rotate%20Image/) | Повернуть изображение | 🟡 Medium | ⬜ |
| 54 | [Spiral Matrix](medium/54.%20Spiral%20Matrix/) | Спиральная матрица | 🟡 Medium | ⬜ |
| 73 | [Set Matrix Zeroes](medium/73.%20Set%20Matrix%20Zeroes/) | Установить нули в матрице | 🟡 Medium | ⬜ |
| 202 | [Happy Number](easy/202.%20Happy%20Number/) | Счастливое число | 🟢 Easy | ⬜ |
| 66 | [Plus One](easy/66.%20Plus%20One/) | Плюс один | 🟢 Easy | ⬜ |
| 50 | [Pow(x, n)](medium/50.%20Pow%28x%2C%20n%29/) | Возведение в степень | 🟡 Medium | ⬜ |
| 43 | [Multiply Strings](medium/43.%20Multiply%20Strings/) | Умножение строк | 🟡 Medium | ⬜ |
| 2013 | [Detect Squares](medium/2013.%20Detect%20Squares/) | Обнаружение квадратов | 🟡 Medium | ⬜ |

### Bit Manipulation

| # | Задача | По-русски | Сложность | Статус |
|--:|---|---|:-:|:-:|
| 136 | [Single Number](easy/136.%20Single%20Number/) | Одиночное число | 🟢 Easy | ⬜ |
| 191 | [Number of 1 Bits](easy/191.%20Number%20of%201%20Bits/) | Количество единичных битов | 🟢 Easy | ⬜ |
| 338 | [Counting Bits](easy/338.%20Counting%20Bits/) | Подсчет битов | 🟢 Easy | ⬜ |
| 190 | [Reverse Bits](easy/190.%20Reverse%20Bits/) | Перевернуть биты | 🟢 Easy | ⬜ |
| 268 | [Missing Number](easy/268.%20Missing%20Number/) | Пропавшее число | 🟢 Easy | ⬜ |
| 371 | [Sum of Two Integers](medium/371.%20Sum%20of%20Two%20Integers/) | Сумма двух целых чисел | 🟡 Medium | ⬜ |
| 7 | [Reverse Integer](medium/7.%20Reverse%20Integer/) | Перевернуть целое число | 🟡 Medium | ⬜ |

## Дополнительные задачи

Задачи вне списка NeetCode 150.

| # | Задача | По-русски | Сложность | Тема | Статус |
|--:|---|---|:-:|---|:-:|
| 9 | [Palindrome Number](easy/9.%20Palindrome%20Number/) | Число-палиндром | 🟢 Easy | Math & Geometry | [✅](easy/9.%20Palindrome%20Number/main.go) |
| 26 | [Remove Duplicates from Sorted Array](easy/26.%20Remove%20Duplicates%20from%20Sorted%20Array/) | Удаление дубликатов из отсортированного массива | 🟢 Easy | Two Pointers | [✅](easy/26.%20Remove%20Duplicates%20from%20Sorted%20Array/main.go) |
| 27 | [Remove Element](easy/27.%20Remove%20Element/) | Удаление элемента | 🟢 Easy | Two Pointers | [✅](easy/27.%20Remove%20Element/main.go) |
| 977 | [Squares of a Sorted Array](easy/977.%20Squares%20of%20a%20Sorted%20Array/) | Квадраты отсортированного массива | 🟢 Easy | Two Pointers | [✅](easy/977.%20Squares%20of%20a%20Sorted%20Array/main.go) |
