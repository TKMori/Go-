package main

import "fmt"

func main() {
	a, b := 5, 10

	// 在此呼叫交換函式
	swap(&a, &b)

	fmt.Println(a == 10, b == 5)
}
func swap(a *int, b *int) {
	// 在此交換變數值
	tmp := *b

	*b = *a
	*a = tmp
}
