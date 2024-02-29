package main

import (
	"fmt"
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
	index := rand.Intn(len(strs))
	fmt.Println(strs[index])
}
