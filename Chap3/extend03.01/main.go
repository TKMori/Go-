package main

import "fmt"

// 計算稅金並加總售價。
func calcTotal(price, pTax float64) float64 {
	var tax float64
	tax = price / 100 * pTax

	return tax + price
}

func main() {
	tatal := calcTotal(0.99, 10)

	fmt.Printf("Sales Tax Total: %f", tatal)
}
