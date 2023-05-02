package main

import(
	"fmt"
)

func search(nums []int, target int) int {

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

	return -1
}

func main(){
	nums := []int{-1,0,3,5,9,12} 
	target := 15
	res := search(nums,target)
	fmt.Println(res)
}

