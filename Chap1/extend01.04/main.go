package main

import "fmt"

// 闖禍的朋友又來找你了, 這回他又遇上一個程式臭蟲。程式碼應該輸出結果為 true, 但總是輸出 false。你能修正他犯的錯嗎？
func main() {
	count := 0

	if count < 5 {
		count = 10
		count++
	}

	fmt.Println(count == 11)
}
