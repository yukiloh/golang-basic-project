package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//接口的练习,对结构体的切片进行排序

//1.声明结构体
type Person struct {
	Name string
	Age  int
}

//2.声明Person类型的切片
type PersonSlice []Person

//3.实现sort包中的Sort接口(3个方法)
//Len:集合中元素个数
func (p PersonSlice) Len() int {
	return len(p)
}

//Less:2个元素互相比较大小,返回bool,大于小于决定了降序升序
func (p PersonSlice) Less(i, j int) bool {
	return p[i].Age < p[j].Age //按照年龄进行排序
	//return p[i].Name < p[j].Name		//按照姓名排序(有问题,意思意思)
}

//Swap:交换2个元素
func (p PersonSlice) Swap(i, j int) {
	//交换元素传统的写法
	//temp := p[i]
	//p[i] = p[j]
	//p[j] = temp

	//golang独有的写法
	p[i], p[j] = p[j], p[i]
}

func main() {
	//测试接口排序
	//创建一个切片,并定义10个元素
	var ps PersonSlice
	for i := 0; i < 10; i++ {
		p := Person{
			Name: fmt.Sprintf("Person%v", rand.Intn(11)+1),
			Age:  rand.Intn(20) + 1,
		}
		ps = append(ps, p)
	}

	//测试打印
	fmt.Println(ps)

	//调用sort.Sort接口进行排序
	sort.Sort(ps)

	fmt.Println(ps)

}
