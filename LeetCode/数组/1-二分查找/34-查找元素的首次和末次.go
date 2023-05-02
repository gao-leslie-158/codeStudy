package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	/*
	1. 首先在nums中用二分查找找target
	2. 如果没找到，返回 -1 , searchRange返回[-1,-1]
	3. 如果找到，利用左右指针滑动找左右边界。 searchRange返回[left , right]
	*/
	index := binarySearch(nums,target)
	if index == -1{
		return []int{-1,-1}
	}else{
		n := len(nums)
		left , right := index , index
		//往左边滑动
		for ; left - 1 >= 0  && nums[left - 1] == target; left --{}
		//往右边滑动
		for ; right + 1 < n && nums[right + 1] == target; right ++{}
		return []int{left,right}
	}
}

func binarySearch(nums []int, target int) int {
	left , right := 0 , len(nums) -1
	for left <= right{
		mid := left + ((right - left) >> 1)
		if nums[mid] > target{
			right = mid -1
		}else if nums[mid] < target{
			left = mid + 1
		}else{
			return mid
		}
	}
	return -1
}

func main(){
	nums := []int{0} 
	target := 0
	res := searchRange(nums,target)
	fmt.Println(res)
}

