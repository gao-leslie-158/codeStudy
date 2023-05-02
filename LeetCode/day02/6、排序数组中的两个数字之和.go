package main
import (
	"fmt"
)

/*
输入：numbers = [1,2,4,6,10], target = 8
输出：[1,3]
解释：2 与 6 之和等于目标数 8 。因此 index1 = 1, index2 = 3 。
*/

func twoSum3(numbers []int, target int) []int {
	/*双指针
	l,h 初始为两端
	1、if target > l+h值 ==> l++
	2、if target < l+h值 ==> h--
	*/
	l , h := 0 , len(numbers) - 1
	for l != h{
		sum := numbers[l] + numbers[h]
		if target > sum{
			l++
		}else if target < sum {
			h--
		}else {
			return []int{l,h}
		}
	}
	return []int{1,-1}
}

func twoSum2(numbers []int, target int) []int {
	/*利用有序数组性质，采用二分查找思想*/
	/*
	先固定一个数v1，从v1右边找第二个数v2：v2 = target - v1
	如果 v2 < mid值  ==> high = mid -1
	如果 v2 > mid值  ==> low = mid +1
	*/
	for i := 0; i < len(numbers); i++{
		l , h := i+1 , len(numbers)-1
		for l <= h{
		// 不使用mid=(h+l)/2原因就是避免如果两个数都是超级大就会引起overflow
			mid := (h - l)/2 + l
			v2 := target - numbers[i]
			if v2 == numbers[mid]{
				return []int{i,mid}
			}else if v2 > numbers[mid]{
				l = mid + 1
			}else {
				h = mid - 1
			}
		}
	}
	return []int{1,-1}
}

func twoSum1(numbers []int, target int) (index []int) {
	//直接无脑暴力
	for i , v1 := range numbers[:]{
		for j , v2 := range numbers[:i]{	
			if v1 + v2 == target{
				index = append(index, j,i)		//i > j	
			}
		}
	}
	return
}

func main(){

	numbers := [...]int{1,2,4,6,10}
	res := twoSum3(numbers[:],6)
	fmt.Println(res)
}