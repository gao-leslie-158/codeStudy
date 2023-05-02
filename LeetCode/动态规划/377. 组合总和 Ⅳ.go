package main
import (
	"fmt"
)
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	n := len(nums)
	dp[0] = 1
	for j := 0; j <= target; j++{
		for i := 0; i < n; i ++{
			if j >= nums[i]{
				dp[j] += dp[j - nums[i]]
			}
		}
	}
	return dp[target]
}

func main(){

}