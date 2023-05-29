package main

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
输入：s = "()"
输出：true
输入：s = "()[]{}"
输出：true
输入：s = "(]"
输出：false
*/

func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	//num := len(s) / 2
	//var sbyte []byte
	//sbyte = []byte(s)
	return false
}
func main() {

}
