package main

import "fmt"

//golang中的接口	接口是引用类型,默认值是nil

//定义一个接口
type Usb interface {
	//语法格式: 方法名(参数) 返回结果	这部分也称为函数的签名?
	Test()
	Insert()  //插入
	Extract() //拔出
}

//用于演示实现多个接口
type Power interface {
	PowerSupply()
}

//用于演示空接口
type Empty interface {
}

//定义一个结构体,去实现Usb接口(方法名相同)
type Device struct {
}

func (d Device) Insert() { //实现接口时,传入的类型也必须是定义时的类型,*Device和Device不可混用!
	fmt.Println("device insert...")
}
func (d Device) Extract() {
	fmt.Println("device extract...")
}
func (d Device) Test() {
	fmt.Println("device testing...")
}

//属于device独有的方法(keyboard中不存在的)
func (d Device) Specific() {
	fmt.Println("device specific testing...")
}

//实现Power接口中的方法
func (d Device) PowerSupply() {
	fmt.Println("powering")
}

//再定义一个结构体，实现Usb接口，用于动态数组的展示
type KeyBoard struct {
}

func (k KeyBoard) Insert() {
	fmt.Println("keyboard insert...")
}
func (k KeyBoard) Extract() {
	fmt.Println("keyboard extract...")
}
func (k KeyBoard) Test() {
	fmt.Println("keyboard testing...")
}

//定义调用接口的结构体
type Computer struct {
}

//定义一个c的方法,传入接口类型为Usb的参数,从而实现接口
func (c Computer) UseUsb(u Usb) { //此时传入的Usb便是多态参数,并非只能是Device类型
	u.Test()
	u.Insert()
	u.Extract()

	//进行类型断言
	if u, flag := u.(Device); flag {
		u.Test()
	}
}

//====================
//用于演示断言	定义一个坐标，x y
type Coordinate struct {
	x int
	y int
}

//演示类型断言 判断参数的类型
func CheckVarType(items ...interface{}) { //...代表可传入多个
	for index, value := range items {
		switch value.(type) { //v.(type)貌似是switch中特有的用法
		case int:
			fmt.Println("index: ", index, " var type is int")
		case bool:
			fmt.Println("index: ", index, " var type is bool")
		case string:
			fmt.Println("index: ", index, " var type is string")
		//其他省略,也可以用于判断指针和值
		case Device:
			fmt.Println("index: ", index, " var type is Device")
		case *Device:
			fmt.Println("index: ", index, " var type is *Device")
		}
	}
}

func main() {

	c := Computer{}
	d := Device{}

	//c取调用方法,传入符合接口类型结构体,结构体再调用实现接口的方法
	c.UseUsb(d)

	//golang的接口中只能存放方法,不允许存放任何变量(java是可以的)
	//实现接口中 所有 的方法(与java一致)
	//接口主要体现程序的多态和高内聚低耦合
	//golang中的接口不需要显示实现(关键词implement)
	//接口的继承不是按照类型名!是基于方法!	 如果上述案例中有一个名为Usb2的接口,内含形同的方法,在执行UseUsb时也会被调用!
	//因此golang接口的耦合度比java还低

	//接口本身不能创建实例,但可以指向一个 实现该接口类型的变量
	var d2 Device //类似于java中的实现
	//var u Usb	//此时调用u会发生空指针
	var u Usb = d2 //让接口指向一个实现该接口类型的变量
	u.Test()
	//自定义类型也可以实现接口,并非只是结构体(演示略)

	//一个结构体实现多个接口的示例,结构体device可以同时传输数据和供电
	var d3 Device
	var u1 Usb = d3 //实现接口u1和p1
	var p1 Power = d3
	u1.Test()
	p1.PowerSupply()

	//多继承也类似于java,如果A接口继承了B,C接口,那么实现A接口时,也必须要实现B,C中的方法
	//如果A继承BC时,但B继承了BD,C继承了CD,D被重复继承的情况则编译器会报错

	//空接口没有任何方法,所有类型都可以实现空接口
	var e Empty = d
	fmt.Println(e) //打印的结果也是空
	var i int = 1
	var e2 Empty = i //可以把任何一个变量赋值给空接口是一个经常用得到的特性!!!
	fmt.Println("empty 2: ", e2)

	//通过接口实现动态数组
	//常规的数组只能存放一个类型的元素，但是通过接口可以存放多种类型（Device和KeyBoard类型）！
	var usbArray [3]Usb
	usbArray[0] = Device{}
	usbArray[1] = KeyBoard{}
	fmt.Println(usbArray)

	//====================
	//类型断言	当不知道接口的类型时，要将其转为相应的类型，就需要断言
	var co1 Coordinate = Coordinate{1, 2} //定义一个结构体
	var emptyI interface{}                //定义一个空的接口
	emptyI = co1                          //让结构体给空接口赋值

	var co2 Coordinate        //再定义一个结构体
	co2 = emptyI.(Coordinate) //直接让空接口赋给2不可行，此处便需要断言
	//只有空接口原先是指向Co类型的结构体，才可以进行断言	比较严格，int32 → int64 也会报错

	fmt.Println(emptyI, co2)

	//判断是否断言成功（常用）
	//传统写法
	//co3,flag := emptyI.(Coordinate)
	//if flag {
	if co3, flag := emptyI.(Coordinate); flag { //合并写法

		fmt.Println("assert succeed,", co3)
	} else {
		fmt.Println("assert failed...")
	}

	//类型断言的案例1
	//当两个结构体(device&keyboard)都实现相同的接口时(usb)
	//其中一个接口(device)存在自己独有的方法,同时调用时为了避免另一个(keyboard)发生类型为定义
	//此时便可以通过类型断言来进行判断(具体代码在computer中)
	var usbArray2 [2]Usb
	usbArray2[0] = Device{}
	usbArray2[1] = KeyBoard{}

	var computer Computer
	for _, v := range usbArray2 {
		computer.UseUsb(v) //遍历调用
	}

	//类型断言的案例2
	//判断参数的类型
	CheckVarType(true, 1, "abc", &Device{}, Device{})

}
