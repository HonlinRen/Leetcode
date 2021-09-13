package Test

import (
	"fmt"
	"testing"
)

func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

func Test58(t *testing.T) {
	s := "abcdefg"
	k := 2
	fmt.Println(reverseLeftWords(s,k))

}



