

func slidingWindow(s string){

	n := len(s)
	sum := 0
	/*定义左右窗口边界*/
	start , end := 0 , 0
	
	/*右边扩大窗口*/
	for end < n{
		//对加入窗口的元素操作
		sum += s[end]

		//右移窗口
		end ++

		//进行窗口数据内的更新

		//左边收缩窗口,收缩窗口条件
		for sum > 78{
			//移出窗口元素
			
			//左移窗口
			start ++

			//进行窗口内的数据更新
		}
	}
}