package main
import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	/*每个不等于val的值，向前移动前面count（前面val值的个数）*/
	n := len(nums)
	count := 0 //统计当前值前面val的个数
	for  i := 0; i < n; i ++{
		if nums[i] == val{
			count ++
		}else{
			nums[i - count] = nums[i]
		}
	}
	return n - count
}

func main(){

	
}