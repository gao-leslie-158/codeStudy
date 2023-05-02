package main

import (
	"fmt"
	"sort"
)

/*
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]] 输出对应的三元组值就行
*/

func threeSum(nums []int) [][]int {
	/*
	借鉴两数之和方法，但是此时不有序且会重复,所以先排序
	1-排序变为有序数组后
	2、先固定v1,在target = -v1，用双指针找b、c，
		但是这里不能像两数之和那样，因为这里值不唯一
	3、这里存在重复元素，避免枚举重复的值，要跳过
	*/
	//1、先排个序
	sort.Ints(nums)
	n := len(nums)
	res := make([][]int,0)
	for i1 := 0; i1 < n; i1 ++{
		//保持和上一层不同
		if i1 > 0 && (nums[i1] == nums[i1-1]){
			continue
		}
		//
		i3 := n-1	//第三层，右指针
		target := -nums[i1]
		for i2 := i1+1; i2 < n; i2++{
			//保持跟上一次枚举不同
			if i2 > i1+1 && (nums[i2] == nums[i2-1]){
				continue
			}
			for i2 < i3 && (nums[i2] + nums[i3] > target) {
				i3--
			}
			//如果重合，代表当前b+c > target
			//就算b再增大，也找不到满足b+c = target的，可以跳出当前b循环
			if i2 == i3{
				break
			}
			//相等则添加
			if nums[i1] + nums[i2] + nums[i3] == 0 {
				res = append(res, []int{nums[i1],nums[i2],nums[i3]})
			}
		}
	}
	return res
}

func main(){

	nums := [...]int{-1,0,1,2,-1,-4}
	res := threeSum(nums[:])
	fmt.Println(res)
}
