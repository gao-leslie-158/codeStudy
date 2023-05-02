package main
import (
	"fmt"
)
var (
	wordDictSet map[string]bool
)

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	wordDictSet = make(map[string]bool,len(wordDict))
	for _ , w := range wordDict{
		wordDictSet[w] = true
	}
	memory := make([]bool, n)
	return backtracking(s,wordDict,memory,0)
}

func backtracking(s string,wordDict []string,memory []bool,startindex int) bool{
	n := len(s)
	if startindex >= n{
		return true
	}
	if memory[startindex] == false{
		return memory[startindex]
	}
	for i := startindex+1; i <= n; i ++{
		str := s[startindex:i-startindex+1]
		if wordDictSet[str] && backtracking(s,wordDict,memory,i){
			return true
		}
	}
	memory[startindex] = false
	return false
}

func main(){
	s := "abcasdha"
	str := s[3:5]
	fmt.Println(str)
} 