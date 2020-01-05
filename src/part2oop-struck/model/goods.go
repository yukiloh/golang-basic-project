package model

import "fmt"

//用于展示继承	如果此处(匿名)结构体小写，哪怕被公有结构体继承也无法访问
type Goods struct {
	Name  string
	Price int
}

func (g *Goods) ShowGoods() {
	fmt.Println("goods show... name:", g.Name, " price:", g.Price)
}

type Fruit struct {
	Goods        //通过匿名结构体来使用继承，简洁,如果添加了结构体名则是有名结构体!
	Quantity int //数量
}

func (f *Fruit) ShowFruit() {
	fmt.Println("fruit show... name:", f.Name, " price:", f.Price, " quantity:", f.Quantity)
}

//通过idea自动创建，所以说一般情况都是传指针
func (f *Fruit) Buy(i int) {
	f.Quantity += i
}

//蔬菜
type Greens struct {
	Goods    Goods //有名结构体
	Quantity int   //数量
	int            //结构体中也允许基本类型,但你这是小写外面无法访问.不允许重复匿名基本类型
}
