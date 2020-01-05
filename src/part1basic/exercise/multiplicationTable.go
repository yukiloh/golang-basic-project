package exercise

import "fmt"

//MultTable 九九乘法表
func MultTable() {
	for a := 1; a < 10; a++ {
		for b := a; b < 10; b++ {
			fmt.Printf("%d*%d = %d  ", a, b, a*b)
		}
		fmt.Println("")
	}
}
