// https://leetcode.com/problems/time-based-key-value-store/description/
package main

import "fmt"

func main() {
	timeMap := Constructor()
	timeMap.Set("foo", "bar", 1)              // сохраняем ключ "foo" со значением "bar" и меткой 1
	fmt.Printf("%q\n", timeMap.Get("foo", 1)) // "bar"
	fmt.Printf("%q\n", timeMap.Get("foo", 3)) // "bar"
	timeMap.Set("foo", "bar2", 4)             // сохраняем ключ "foo" со значением "bar2" и меткой 4
	fmt.Printf("%q\n", timeMap.Get("foo", 4)) // "bar2"
	fmt.Printf("%q\n", timeMap.Get("foo", 5)) // "bar2"
}

type TimeMap struct {
	// TODO: решить
}

func Constructor() TimeMap {
	return TimeMap{}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
}

func (this *TimeMap) Get(key string, timestamp int) string {
	return ""
}
