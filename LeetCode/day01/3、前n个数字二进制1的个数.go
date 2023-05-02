package main
import (
	"fmt"
)

func countBits(n int) []int {
	/*
	给定一个非负整数 n ，请计算 0 到 n 之间的
	每个数字的二进制表示中 1 的个数，并输出一个数组。
	*/
	var a []int
	for i := 0; i <= n; i ++{
		count := 0
		for j := i; j != 0 ; j /= 2 {
			res := j % 2
			if res != 0{
				count += 1
			}
		}
		a = append(a, count)
	}
	return a
}

func main(){

	a := countBits(5)
	fmt.Println(a)
}