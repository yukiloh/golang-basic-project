package exercise

import "fmt"

var num1, num2, num3 int = 1, 2, 3

//GetMaxFromTwo 求2个数的最大值
func GetMaxFromTwo() {
	if num1 > num2 {
		num2 = num1
	}
	fmt.Print(num2)
}

//GetMaxFromThree 求3个数的最大值
func GetMaxFromThree() {
	if num1 > num2 {
		num2 = num1
	}

	if num2 > num3 {
		num3 = num2
	}

	fmt.Print(num3)
}

//GetMaxFromArr 不知道怎么定义动态数组长度,暂时定5个
func GetMaxFromArr(arrInt [5]int) {
	max := 0
	for _, i := range arrInt {
		if i > max {
			max = i
		}
	}
	fmt.Println("max:", max)
}
