package main

import (
	"fmt"
)

// 或运算：全0为0
func singleNumber(nums []int) int {
	// 转化为int32的原因是因为int在64位的操作系统中数值范围太大了，
	// 会将答案的符号位误认为是数值位
	res :=int32(0)	//结果
	for i := 0; i < 32; i++{
		curSum := int32(0)	//当前第i位累积和
		for _ , num := range nums{
			//得到第i位的累加结果
			curSum += int32(num >> i) & 1	//(x >> i) & 1 得到第i个二进制位
		}
		if curSum % 3 > 0{
			res = res | (1<<i)
		}
	}
	return int(res)
}

func singleNumber1(nums [10]int) int {
	var res int
	var count map[int]int
	count = make(map[int]int)
	for _ , num := range nums{
		count[num] += 1
	}
	for k , v := range count{
		if v == 1{
			res = k
		}
	}
	return res
}

func main(){

	a := [...]int{2,2,2,6,6,6,10,10,10,-5}
	res := singleNumber1(a)
	fmt.Println(res)

}