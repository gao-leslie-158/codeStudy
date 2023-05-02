package main

import (
	"fmt"
	"math"
)

func findTargetSumWays(nums []int, target int) int {
	n := len(nums)
	sum := 0
	for _ , v := range nums{
		sum += v
	}
	if (target + sum) % 2 == 1{
		// sum是5，target是2，无解,抵消不了
		// -2+3 = 1，+2-3 = -1 
		return 0
	}
	if int(math.Abs(float64(target))) > sum{
		// nums[i]全+或全-也还是比target小
		return 0
	}
	max_cap := (sum + target) / 2
	dp := make([]int,max_cap+1)
	dp[0] = 1
	for i := 0; i < n; i++{
		for j := max_cap; j >= nums[i]; j--{
			dp[j] += dp[j - nums[i]]
		} 
	}
	return dp[max_cap]
}

func main(){
	nums := []int{1,1,1,1,1}
	target := 4
	res := findTargetSumWays(nums,target)
	fmt.Println(res)
}