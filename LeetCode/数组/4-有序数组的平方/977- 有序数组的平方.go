package main

import (

)

func main(){
	sortedSquares([]int{-1,-3,0,2,3,5,8})
	
}

func sortedSquares1(nums []int) []int {
	/*简单点，先平方，后排序*/
	// 平方
	for i , v := range nums{
		nums[i] = v * v
	}
	// sort.Ints(nums)
	return nums
}

func sortedSquares(nums []int) []int {
	/*双指针
	利用双指针分别指向0 ，n-1，每次比较两个指针对应的数，选择平方较大
	的那个逆序的放入ans中
	*/
	n := len(nums)
	i , j := 0 , n-1
	ans := make([]int, n)
	for pos := n - 1; pos >= 0; pos --{
		if v , w := nums[i]*nums[i],nums[j]*nums[j]; v > w{
			ans[pos] = v
			i ++
		}else {
			ans[pos] = w
			j --
		}
	}
	return ans
}