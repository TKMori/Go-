package main

import (
	"fmt"
)

// 1. 寫一段可以從 1 顯示到 100 的程式。
// 2. 如果顯示的數字是 3 的倍數 , 改為顯示文字『 Fizz』。
// 3. 如果顯示的數字是 5 的倍數 , 改為顯示『 Buzz』。
// 4. 如果顯示的數字是 3 和 5 的公倍數 , 改為顯示『 FizzBuzz』。
func main() {
	for i := 1; i < 101; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
