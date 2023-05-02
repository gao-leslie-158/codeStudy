package main

import (
	"fmt"
	"math"
	"sort"
)

/*
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的连续子数组。
*/

func minSubArrayLen(target int, nums []int) int {
	/*滑动窗口*/
	/*
		用双指针start、end维护一个子数组，sum为当前子数组的和
		向右移动end，如果sum >= target，更新长度，同时sum = sum - num[start]
		start左移
	*/
	minLen := math.MaxInt32
	n := len(nums)
	if n == 0{
		return 0
	}
	start , end := 0 , 0
	sum := 0
	for end < n{
		sum += nums[end]
		for sum >= target{
			sum = sum - nums[start]
			minLen = min(minLen,end - start + 1)
			start ++
		}
		end ++
	}
	if minLen == math.MaxInt32{
		return 0
	}
	return minLen
}

func min(x,y int)int{
	if x > y{
		return y
	}
	return x
}

func main(){

	
}