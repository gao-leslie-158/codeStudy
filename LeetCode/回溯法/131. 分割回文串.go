package main
import (
	"fmt"
)

var (
	path []string
	res [][]string
)

func partition(s string) [][]string {
	path , res = make([]string, 0) , make([][]string, 0)
	backtracking(s,0)
	return res
}

func backtracking(s string,startindex int){
	if startindex == len(s){
		tmp := make([]string, len(path))
		copy(tmp,path)
		res = append(res, tmp)
		return
	}

	for i := startindex; i < len(s); i ++{
		str := s[startindex:i+1]
		if isPalindrome(str) == true{	//是回文子串
			path = append(path, str)	
			backtracking(s,i+1)			// 找i+1为起始位置的子串
			path = path[:len(path)-1]	// 回溯，弹出本次填充的子串
		}
	}
}

func isPalindrome(s string) bool{
	// 双指针法
	for i , j := 0 , len(s)-1; i < j; i , j = i+1 , j-1{
		if s[i]  != s[j]{
			return false
		}
	}
	return true
}

func main(){

}