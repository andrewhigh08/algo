# 355. Design Twitter

**Разработка Твиттера** · 🟡 Medium · Heap / Priority Queue · [LeetCode](https://leetcode.com/problems/design-twitter/description/)

## Условие

Спроектируйте упрощённую версию Твиттера, в которой пользователи могут публиковать твиты, подписываться на других пользователей и отписываться от них, а также видеть `10` самых свежих твитов в своей ленте новостей.

Реализуйте класс `Twitter`:

- `Twitter()` — инициализирует объект Твиттера.
- `void postTweet(int userId, int tweetId)` — публикует новый твит с ID `tweetId` от имени пользователя `userId`. Каждый вызов этой функции выполняется с уникальным `tweetId`.
- `List<Integer> getNewsFeed(int userId)` — возвращает `10` самых свежих ID твитов из ленты новостей пользователя. Каждый твит в ленте должен быть опубликован либо пользователем, на которого подписан данный пользователь, либо им самим. Твиты должны быть **упорядочены от самых новых к самым старым**.
- `void follow(int followerId, int followeeId)` — пользователь с ID `followerId` подписывается на пользователя с ID `followeeId`.
- `void unfollow(int followerId, int followeeId)` — пользователь с ID `followerId` отписывается от пользователя с ID `followeeId`.

## Примеры

**Пример 1:**

```text
Input
["Twitter", "postTweet", "getNewsFeed", "follow", "postTweet", "getNewsFeed", "unfollow", "getNewsFeed"]
[[], [1, 5], [1], [1, 2], [2, 6], [1], [1, 2], [1]]
Output
[null, null, [5], null, null, [6, 5], null, [5]]

Explanation
Twitter twitter = new Twitter();
twitter.postTweet(1, 5); // пользователь 1 публикует новый твит (id = 5)
twitter.getNewsFeed(1);  // лента пользователя 1 должна вернуть список из 1 ID твита -> [5]. Возвращает [5]
twitter.follow(1, 2);    // пользователь 1 подписывается на пользователя 2
twitter.postTweet(2, 6); // пользователь 2 публикует новый твит (id = 6)
twitter.getNewsFeed(1);  // лента пользователя 1 должна вернуть список из 2 ID твитов -> [6, 5]. Твит 6 должен идти раньше твита 5, так как он опубликован позже
twitter.unfollow(1, 2);  // пользователь 1 отписывается от пользователя 2
twitter.getNewsFeed(1);  // лента пользователя 1 должна вернуть список из 1 ID твита -> [5], так как пользователь 1 больше не подписан на пользователя 2
```

## Ограничения

- `1 ≤ userId, followerId, followeeId ≤ 500`
- `0 ≤ tweetId ≤ 10⁴`
- Все твиты имеют **уникальные** ID.
- Будет сделано не более `3 * 10⁴` вызовов `postTweet`, `getNewsFeed`, `follow` и `unfollow`.
- Пользователь не может подписаться сам на себя.
