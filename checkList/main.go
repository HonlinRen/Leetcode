package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("success")
}
func checkList(list []int) (val int, ok bool) {
	if len(list) < 3 {
		return 0, false
	}
	// 检查等差，等比数列
	var tmp int
	commonDifference := true
	commonRatio := true
	for i := 1; i < len(list)-1; i++ {
		if list[i]*2 != list[i-1]+list[i+1] {
			commonDifference = false
		}
		if list[i]*list[i] != list[i-1]*list[i+1] {
			commonRatio = false
		}
	}
	if commonDifference {
		tmp = int(list[len(list)-1]*2 - list[len(list)-2])
		return tmp, true
	}
	zeroFlag := false
	for i := 0; i < len(list); i++ {
		if list[i] == 0 {
			zeroFlag = true
		}
	}
	if commonRatio && !zeroFlag {
		tmp = int(list[len(list)-1] * list[len(list)-1] / list[len(list)-2])
		return tmp, true
	}
	//平方根后的等差数列
	otherType := true
	var tmpList []int
	for i := 0; i < len(list); i++ {
		if list[i] < 0 {
			otherType = false
			break
		}
		tmp = int(math.Sqrt(float64(list[i])))
		if tmp*tmp != list[i] {
			otherType = false
			break
		}
		tmpList = append(tmpList, tmp)
	}
	for i := 1; i < len(tmpList)-1; i++ {
		if tmpList[i]*2 != tmpList[i-1]*tmpList[i+1] {
			otherType = false
		}
	}
	if otherType {
		tmp = tmpList[(len(tmpList)-1)*2] - tmpList[len(tmpList)-2]
		return tmp, true
	}
	//递归算法 数列循环嵌套

	return 1, true
}
