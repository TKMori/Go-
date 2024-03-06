package main

import "fmt"

// 假設你有一個像下面這樣的資料表。你想要找出其中出現次數最多的字, 然後把這個字和次數印出來。
func main() {
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Never": 2,
		"Give":  1,
		"Up":    4,
	}

	maxKey := ""
	maxValue := 0

	for k, v := range words {
		if v > maxValue {
			maxKey = k
			maxValue = v
		}
	}

	fmt.Printf("出現最多次的字: %s\n", maxKey)
	fmt.Printf("出現次數: %d\n", maxValue)
}
