package main

import (
	"fmt"
	"strings"
)

var (
	chessbord [][]string
	res [][]string 
)

func solveNQueens(n int) [][]string {
	chessbord , res = make([][]string, n) , make([][]string, 0)
	// 初始化棋盘
	for i := 0; i < n; i++{
		chessbord[i] = make([]string, n)
	}
	for i := 0; i < n; i++{
		for j := 0; j < n; j++{
			chessbord[i][j] = "."
		}
	}
	backtracking(n,0)
	return res
}

func backtracking(n int, row int){
	if row == n{
		tmp := make([]string, n)
		for i , rowStr := range chessbord{
			tmp[i] = strings.Join(rowStr,"") //rowStr是[".",".","Q","."]这种形式
		}
		res = append(res, tmp)
		return 
	}
	for colum := 0; colum < n; colum++{
		if isValid(n,row,colum,chessbord){
			chessbord[row][colum] = "Q"
			backtracking(n,row+1)
			chessbord[row][colum] = "."
		}
	}
}

func isValid(n,row,colum int, chessbord [][]string) bool{
	// 1、判断行合法，因为单层遍历就已经每层取不同的元素了，不用判断
	// 2、判断列合法
	for j := 0; j < row; j++{
		if chessbord[j][colum] == "Q"{
			return false
		}
	}
	// 3、判断45°角合法
	for i,j := row-1,colum-1; i >= 0 && j >= 0; i,j = i-1, j-1{
		if chessbord[i][j] == "Q"{
			return false
		}
	}
	// 4、判断135°角合法
	for i,j := row-1,colum+1; i >= 0 && j < n; i,j = i-1,j+1{
		if chessbord[i][j] == "Q"{
			return false
		}
	}
	return true
}

func main(){

} 