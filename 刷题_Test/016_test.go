package 刷题_Test

import (
	"fmt"
	"testing"
)

func Test016(t *testing.T) {
	fmt.Println(myPow(2, -16))
}
func myPow(x float64, n int) float64 {
	// 递归
	if n == 1 {
		return x
	}
	if n == -1 {
		return 1/ x
	}
	if n==0 {
		return 1
	}
	// 如果n为奇数，则n-1退为偶数
	if n & 1 != 0 {
		return x * myPow(x, n-1)
	}
	// 底数平方，指数除2
	return myPow(x*x, n >> 1)

}
