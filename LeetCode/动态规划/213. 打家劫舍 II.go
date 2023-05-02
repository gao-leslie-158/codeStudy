package main
import (
	"fmt"
)

func rob(nums []int) int{
	n := len(nums)
	if n == 0{
		return 0
	}
	if n == 1{
		return nums[0]
	}
	res1 := robRange(nums,0,n-2)	// 含首不含尾
	res2 := robRange(nums,1,n-1)	// 含尾不含首
	return max(res1,res2)
}

func robRange(nums []int,start,end int) int {
	if end == start{
		return nums[start]
	}
	dp := make([]int, len(nums))
	dp[start] , dp[start+1] = nums[start] , max(nums[start],nums[start+1])
	for i := start + 2; i <= end; i ++{
		dp[i] = max(dp[i-1],dp[i-2]+nums[i])
	}
	return dp[end]
}

func max(x , y int) int{
	if x > y{
		return x
	}
	return y
}

func main(){

} 