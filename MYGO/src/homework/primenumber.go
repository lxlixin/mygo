package homework

import (
	"fmt"
)

// 获取素数
func GetPrimeNumber(startN, endN int) {

	for ; startN <= endN; startN++ {
		if startN < 4 {
			fmt.Println(startN)
			continue
		}
		for i := 2; i*i < startN; i++ {
			if (i+1)*(i+1) > startN {
				counter := 0
				for j := 2; j <= i; j++ {
					if startN%j != 0 {
						counter++
					}
				}
				if counter >= i-1 {
					fmt.Println(startN)
				}
			}
		}
	}
}

//九九乘法表
func Get99(startN, endN int) {
	for i := startN; i <= endN; i++ {
		for j := startN; j <= i; j++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}
}
