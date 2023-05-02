package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main(){
	a := 8080
	b := "L" + strconv.Itoa(a)
	fmt.Println(b)
} 