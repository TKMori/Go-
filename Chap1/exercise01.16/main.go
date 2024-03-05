package main

import (
	"fmt"
	"log"
)

// 設計一個可以 cache book 跟 cd 的 function
// 並且使用 const 限定儲存上限和定義相關名稱

const (
	cacheLimit int = 10

	bookPreWord string = "book_"
	cdPreWord   string = "cd_"
)

var cache map[string]string

func getCache(key string) (string, error) {
	val, ok := cache[key]

	if !ok {
		return "", fmt.Errorf("key: %s is not exist", key)
	}

	return val, nil
}

func setBookCache(isbn, name string) error {
	if len(cache)+1 >= cacheLimit {
		return fmt.Errorf("cache limit")
	}

	cache[bookPreWord+isbn] = name

	return nil
}

func setCDCache(isbn, name string) error {
	if len(cache)+1 >= cacheLimit {
		return fmt.Errorf("cache limit")
	}

	cache[cdPreWord+isbn] = name

	return nil
}

func main() {
	cache = make(map[string]string, cacheLimit)

	if err := setBookCache("123456789", "book1"); err != nil {
		log.Fatal(err)
	}

	if err := setCDCache("987654321", "cd1"); err != nil {
		log.Fatal(err)
	}

	result1, err := getCache(bookPreWord + "123456789")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result1)

	result2, err := getCache(cdPreWord + "987654321")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result2)

	result3, err := getCache(cdPreWord + "123456789")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result3)
}
