package Test

import (
	"fmt"
	"testing"
)
func Test59(t *testing.T){
	nums :=[]int {1,3,-1,-3,5,3,6,7}
	k := 3
	fmt.Println(maxSlidingWindow(nums,k))
}
func maxSlidingWindow(nums []int, k int) []int {
	var result []int
	if len(nums) == 0{
		result[0] = 0
		return result
	}
	result = make([]int,len(nums) - k + 1)
	i := 0

	for i+k-1<len(nums){
		result[i] = max(nums[i:i+k])
		i++
	}
	return result
}
func max(nums []int) int {
	maxNum := nums[0]
	for _,v:=range nums{
		if maxNum < v{
			maxNum = v
		}
	}
	return maxNum
}
