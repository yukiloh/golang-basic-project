package exercise

//BubbleSort  冒泡排序
func BubbleSort(slice []int) []int {
	if len(slice) == 1 { //如果长度1就没必要交换
		return slice
	}
	var temp int
	flag := true
	for i := 0; i < len(slice); i++ {
		flag = true                           //每次循环都初始化,变量外置避免重复申明
		for j := 0; j < len(slice)-i-1; j++ { //因为冒泡每次遍历都会把最大数放到最后一位,所以每次外层遍历完后内层遍历可以少一位
			if slice[j] > slice[j+1] {
				temp = slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = temp
				flag = false //表面做过交换
			}
		}
		if flag {
			break //如果没做过则中断
		}
	}
	return slice
}
