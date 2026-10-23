# 139. Word Break

**Разбиение на слова** · 🟡 Medium · 1-D Dynamic Programming · [LeetCode](https://leetcode.com/problems/word-break/description/)

## Условие

Дана строка `s` и словарь строк `wordDict`. Верните `true`, если `s` может быть сегментирована в последовательность из одного или нескольких словарных слов, разделённых пробелами.

**Примечание:** одно и то же слово из словаря может быть использовано при сегментации несколько раз.

## Примеры

**Пример 1:**

```text
Input: s = "leetcode", wordDict = ["leet","code"]
Output: true
Explanation: Верните true, потому что "leetcode" можно сегментировать как "leet code".
```

**Пример 2:**

```text
Input: s = "applepenapple", wordDict = ["apple","pen"]
Output: true
Explanation: Верните true, потому что "applepenapple" можно сегментировать как "apple pen apple".
Обратите внимание, что вам разрешено повторно использовать словарное слово.
```

**Пример 3:**

```text
Input: s = "catsandog", wordDict = ["cats","dog","sand","and","cat"]
Output: false
```

## Ограничения

- `1 ≤ s.length ≤ 300`
- `1 ≤ wordDict.length ≤ 1000`
- `1 ≤ wordDict[i].length ≤ 20`
- `s` и `wordDict[i]` состоят только из строчных английских букв.
- Все строки в `wordDict` **уникальны**.
