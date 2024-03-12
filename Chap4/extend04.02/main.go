package main

import "fmt"

// 1. 建立一段切片, 把一週七天的名稱都放進去, 從週一到週日：
// "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"
// 2. 利用切片的範圍擷取寫法和 append() 更改切片, 讓切片變成從週日開始、至週六結束。
// 3. 將切片顯示至主控台。
func main() {
	week := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	newWeek := week[len(week)-1:]
	week = week[:len(week)-1]
	newWeek = append(newWeek, week...)
	fmt.Println(newWeek)
}
