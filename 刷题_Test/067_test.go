package 刷题_Test

import (
	"fmt"
	"math"
	"testing"
)


func Test067(t *testing.T) {
	str := "-2147483641"

	fmt.Println(myAtoi(str))
}

func myAtoi(s string) int {
	var (
		negative  bool // 负数
		i         int32
		signed    bool // 是否出现过'-'，'+'
		startZero bool // 是否开始是0,再出现'-'，'+'，return 0
	)
	for k, v := range s { // 去除前空格
		if v != ' ' {
			s = s[k:]
			break
		}
	}
	for _, v := range s {
		if v == '0' && !signed && i == 0 { // 0000-898
			startZero = true
			continue
		}
		if startZero && (v == '+' || v == '-') {
			break
		}
		if v == '+' && !signed { // +
			signed = true
			continue
		}
		if v == '-' && !signed && i == 0 { // -
			negative = true
			signed = true
			continue
		}
		if v < '0' || v > '9' {
			break
		}
		if negative{
			if -i < (math.MinInt32 + (v - '0'))/10{
				i = math.MinInt32
				break
			}
		}else {
			if i > (math.MaxInt32 - (v - '0'))/10{
				i = math.MaxInt32
				break
			}
		}
		i = v - '0' + i*10
	}
	if negative {
		i = -i
	}
	return int(i)
}

