package basicknowledge

import (
	"testing"
)

func TestAdd(t *testing.T) {
	x := 1
	y := 2
	result := add(x, y)
	if result != 3 {
		t.Log("这个方法计算不正确呀！赶紧改一改！")
	}
}
