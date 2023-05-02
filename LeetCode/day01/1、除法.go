package main

import (
	"fmt"
)

func divide(a int, b int) int {
	//确定符号
	var symb int
	if (a > 0 && b >0) || (a < 0 && b < 0){
		symb = 1
	}else {
		symb = -1
	}
	if a == -1<<31 && b == -1{	//溢出情况
		return 1<<31 - 1
	}

	//变绝对值
	if a < 0 {
		a = -a
	}
	if b < 0{
		b = -b
	}

	shang , yu := 0 , a
	for yu >= b{	//商>余数时结束
		//定义左右区间
		l , h := 0 , 31
		mid := (l+h)>>1
		res := b*(1<<mid)
		for yu < res {
			h = mid
			mid = (l+h)>>1
			res = b*(1<<mid)
		}
		if yu >= res{
			yu -= res
			shang += (1<<mid)
		}
	}
	return symb*shang
}

func main(){

	res := divide(-1<<31,-1)
	fmt.Println(res)

}