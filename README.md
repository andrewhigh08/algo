# algo

Решения алгоритмических задач на Go: LeetCode (NeetCode 150), Codeforces, Codewars и Яндекс.

| Платформа | Что внутри | Решено |
|---|---|:-:|
| [LeetCode](leetcode/) | NeetCode 150 по темам и ещё 4 задачи; условия на русском | 17 / 154 |
| [Codeforces](codeforces/) | задачи A и B из Div. 2, рейтинг 800–1500 | 66 / 66 |
| [Codewars](codewars/) | ката 4 kyu | 1 / 1 |
| [Яндекс](yandex/) | «Тренировки по алгоритмам 5.0» | 1 / 1 |

## Структура

```text
algo/
├── leetcode/
│   ├── README.md              оглавление по темам NeetCode 150 и прогресс
│   └── easy/ medium/ hard/
│       └── <номер>. <название>/
│           ├── README.md      условие на русском и разбор решения
│           └── main.go        решение или заготовка (// TODO: решить)
├── codeforces/
│   ├── README.md              таблица задач с рейтингом и тегами
│   └── <контест><буква>. <название>/main.go
├── codewars/
│   └── <ранг>/<название>/     README.md + main.go
└── yandex/
    └── <задача>/              README.md + main.go + input.txt / output.txt
```

## Запуск

Каждая задача — отдельная программа в `package main`. Файла `go.mod` в репозитории нет: в именах папок есть пробелы, а такие пути не годятся для импорта. Поэтому `go run` получает файл явно:

```bash
cd "leetcode/easy/1. Two Sum"
go run main.go
```

Решения Codeforces читают входные данные из stdin, а задача Яндекса — из файла `input.txt` рядом с `main.go`, поэтому её запускают из корня репозитория:

```bash
echo "6 6 4" | go run "codeforces/1A. Theatre Square/main.go"   # 4
go run yandex/A1/main.go                                         # ответ попадёт в yandex/A1/output.txt
```

Проверить форматирование всего кода: `gofmt -l .`
