package main

import (
	"fmt"
	"reflect"
)

//关于反射
func main() {
	//反射可以在程序运行时动态获取变量的各种信息
	//如果变量为结构体，则还可以获取其本身的字段，方法等
	//通过反射可以修改变量的值，或者调用其方法

	//反射时，变量、interface{}、reflect.Value一般可以互相转换（常用）
	//变量 → 作为参数，传递至函数中，通过空接口interface{}接收 → 通过reflect.Value(),获得反射value类型,进行数据处理
	//返回会变量时,通过reflect.Value() → 获得v.interface()类型  → 通过类型断言,转为变量

	//获取反射的演示

	//通过反射获取int类型（基本类型）
	var i = 10
	showRefInt(i)

	//通过反射获取结构体中的值
	goudan := student{"狗蛋",12,"M"}
	showRefStruct(goudan)

	//通过反射修改基本类型的值
	changeValueByRef(&i)		//注意，需要传入地址
	fmt.Println(i)

	//通过反射获取一个结构体的类型
	GetAnyStructInfo(goudan)
	//通过反射修改任意结构体中字段的值
	ChangeAnyStructData(&goudan)	//同样需要传入地址

	fmt.Println(goudan)	//打印查看结果
}


//修改任意类型结构体中的值
func ChangeAnyStructData(stu *student) {	//golang中的参数传递都是值传递（传入值的拷贝），如果要修改值，必须传入指针，否则只影响复制的对象
	refValPtr := reflect.ValueOf(stu)		//获取反射的指针；另已测试，TypeOf没有SetString函数
	refValPtr.Elem().FieldByName("Name").SetString("猫蛋")	//根据字段的name，修改字段的值
}


//获取一个任意结构体的信息
func GetAnyStructInfo(i interface{})  {
	refVal := reflect.ValueOf(i)
	refType := reflect.TypeOf(i)
	refKind := refVal.Kind()       //获取传入参数的值的类型kind
	if refKind != reflect.Struct { //右侧为常量值,源码采用iota来
		fmt.Println("param is not struct")
		return
	}

	//获取字段数量,val和type都可获取
	//fmt.Println(refVal.NumField())
	//fmt.Println(refType.NumField())

	//获取tag,注意,tag只可以通过type获取	所以序列化的时候也是通过反射来获取key的名称,也即为什么tag必须要指定使用json
	fmt.Println(refType.Field(0).Tag.Get("json"))


	//===================================
	//获取方法数量,val和type都都可以
	fmt.Println(refVal.NumMethod())
	fmt.Println(refType.NumMethod())

	//golang反射获取方法时,方法遵照的排序是方法的名字,abcd这样
	refVal.Method(1).Call(nil)	//第一个方法：

	//如果需要传参，则需要先定义reflect.Value类型的切片
	var params = []reflect.Value{
		reflect.ValueOf("nil"), 		//内部的值也需要通过ValueOf转换
		reflect.ValueOf(12),
	}
	params = append(params,reflect.ValueOf("F"))		//也可以通过append的方式来添加
	result := refVal.Method(0).Call(params)			//注意，返回的结果也是[]reflect.Value 切片！
	fmt.Println(result[0].Bool())						//通常情况此处需要添加断言，省略了

}


func changeValueByRef(i interface{}) {
	rValue := reflect.ValueOf(i)		//因为传入的是地址,此时rVal类型是指针ptr
	rValue.Elem().SetInt(100)		//需要通过Elem函数来获取指针指向的值,才能进行修改
}

type student struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func (stu student) ShowStuInfo() {
	fmt.Println(stu.Name,stu.Age,stu.Gender)
}

func (stu student) ShowGivenStuInfo(name string, age int, gender string) (result bool) {
	stu.Name = name
	stu.Age = age
	stu.Gender =gender
	fmt.Println(stu)
	return true
}

func showRefStruct(stu interface{}) {
	rType := reflect.TypeOf(stu) //获取type
	rVal := reflect.ValueOf(stu) //获取value,类型并非student struct
	fmt.Println(rType,rVal)		//rType打印的值:main.student,即main包下的student类型


	//因为是结构体,无法通过rVal.Xxx()转换为基本类型,因此需要先转为空接口,再进行断言
	rValStruct,ok := rVal.Interface().(student)
	if ok {
		fmt.Println(rValStruct.Name,rValStruct.Age,rValStruct.Gender)
	}

	//也可以通过reflect.Value.Kind() 获取反射的类型
	//与TypeOf不同的是,TypeOf获取的是具体范围(main.student),而kind则是广域范围(struct)
	//对于基本类型这两者的结果是相同的
	//↓的2个结果都是struct
	fmt.Println(rVal.Kind(),rType.Kind())

}


func showRefInt(i interface{}) {
	//
	rType := reflect.TypeOf(i)	//获取参数的type：int
	rVal := reflect.ValueOf(i)	//获取value，但rVal的类型是reflect.Value!，并不是int！
	rValInt := rVal.Int() + 1	//此时才是int类型，如转为其他类型触发panic
	fmt.Println(rType,rVal,rValInt)

	rValInterface := rVal.Interface() //或者先转为interface{}，再断言为需要的类型
	rValInt2 := rValInterface.(int)
	fmt.Println(rValInt2+1)

}
