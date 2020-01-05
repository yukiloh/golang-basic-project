package exercise

import "fmt"

//ExchangeAB 不用中间值,交换ab的值
func ExchangeAB() {
	a := 10
	b := 20

	a = a + b
	b = a - b
	a = a - b

	fmt.Print(a, ",", b)
}
