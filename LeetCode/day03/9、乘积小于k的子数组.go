package main
import (
	"fmt"
)

/*连续子数组
输入: nums = [10,5,2,6], k = 100
输出: 8
解释: 8 个乘积小于 100 的子数组分别为: [10], [5], [2], [6], [10,5], [5,2], [2,6], [5,2,6]。
需要注意的是 [10,5,2] 并不是乘积小于100的子数组。
*/
func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	/*滑动窗口*/
	/*
		k为0值，直接返回0
		两个指针start、end，一个乘积当前窗口乘积product
		start=end=0，当前窗口乘积product < k 时，个数+1，end右滑
		product >= k 时，product = product / nums[start]，start右滑
		新增元素为0时，start直接跳到0后，end也后移
	*/
	prod, start := 1, 0
	for end, num := range nums {
		prod *= num
		for ; start <= end && prod >= k; start++ {
			prod /= nums[start]
		}
		ans += end - start + 1	//以右端点end为结尾的合法子数组个数
	}
	return
}


func main(){

	nums := [...]int{10,5,2,6}
	res := numSubarrayProductLessThanK(nums[:],100)
	fmt.Println(res)
}