package main
import (
	"fmt"
)

func moveZeroes(nums []int)  {
	/*
	1-当前元素==0，count ++
	2-当前元素!=0，向前移动count位，且将当前元素置零
	3-count == 0时,当前元素不用置零
	*/
	n := len(nums)
	count := 0 //统计当前值前面val的个数
	for  i := 0; i < n; i ++{
		if nums[i] == 0{
			count ++
		}else {
			nums[i - count] = nums[i]
			if count != 0{
				nums[i] = 0
			}
		}
	}
}

func main(){

	
}