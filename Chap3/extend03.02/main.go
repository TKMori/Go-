package main

import (
	"fmt"
)

// 1. 信用分數高於 450 者屬於信用良好, 可享有 15% 的優惠年利率。
// 2. 信用分數低於 450 者, 年利率為 20%。
// 3. 信用良好的人, 每月付款金額最高不得超過每月收入的 20%。
// 4. 信用分數低於450者, 每月付款金額最高不得超過每月收入的 10%。
// 5. 信用分數、每月收入、貸款金額、貸款期數小於 0 時必須示警 (丟出錯誤)。
// 6. 貸款期數如果不能被 12 整除, 也必須示警。
// 7. 利息計算方式為簡單的單利計算：貸款金額乘以年利率再除以 12, 即為每月支付利息。
func calcCredit(no, creditScore, creditCount, income, lornPeriod int) error {
	if creditScore < 0 || income < 0 || creditCount < 0 || lornPeriod < 0 {
		return fmt.Errorf("some thing wrong")
	}

	var annualInterestRate int

	if creditScore > 450 {
		annualInterestRate = 15
	} else {
		annualInterestRate = 20
	}

	totalInterest := creditCount / 100 * annualInterestRate
	monthInterest := totalInterest / 12

	if (creditScore > 450 && monthInterest > income/100*20) || (creditScore < 450 && monthInterest > income/100*10) {
		return fmt.Errorf("monthInterest error")
	}

	if lornPeriod%12 != 0 {
		return fmt.Errorf("lornPeriod error")
	}

	return nil
}

func main() {
	err := calcCredit(1, 350, 10000, 1000, 12)
	if err != nil {
		fmt.Println("False")
		return
	}

	fmt.Println("True")
}
