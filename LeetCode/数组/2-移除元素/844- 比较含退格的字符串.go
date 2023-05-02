
func backspaceCompare1(s string, t string) bool {
	/*双指针法：分别指向末尾
	1-逆序遍历字符，skip表示当前要删除的字符数
	2-当字符为退格符"#"，skip+1
	3-当字符为普通字符：
		·skip == 0 ，不需要删除
		·skip != 0 ，删除当前字符 

	逻辑判断：
	false：
		均没遍历完，当前值不等
		至少有一方先遍历完
	true:
		同时遍历完
	*/

	skipS , skipT := 0 , 0
	si , ti := len(s) - 1, len(t) - 1
	for si >= 0 || ti >= 0{
		for si >= 0{
			if s[si] == '#'{
				skipS ++
				si --
			}else if skipS > 0{
				skipS --
				si --
			}else {
				break	//不为退格符且skipS==0
			}
		}
		for ti >= 0{
			if t[ti] == '#'{
				skipT ++
				ti --
			}else if skipT > 0{
				skip --
				ti --
			}else {
				break
			}
		}
		if si >= 0 && ti >= 0{
			if s[si] != t[ti]{	//如果没有遍历完，且删了后当前两值不等
				return false
			}
		}else if si >= 0 || ti >= 0{
			return false	//至少有一方先遍历完
		}
		si --
		ti --
	}
	return true	//遍历完，且最后相等
}

func build(str string) string{
	var stackS []byte
	for i := range str{
		if str[i] != '#'{
			stackS = append(stackS,str[i])
		}else if len(stackS) > 0{
			stackS = stackS[:len(stackS) -1]
		}
	}
	return string(stackS)
}

func backspaceCompare(s string, t string) bool {
	/*利用栈，从前往后遍历
	1-为退格符，且栈不为空，弹出栈顶
	2-为普通字符，入栈
	最后比较两个栈内元素是否相等
	*/
	return build(s) == build(t)
}