package main

import (
	"fmt"
	"strconv"
)

/*
思路:(121是回文,-121不是)
1. 判断是不是负数，是负数直接输出false，如果是正数再处理
*/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	//123,0-->12,3-->1,32-->0,321
	num, res := x, 0
	for num > 0 {
		res = res*10 + num%10
		num = num / 10
	}
	return res == x
}

/*思路2(麻烦，因为golang中s[]返回的是asci值):
1. 转换成字符串在进行判别
*/
func isPalindrome_2(x int) bool {
	var b bool
	s := strconv.Itoa(x)
	fmt.Printf("%v\n", s[1])

	return b
}

func main() {
	a := 121
	fmt.Printf("%v", isPalindrome(a))
}
