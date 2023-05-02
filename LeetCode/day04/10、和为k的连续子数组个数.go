package main

import (
	"fmt"
)

/*
输入:nums = [1,1,1], k = 2
输出: 2
解释: 此题 [1,1] 与 [1,1] 为两种不同的情况
*/

func subarraySum1(nums []int, k int) int {
	/*枚举*/
	/*
	*/
	count , n := 0 , len(nums)
	for start := 0; start < n; start ++{
		sum := 0
		for end := start; end < n; end ++{
			sum += nums[end]
			if sum == k{
				count ++
			}
		}
	}
	return count
}

func subarraySum(nums []int, k int) int {
	/*
	这道题目非常简洁，就是求数组中何为整数k的连续子数组个数。
	如果这道题的取值没有负数，那就是标准的滑窗问题，但因为有了负数，
	滑窗思想不能用了。 通过分析，这道题应该属于我们上面列举四种情况的最后一种。
	具体思路如下：
	
		初始化一个空的哈希表和pre_sum=0的前缀和变量
		设置返回值ret = 0，用于记录满足题意的子数组数量
		循环数组的过程中，通过原地修改数组的方式，计算数组的累加和
		将当前累加和减去整数K的结果，在哈希表中查找是否存在
		如果存在该key值，证明以数组某一点为起点到当前位置满足题意，ret加等于将该key值对应的value
		判断当前的累加和是否在哈希表中，若存在value+1，若不存在value=1
		最终返回ret即可
	但在这里要注意刚才说到的前缀和边界问题。 我们在计算这种场景时，
	需要考虑如果以数组nums[0]为开头的连续子数组就满足题意呢？
	此时候我们的哈希表还是空的，没办法计算前缀和！所以遇到这类题目，
	都需要在哈希表中默认插入一个{0:1}的键值对，
	用于解决从数组开头的连续子数组满足题意的特殊场景。 下面就开始解题吧！
	*/
		count , presum := 0 , 0
		presum_map := map[int]int{}
		presum_map[0] = 1
		for _ , v := range nums{
			presum += v
			if _ , ok := presum_map[presum - k]; ok{
				count += presum_map[presum - k]
			}
			presum_map[presum] += 1
		}
		return count
	}


func main(){
	// nums := [...]int{1,1,1,2,3}
	nums := [...]int{1,1,1}
	res := subarraySum(nums[:],2)
	fmt.Println(res)
	
}