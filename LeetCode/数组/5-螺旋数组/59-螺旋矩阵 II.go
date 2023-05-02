package main
import (
	"fmt"
)

func generateMatrix(n int) [][]int {
	/*直接用模拟
	对于每层，从左上方开始以顺时针的顺序填入所有元素。假设当前层的左上角位于 (top,left)
	从左到右填入上侧元素，依次为 (top,left)到 (top,right)。
	从上到下填入右侧元素，依次为 (top+1,right) 到 (bottom,right)

	如果 left < right 且 top < bottom
	则从右到左填入下侧元素，依次为 (bottom,right−1)到(bottom,left+1)
	从下至上填入下列元素，(bottom,left)到(top+1,left)
	*/
	ans := make([][]int, n)
    for i := range ans {
        ans[i] = make([]int, n)
    }

	t , b , l , r := 0 , n-1 , 0 , n-1 //上下左右
	for num := 1; num <= n*n;{
		for i := l; i <= r; i ++{
			ans[t][i] = num
			num ++
		}
		t ++
		for i := t; i <= b; i ++{
			ans[i][r] = num
			num ++
		}
		r --
		for i := r; i >= l; i --{
			ans[b][i] = num
			num ++
		}
		b --
		for i := b; i >= t; i --{
			ans[i][l] = num
			num ++
		}
		l ++ 
	}
	return ans
}

func main(){
	ans := generateMatrix(3)
	fmt.Println(ans)
}