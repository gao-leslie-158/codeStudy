package main

import (
	"fmt"

	"golang.org/x/tools/go/analysis/passes/copylock"
)

var(
    path []int  //存放当前路径的集合
    res [][]int // 存放符合条件结果的集合
)

func combine(n int, k int) [][]int {
    path , res = make([]int, 0,k),make([][]int, 0)
    dfs(n,k,1)
    return res
}

func dfs(n int, k int, start int){
    if len(path) == k{
        tmp := make([]int,k)    
        copy(tmp,path)  // 切片是引用变量
        res = append(res, tmp)
        return
    }
    for i := start; i <= n - (k - len(path)) + 1; i ++{
        // if n - i + 1 < k - len(path){   // 剪枝？？
        //     //  列表中剩余元素个数 < 还需要的元素个数
        //     break
        // }
        path = append(path, i)  //加入
        dfs(n,k,i+1)
        path = path[:len(path) - 1] //推出
    }
}

func main(){

}