// https://leetcode.com/problems/reconstruct-itinerary/description/
package main

import "fmt"

func main() {
	fmt.Printf("%q\n", findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}))                 // ["JFK" "MUC" "LHR" "SFO" "SJC"]
	fmt.Printf("%q\n", findItinerary([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}})) // ["JFK" "ATL" "JFK" "SFO" "ATL" "SFO"]
}

func findItinerary(tickets [][]string) []string {
	// TODO: решить
	return nil
}
