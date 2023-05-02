package main

import (
	"fmt"
)
func minWindow(s string, t string) string {

	/*滑动窗口+map
	1-首先用map1存t串元素对应出现的次数
	2-用map2存放当前窗口元素对应的出现次数
	3-如果map2值元素次数小于对应map出现的次数，end++
	4-如果map2元素值大于>=map1，start++
		且当前如果更新ans
	*/
	smap , tmap := map[byte]int{} , map[byte]int{}
	sn , tn := len(s) , len(t)
	start , end := 0 ,0
	minlen := math.MaxInt32
	left , right := -1 , -1	//最小窗口的左右边界

	// 统计t中元素出现的个数
	for  i := 0; i < tn; i++{
		tmap[t[i]] ++
	}
	//检查当前窗口smap元素出现次数是否大于tmap
	check := func() bool{
		for k , v := range tmap{
			if smap[k] < v{
				return false
			}
		}
		return true	//true为>=
	}

	for end < sn{
		if tmap[s[end]] > 0{
			smap[s[end]] ++
		}
		for check() {
			//更新最小窗口
			l := end - start + 1
			if l < minlen{
				minlen = l
				left , right = start , start + minlen
			}
			//如果左边界元素在tmap中，smap对应值-1
			if _ , ok := tmap[s[start]];ok{
				smap[s[start]] --
			}
			start ++
		}
		end ++
	}
	if left == -1{
		return ""
	}
	return s[left:right]
}


func main(){
	ans := minWindow("ADOBECODEBANC","ABC")
	fmt.Println(ans)
}