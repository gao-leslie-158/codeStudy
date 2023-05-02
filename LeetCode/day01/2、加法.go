package main

import (
	"fmt"
	"strconv"
)

/*
给定两个 01 字符串 a 和 b ，请计算它们的和，并以二进制字符串的形式输出。

输入为 非空 字符串且只包含数字 1 和 0。
*/
func addBinary(a string, b string) string {

	ans := ""
	lenA , lenB := len(a) , len(b)
	n := maxLen(lenA,lenB)
	carry := 0
	for i := 0;i < n; i++{
		//carry = carry + a[] + b[]
		//索引得到的是字符的ASCII，必须减去字符‘0’才是值
		if i < lenA{
			carry += int(a[lenA-1-i] - '0')
		}
		if i < lenB{
			carry += int(b[lenB-1-i] - '0')
		}
		//ans = carry%2 + ans
		ans = strconv.Itoa(carry%2) + ans

		carry /= 2
	}
	if carry > 0{
		return "1" + ans
	}
	return ans
}
func maxLen(lenA int,lenB int) int {
	if lenA > lenB{
		return lenA
	}
	return lenB
}

func main(){

	res := addBinary("1001","01000001")
	fmt.Println(res)
}

