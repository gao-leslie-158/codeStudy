var (
	path []int
	res [][]int
)

func findSubsequences(nums []int) [][]int {
	path , res = make([]int, 0,len(nums)) , make([][]int, 0)
	backtracking(nums,0)
	return res
}

func backtracking(nums []int, startindex int){
	if len(path) >= 2{
		tmp := make([]int, len(path))
		copy(tmp,path)
		res = append(res, tmp)
	}
	used := make(map[int]bool, len(nums))	// 对同层元素去重
	for i := startindex; i < len(nums); i ++{
		if used[nums[i]] {	//同一个父节点的同层不能重复使用
			continue
		}
		if len(path) == 0 || nums[i] >= path[len(path)-1] {
			path = append(path, nums[i])
			used[nums[i]] = true
			backtracking(nums,i+1)
			path = path[:len(path)-1]
		}
	}
}