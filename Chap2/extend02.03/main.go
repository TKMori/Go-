package main

import (
	"fmt"
)

// 實現氣泡排序法
func main() {
	nums := []int{9, 2, 4, 1, 7, 3, 5, 6, 8}
	isChanged := true

	// 最外層 for 是最大跑的次數
	for i := len(nums) - 1; i > 0; i-- {

		// 如果上一 round 沒有任何交換順序行為，表示已經排序完成
		if !isChanged {
			break
		}

		isChanged = false
		// 實際比較每一 round 的順序
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				isChanged = true
			}
		}
	}

	fmt.Println(nums)
}
