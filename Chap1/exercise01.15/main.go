package main

import "fmt"

func add5Value(count int) {
	count += 5

	fmt.Println(count)
}

func add5Point(count *int) {
	*count += 5

	fmt.Println(*count)
}

func main() {
	count := 0

	add5Value(count)
	fmt.Println(count)

	add5Point(&count)
	fmt.Println(count)
}
