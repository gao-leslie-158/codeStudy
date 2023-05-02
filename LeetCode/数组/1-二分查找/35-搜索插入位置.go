package main

import(
	"fmt"
)

func searchInsert(nums []int, target int) int {
	// 二分查找，没找到时left > right,返回left
	left , right := 0 , len(nums) -1

	for left <= right{
		mid := (left + right) / 2
		if nums[mid] > target{
			right = mid -1
		}else if nums[mid] < target{
			left = mid + 1
		}else{
			return mid
		}
	}

	return left
}

func main(){
	nums := []int{-1,0,3,5,9,12} 
	target := 15
	res := searchInsert(nums,target)
	fmt.Println(res)
}

