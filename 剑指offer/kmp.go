package main

import (
	"fmt"
)

func main() {
	mainString := "123451237123"

	subString := "123451238123"

	fmt.Println("\n子串在主串中第一次出现的位置", KMP(mainString, subString))
}

/**
KMP
**/
//getNext next数组获取
func getNext(T string) [20]int {
	var i, j int
	var next [20]int
	var nextstring [20]string
	i = 0
	j = -1
	next[0] = -1
	nextstring[i] = fmt.Sprintf("%c %d j", T[i], j)
	for i < len(T)-1 {
		if j == -1 || T[i] == T[j] {
			j++
			i++
			if T[i] == T[j] {
				next[i] = next[j]
			} else {
				next[i] = j
			}
		} else {
			j = next[j]
		}
	}
	return next
}

// KMP 字符串匹配
func KMP(S, T string) int {
	next := getNext(T)
	fmt.Print("[")
	for k := 0; k < len(T); k++ {
		fmt.Printf("%2c ", T[k])
	}
	fmt.Println("]")
	fmt.Printf("%2d", next[:len(T)])
	i := 0
	j := 0
	//同时满足才可以  找除字符串出现的第一个位置
	for i < len(S) && j < len(T) {
		if j == -1 || S[i] == T[j] {
			//当字符匹配时 i j 都加1
			i++
			j++
		} else {
			//子串的 偏移量 从next数组中取  i 不变
			j = next[j]
		}
	}
	//如果 j 大于 或者 等于 T串的长度 说明匹配成功
	if j >= len(T)-1 {
		return i - len(T) + 1
	}
	return 0
}
