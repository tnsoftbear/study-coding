package main

import "fmt"

const NO = "NO"
const YES = "YES"

func ll(str string) {
	println(str)
}

func Contains(haystack []int, needle int) int {
	ll(fmt.Sprintf("Searching needle %v in haystack %v", needle, haystack))
	for i, el := range haystack {
		if el == needle {
			ll(fmt.Sprintf("Needle found at index %d", i))
			return i
		}
	}
	ll("Needle not found")
	return -1
}

func Remove(haystack []int, value int) []int {
	// ll("Remove %d -> Haystack is %v", value, haystack)
	var index = Contains(haystack, value)
	haystack = append(haystack[:index], haystack[index+1:]...)
	ll(fmt.Sprintf("Remove %d -> Haystack is %v", value, haystack))
	return haystack
}

func Tickets(peopleInLine []int) string {
	var pocket = []int{}
	for i, banknote := range peopleInLine {
		ll(fmt.Sprintf("Iteration #%d cache %d", i, banknote))
		switch banknote {
		case 50:
			var foundI = Contains(pocket, 25)
			if foundI != -1 {
				pocket = Remove(pocket, 25)
			} else {
				return NO
			}
		case 100:
			var found25I = Contains(pocket, 25)
			var found50I = Contains(pocket, 50)
			if found25I != -1 && found50I != -1 {
				pocket = Remove(pocket, 25)
				pocket = Remove(pocket, 50)
			} else if found50I != -1 && found25I == -1 {
				return NO
			} else if found50I == -1 && found25I != -1 {
				pocket = Remove(pocket, 25)
				found25I = Contains(pocket, 25)
				if found25I == -1 {
					return NO
				}
				pocket = Remove(pocket, 25)
				found25I = Contains(pocket, 25)
				if found25I == -1 {
					return NO
				}
				pocket = Remove(pocket, 25)
			} else {
				return NO
			}
		}
		pocket = append(pocket, banknote)
		ll(fmt.Sprintf("Append -> Haystack is %v", pocket))
	}
	return YES
}

func main() {
	ll(Tickets([]int{25, 25, 50, 100}))
}

/**

The new "Avengers" movie has just been released! There are a lot of people at the cinema box office standing in a huge line. Each of them has a single 100, 50 or 25 dollar bill. An "Avengers" ticket costs 25 dollars.
Vasya is currently working as a clerk. He wants to sell a ticket to every single person in this line.
Can Vasya sell a ticket to every person and give change if he initially has no money and sells the tickets strictly in the order people queue?
Return YES, if Vasya can sell a ticket to every person and give change with the bills he has at hand at that moment. Otherwise return NO.

Examples:
Tickets([]int{25, 25, 50}) // => YES
Tickets([]int{25, 100}) // => NO. Vasya will not have enough money to give change to 100 dollars
Tickets([]int{25, 25, 50, 50, 100}) // => NO. Vasya will not have the right bills to give 75 dollars of change (you can't make two

*/
