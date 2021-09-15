package 刷题_Test

import (
	"math"
	"testing"
)

func name017 (t testing.T)  {

}
func printNumbers(n int) []int {
	var out []int
	 for i:=1;i<int(math.Pow10(n));i++{
	 	out = append(out,i)
	}
	return out
}


