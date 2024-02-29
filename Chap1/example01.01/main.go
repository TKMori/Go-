package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// 輸出隨機字串
func main() {
	strs := []string{
		"string1",
		"string2",
		"string3",
		"string4",
		"string5",
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(strs) + 5) // 這邊故意讓 index 超出範圍，讓後面有機會觸發 error

	s, err := selector(strs, index)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(s)
}

func selector(strs []string, index int) (string, error) {
	if index < 0 || index >= len(strs) {
		return "", fmt.Errorf("out of range: %d", index)
	}

	return strs[index], nil
}
