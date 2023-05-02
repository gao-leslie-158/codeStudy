package main
import (
	"fmt"
)

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp{
		dp[i] = make([]int, n+1)
	}
	// 遍历字符串
	for _ , str := range strs{
		// 统计每个字符串0-1个数
		zeroNum , oneNum := 0 , 0
		for _ , c := range str{
			if c == '0'{
				zeroNum ++
			}else {
				oneNum ++
			}
		}
		for i := m; i >= zeroNum; i--{
			for j := n; j >= oneNum; j--{
				dp[i][j] = max(dp[i][j],dp[i-zeroNum][j-oneNum]+1)
			}
		}
	}
	return dp[m][n]
}

func max(x,y int) int{
	if x > y{
		return x
	}
	return y
}

func main(){
	strs := []string{"10", "0001", "111001", "1", "0"}
	m , n := 5 , 3
	res := findMaxForm(strs,m,n)
	fmt.Println(res)
}