package main

import "fmt"

// 1. 建立一個切片, 元素順序如下：
// "Good", "Good", "Bad", "Good", "Good"
// 2. 利用切片的範圍擷取寫法和 append() 更改切片 (提示：把切片切成兩半, 重新接起來時漏掉 Bad 這個元素)。
// 3. 將新切片顯示至主控台。
func main() {
	s := []string{"Good", "Good", "Bad", "Good", "Good"}
	s1 := s[:2]
	s2 := s[3:]
	s1 = append(s1, s2...)

	fmt.Println(s1)
}
