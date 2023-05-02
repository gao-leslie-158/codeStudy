package main

import (
	"fmt"
	"strings"
)

/*
输入: words = ["abcw","baz","foo","bar","fxyz","abcdef"]
输出: 16
解释: 这两个单词为 "abcw", "fxyz"。它们不包含相同字符，且长度的乘积最大。
*/

func maxProduct1(words []string) int {
	/*暴力解*/
	lenW := len(words)
	maxLen := 0	//最大长度
	for i := 0; i < lenW; i++{
		for j := i+1; j < lenW; j++{
			if !strings.ContainsAny(words[i],words[j]){
				newLen := len(words[i]) * len(words[j])
				if newLen > maxLen{
					maxLen = newLen
				}
			}
		}
	}
	return maxLen
}

func maxProduct(words []string) int {
	/*采用掩码然后位运算&*/
	n := len(words)
	mask := make([]int,n)	//掩码切片
	//先得到每个单词mask
	for i , word := range words{
		for _ , s := range word{
			mask[i] |= 1 << (s - 'a')	//用或 | 累积移位
		}
	}
	//比较每个单词mask
	maxLen := 0	//最长长度乘积
	for i , x := range mask{
		for j , y := range mask[:i]{	//mask[:i]这个用法不错
			//如果采用mask[i+1:]这里索引下标就会乱，因为这是一个新的切片
			// & 不同为0 ，相同为1
			if x&y == 0{
				newLen := len(words[i]) * len(words[j])
				if newLen > maxLen{
					maxLen = newLen
				}	
			}
		}
	}
	return maxLen
}


func main(){

	words := [...]string{"a","ab","abc","d","cd","bcd","abcd"}
	// lenW := len(words)
	res := maxProduct(words[:])
	fmt.Println(res)
}