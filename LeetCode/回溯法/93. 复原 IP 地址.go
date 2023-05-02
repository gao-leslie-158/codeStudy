package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	path []string
	res []string
)

func restoreIpAddresses(s string) []string {
	path , res = make([]string,0,len(s)) , make([]string, 0)
	backtracking(s,0)
	return res
}

func backtracking(s string,startindex int){
	if len(path) == 4{	// ip地址有4段
		if startindex == len(s){	// 分割线到最后，全部分割完，且分成了合法的四份
			str := strings.Join(path,".")
			res = append(res, str)
		}
		return
	}
	
	for i := startindex; i < len(s); i ++{
		if i != startindex && s[startindex] == '0'{	//含有符号0，无效
			break
		}
		// startindex 当前层开始的地方
		str := s[startindex:i+1]	// 当前截取的段
		num , _ := strconv.Atoi(str)
		if num >= 0 && num <= 255{	//符合条件，进入下一层
			path = append(path, str)
			backtracking(s,i+1)
			path = path[:len(path)-1]
		}else{
			break
		}
	}
}

func main(){

} 