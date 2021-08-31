package main

import (
	"fmt"
	"strings"
)
//空格替换位 %20
func main(){
	str := "a   b c  "
	arr :=strings.Split(str," ")
	//arr :=strings.Fields(str)
	var strcopy string

	for k,v := range arr{
		if k < len(arr) - 1{
			v = v + "%20"
			strcopy = strcopy + v
		}

	}
	fmt.Println(strcopy)
}