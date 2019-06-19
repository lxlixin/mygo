package homework

import (
	"strings"
)

//Palindrome 判断输入的字符串是否是回文
func Palindrome(str string) bool {
	str = strings.ToLower(str)
	st := []rune(str)
	l := len(st)
	var index int
	m := l / 2
	var s1, s2 rune

	for index < m {
		s1 = st[index]
		s2 = st[l-index-1]
		if s1 != s2 {
			return false
		}
		index++
	}
	return true
}
