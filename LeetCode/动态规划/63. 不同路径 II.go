package main
import (
	"fmt"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m , n := len(obstacleGrid) , len(obstacleGrid[0])
	//考虑ob = [[0]]，起点即终点
	if m == 1 && n ==1 && obstacleGrid[0][0] == 0{
		return 1
	}
	//考虑起点或者终点有障碍
	if obstacleGrid[m-1][n-1] == 1 || obstacleGrid[0][0] == 1{
		return 0
	}
	dp := make([][]int,m)
	for i := range dp{
		dp[i] = make([]int, n)
	}
	for i := 1; i < m && obstacleGrid[i][0] == 0; i++{
		dp[i][0] = 1
	}
	for j := 1; j < n && obstacleGrid[0][j] == 0; j++{
		dp[0][j] = 1
	}
	for i := 1; i < m; i++{
		for j := 1; j < n; j++{
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}			
		}
	}
	return dp[m-1][n-1]
}

func main(){
	ob := [][]int{{0,1,0},{0,1,0},{0,0,0},}
	a := uniquePathsWithObstacles(ob)
	fmt.Println(a,ob)
}