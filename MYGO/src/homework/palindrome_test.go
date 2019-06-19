package homework

import (
	"testing"
)

//TestPalindrome1 判断回文函数是否正确
func TestPalindrome(t *testing.T) {
	testStr := []string{"abcba", "灯油少油灯", "Madam,I’mAdam"}
	for i, v := range testStr {
		result := Palindrome(v)
		if !result {
			t.Logf("第%d个校验失败！", i)
			t.Fail()
		} else {
			t.Logf("第%d个校验成功！", i)
		}
	}
}

func BenchmarkPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 没有想到优化方案
		// BenchmarkPalindrome1-8           2000000               602 ns/op               0 B/op          0 allocs/op
		Palindrome("有月的长ha中21nihaoaoahin12中ah长的月有")
	}
}
