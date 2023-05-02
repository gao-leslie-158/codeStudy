package main
import (
	"fmt"
)

func totalFruit(fruits []int) int {
	/*题目翻译成人话就是：找至多包含两种元素的最长子串，返回长度
	滑动窗口：
	left、right表示窗口的左右边界，用map存储当前窗口每个元素出现的次数，
	1---right每右移一次，将对应元素加入map，如果map中超过两个key:value
		则将fruit[left]从map中移除，left右移。
	2---将fruit[left]从map中移除后，如果元素对应次数减少为0，则要从map中移除键值。
	*/
	var ans int
	left := 0
	cntMap := map[int]int{}
	for right , x := range fruits{
		cntMap[x] ++
		for len(cntMap) > 2{
			y := fruits[left]
			cntMap[y] --
			if cntMap[y] == 0{
				delete(cntMap,y)
			}
			left ++
		}
		l := right - left + 1
		if l > ans{
			ans = l
		}
	}
	return ans
}

func main(){
	ans := totalFruit([]int{0,1,2,1})
	fmt.Println(ans)
}