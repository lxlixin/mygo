package basicknowledge

import (
	"fmt"
	"time"
)

//time包中有好多关于时间的操作，go语言中的format跟别的语言不一样，需要使用2006/01/02 15:04:05.000
func formatTime(t time.Time) {
	// t := time.Time.Now()
	fmt.Println(t.Day())
	fmt.Println(t.Format("2006/01/02 15:04:05"))
	t1 := t.UnixNano()
	fmt.Println(t1)
}
