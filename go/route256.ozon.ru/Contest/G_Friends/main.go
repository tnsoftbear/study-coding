package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func addFriend(user1, user2 int, friends *map[int][]int) {
	var exists = false
	for _, friend := range (*friends)[user1] {
		if friend == user2 {
			exists = true
			break
		}
	}
	if !exists {
		(*friends)[user1] = append((*friends)[user1], user2)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var userCount, pairCount int
	fmt.Fscanf(in, "%d %d\n", &userCount, &pairCount)
	var user1, user2 int
	var friends = make(map[int][]int, userCount)
	for i := 1; i <= userCount; i++ {
		friends[i] = []int{}
	}

	for pairIdx := 0; pairIdx < pairCount; pairIdx++ {
		fmt.Fscanf(in, "%d %d\n", &user1, &user2)
		addFriend(user1, user2, &friends)
		addFriend(user2, user1, &friends)
	}

	for user := range friends {
		sort.Ints(friends[user])
	}

	for user := 1; user <= userCount; user++ {
		friendUsers := friends[user]
		if len(friendUsers) == 0 {
			fmt.Fprintln(out, "0")
			continue
		}

		var friendUserScore = make(map[int]int, len(friendUsers))
		// var subFriendUsers = make(map[int][]int, len(friendUsers))
		for _, friendUser := range friendUsers {
			for _, subFriendUser := range friends[friendUser] {
				if subFriendUser == user {
					continue
				}
				friendUserScore[subFriendUser]++
			}
		}

		var maxScore int
		for _, score := range friendUserScore {
			if maxScore == 0 || score > maxScore {
				maxScore = score
			}
		}

		if maxScore == 0 {
			fmt.Fprintln(out, "0")
			continue
		}

		for subFriendUser, score := range friendUserScore {
			if score == maxScore {
				fmt.Fprintf(out, "%d ", subFriendUser)
			}
		}
		fmt.Fprintln(out)

		// fmt.Fprintf(out, "User: %d, score: %v\n", user, friendUserScore)
	}

	// fmt.Fprintf(out, "%v\n", friends)
}
