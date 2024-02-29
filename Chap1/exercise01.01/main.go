package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// 印出隨機 1 - 5 個 * 符號
func main() {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(5) + 1
	fmt.Print(strings.Repeat("*", index))
}
