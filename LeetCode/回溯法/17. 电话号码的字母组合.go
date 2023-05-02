package main

import (
	"fmt"
	"strings"
)

var (
	m []string	// 数字->字母映射
	res []string	//最终结果
	path []byte	//叶子结点的结果，也就是路径上的结果
)

func letterCombinations(digits string) []string {
	m = []string{"","","abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"}
	path , res = make([]byte, 0) , make([]string, 0)
	if digits == ""{
		return res
	}
	backtrcking(digits,0)
	return res
}

func backtrcking(digits string,index int){
	if len(path) == len(digits){
		tmp := string(path)
		res = append(res, tmp)
		return
	}
	digit := int(digits[index] - '0') //将index对应的数字转为int
	str := m[digit]		// 取数字映射的字符串
	for i := 0; i < len(str); i ++{	//遍历当前数字对应的字符
		path = append(path, str[i])
		backtrcking(digits,index+1)
		path = path[:len(path) - 1]
	}
}

func main(){

}