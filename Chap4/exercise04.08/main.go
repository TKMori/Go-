package main

import (
	"fmt"
	"os"
)

// 使用 os.Args 套件來讀取輸入參數
// 找出輸入的參數中最長的字串
// 輸入參數至少要 3 個

func getPassWord(minLenth int) []string {
	if len(os.Args) < minLenth {
		fmt.Printf("至少輸入 %v 個參數\n", minLenth)
		os.Exit(1)
	}

	var args []string

	// os.Args 讀到的第一個參數是程式名稱不是參數，所以從第二個開始讀取
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}

	return args
}

func findLongestWord(args []string) string {
	longest := ""

	for i := 0; i < len(args); i++ {
		if len(args[i]) > len(longest) {
			longest = args[i]
		}
	}

	return longest
}

func main() {
	longest := findLongestWord(getPassWord(3))
	if len(longest) == 0 {
		fmt.Printf("發生錯誤\n")
		os.Exit(1)
	}

	fmt.Printf("最長的參數是: %v", longest)
}
