package main

import "fmt"

// 1. 定義一個有 10 個 int 元素的陣列。
// 2. 以 for i 迴圈將數字 10 到 19 填入陣列。
// 3. 用 fmt.Println() 將陣列內容顯示在主控台。
func main() {
	var arr [10]int

	num := 10

	for i := 0; i < len(arr); i++ {
		arr[i] = num
		num++
	}

	fmt.Println(arr)
}
