package main

import "fmt"

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron",
		"Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	for _, user := range users {
		for _, char := range user {
			switch char {
			case 'e', 'E':
				distribution[user]++
				coins--
			case 'i', 'I':
				distribution[user] += 2
				coins -= 2
			case 'o', 'O':
				distribution[user] += 3
				coins -= 3
			case 'u', 'U':
				distribution[user] += 4
				coins -= 4
			}
		}
	}

	for user, coin := range distribution {
		fmt.Printf("%s 分到 %d 枚金币\n", user, coin)
	}

	return coins
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}
