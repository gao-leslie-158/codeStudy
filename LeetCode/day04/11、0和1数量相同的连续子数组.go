package main

import (
	"fmt"
	"math"
)

/*
输入: nums = [0,1,0]
输出: 2
说明: [0, 1] (或 [1, 0]) 是具有相同数量 0 和 1 的最长连续子数组。
*/
func findMaxLength(nums []int) int {
	/*还是采用前缀和+哈希表*/
	/*
		可以把0看成-1，也就是变为了求和0的最长子数组
		1、准备一个前缀和哈希表：presum_map
			key：前缀和值 value：第一次出现下标
		2、count记录当前前缀和
		-1 1 -1 -1 1 1 1
			if count存在，代表key对应的value值下标j到当前下标i之间的和为0
				找到一个子数组
			if count不存在，加入
	*/
	count := 0
	presum_map := map[int]int{0:-1}
	maxLen := 0
	for i := 0; i < len(nums); i ++{
		if nums[i] == 0{
			count --
		}else {
			count ++
		}
		if start , ok := presum_map[count]; ok{
			maxLen = max(maxLen,i - start)
		}else {
			presum_map[count] = i
		}
	}
	return maxLen
}



func max(x , y int) int{
	if x > y{
		return x
	}
	return y
}

func main(){
	nums := [...]int{0,1,0,1,1,0}
	res := findMaxLength(nums[:])
	fmt.Println(res)
}