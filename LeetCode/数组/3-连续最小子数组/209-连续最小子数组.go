package main
import (
	"fmt"
)

func minSubArrayLen1(target int, nums []int) int {
	//暴力解
	n := len(nums)
	if n == 0{
		return 0
	}
	ans := n + 1
	for i := 0; i < n; i ++{
		sum := 0
		for j := i; j < n; j ++{
			sum += nums[j]
			if sum >= target {
				ans = minxy(ans , j - i + 1)
				break
			}
		}
	}
	if ans == n + 1{
		return 0
	}
	return ans 
}

func minxy(x int ,y int) int{
	if x < y{
		return x
	}
	return y
}

func minSubArrayLen(target int, nums []int) int {
	/*滑动窗口
	start ，end, sum , ans , target
	1-end < n , sum += nums[end] , end 右移
	2-sum >= target , start右移，sum -= nums[start],更新ans 
	*/
	n := len(nums)
	if n == 0 {
		return 0
	}
	sum , ans := 0 , n + 1
	start , end := 0 , 0
	for end < n {
		sum += nums[end]
		for sum >= target{
			l := end - start + 1
			if l < ans{
				ans = l
			}
			sum -= nums[start]
			start ++
		}
		end ++
	} 

	if ans == n + 1{
		return 0
	}
	return ans
}

func main(){
	ans := minSubArrayLen(11,[]int{1,2,3,4,5})
	fmt.Println("ans:",ans)
}