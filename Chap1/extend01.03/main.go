package main

import "fmt"

// 以下的程式碼出了錯, 原作者也束手無策, 因此向你求助。你能讓它運作起來嗎？
func main() {
	count := 5
	var message string

	if count > 5 {
		message = "Greater than 5"
	} else {
		message = "Not greater than 5"
	}

	fmt.Println(message)
}
