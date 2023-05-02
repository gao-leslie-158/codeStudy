package main

import (
	"fmt"
	"strconv"
)

func wordBreak(s string, wordDict []string) bool {
	n , m := len(s) , len(wordDict)
	wordDictSet := make(map[string]bool, m)	//设置一个单词字典map，用来判断单词是否存在
	for _ , w := range wordDict{
		wordDictSet[w] = true
	}
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++{
		for j := 0; j < i; j ++{
			if dp[j] == true && wordDictSet[s[j:i]]{
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}


func main(){

} 