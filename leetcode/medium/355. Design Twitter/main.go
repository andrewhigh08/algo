// https://leetcode.com/problems/design-twitter/description/
package main

import "fmt"

func main() {
	twitter := Constructor()
	twitter.PostTweet(1, 5)             // пользователь 1 публикует твит (id = 5)
	fmt.Println(twitter.GetNewsFeed(1)) // [5]
	twitter.Follow(1, 2)                // пользователь 1 подписывается на пользователя 2
	twitter.PostTweet(2, 6)             // пользователь 2 публикует твит (id = 6)
	fmt.Println(twitter.GetNewsFeed(1)) // [6 5]
	twitter.Unfollow(1, 2)              // пользователь 1 отписывается от пользователя 2
	fmt.Println(twitter.GetNewsFeed(1)) // [5]
}

type Twitter struct {
	// TODO: решить
}

func Constructor() Twitter {
	return Twitter{}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	return nil
}

func (this *Twitter) Follow(followerId int, followeeId int) {
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
}
