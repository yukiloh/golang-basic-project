package main //go的打包方式。。。

import (
	"encoding/json"
	"fmt"
	"golang-basic-project/src/part2oop-struck/model"
)

//定义结构体
type student struct {
	name   string
	age    int
	gender string //F or M
}

//定义结构体方法
func (s student) test(str string) {
	fmt.Println("用于测试方法的传参:", str)
	s.name = "莫的名字"                         //如果直接在stu的方法内定义name！
	fmt.Println("this is student's method") //此处的s，是传入的student类型的对象s
	fmt.Println("在本方法内的名字: ", s.name)       //此处的s，是传入的student类型的对象s
}

//方法的传递指针
func (s *student) test2() {
	//通过指针修改地址指向的值,此处是简写
	s.age = 14
	//s已经是指针,需要通过(*s)获取指针指向的地址,此为标准写法
	//fmt.Println("用于测试指针:", (*s).name, (*s).age, (*s).gender)
	//也可以用语法糖来简写
	fmt.Println("用于测试指针,通过指针修改了年龄:", s.name, s.age, s.gender)

}

//演示String()	只有传入指针才会触发!
func (s *student) String() (str string) {
	fmt.Println("toString done!")
	str = fmt.Sprintf("name: %v,age: %v,gender: %v", s.name, s.age, s.gender)
	return
}

func main() {

	//===============================================================
	//go的面向对象
	//golang是支持面向对象，而并非他就是面向对象（oop）
	//面向对象的语言都有class，golang中则是struct 结构体
	//golang的面向对象编程非常“简洁”，没有传统的继承、重载、构造函数、析构函数（扫尾工作），也没有this指针
	//golang仍保留继承、封装、多态，实现方式和传统不一样，例如继承通过匿名字段实现
	//golang通过 接口关联，来降低耦合，golang中【面向接口编程】是非常重要的一点

	//===============================================================
	//结构体 struct	值类型 引用以后不会影响原来的结构体！可以通过new创建

	//定义结构体的方式	结构体名和字段名都可以设为大写，即public

	//struct的创建方式1：
	var stuA student
	stuA.name = "二狗"
	stuA.age = 11
	stuA.gender = "M"

	fmt.Println("struct: ", stuA, " name:", stuA.name)

	//方式2	推导
	stuB := student{"狗蛋", 12, "M"}
	fmt.Println("struct: ", stuB)

	//或者用推导,定义字段值,注意此种写法不受字段顺序约束
	stuC := student{
		name:   "狗蛋的兄弟",
		age:    12,
		gender: "M", //逗号必须加
	}
	fmt.Println("stuC: ", stuC)

	//方式3 通过指针类型创建
	var ptn *student = new(student)
	//标准写法
	(*ptn).name = "大柱"
	//简化写法,go底层会自动添加取值运算符	为啥有这种写法?因为符合开发者的习惯
	ptn.age = 13
	ptn.gender = "M"
	fmt.Println("struct: ", *ptn)
	//复习一遍指针
	//ptn = *obj		这是指针类型
	//&ptn 				指针指向的值(地址)
	//*ptn				指针指向的值(地址)的值

	//方式4 取对象的地址
	var ptn2 *student = &student{}
	(*ptn2).name = "翠花" //和↑相同,标准写法
	ptn2.age = 12       //简化
	ptn2.gender = "F"
	fmt.Println("struct: ", *ptn2)

	//结构体在内存中的地址是连续的!
	//2个字段完全相同(字段名,类型,字段个数)的结构体不能直接赋值,但可以强转

	//关于struct的tag	多用于序列化
	type Student2 struct {
		Name   string `json:"name"` //添加tag
		Age    int    `json:"age"`
		Gender string `json:"gender"` //F or M
	}

	var stu2A Student2
	stu2A.Name = "铁柱"
	stu2A.Age = 20
	stu2A.Gender = "M"

	//直接打印时因为是public，所以必定key会变成大写，可以通过添加tag的方式来转为小写
	json1, err := json.Marshal(stu2A) //json工具函数，会返回json和error信息		marshal：编排，marshal底层能读取到tag是利用了反射
	if err != nil {
		fmt.Println("err: ", err)
	} else {
		fmt.Println("json1: ", string(json1))
	}

	//===============================================================
	//golang的方法  只属于自己结构体类型的方法
	//定义二狗
	var stu2dog student = student{"二狗", 11, "M"}
	//让二狗同学调用自己的方法
	stu2dog.test("hello world")        //此时会在方法内部调用给name赋值的语句,因此会打印"莫得名字"
	fmt.Println("对象真名：", stu2dog.name) //调用方法时,将对象复制给方法,因此方法内重新定义name不会影响原来的name,还是二狗,除非传递了地址

	//golang的方法在执行时,会把变量本身也调用(复制一份至方法中),这也是与普通的函数不同的地方

	//golang的方法一般通过值拷贝传递,也可以修改通过指针方式传递,传递地址更高效
	var stu3dog student = student{"三狗", 13, "M"}
	//test2传递的是指针,因此此处需要获取stu的地址
	fmt.Println("三狗原来年龄:", stu3dog.age)
	(&stu3dog).test2() //会发现此时因为传递的是指针,三狗的年龄发生了改变!

	//方法不仅仅可以作用在自定义类型,也可以作用在int,float等上
	//(但需要先转成自定义,例如type integer int,原生的不可 )

	//关于String(),当定义了String(),且打印对象(需要根据String中定义的类型,对象/指针)时
	//fp时会默认调用String()方法,类似于toString()
	fmt.Println("String() : ", &stu3dog)

	//方法和函数的区别
	//1.调用方式不同；函数:函数名(参数),方法:变量.方法名(参数)	我觉得是差不多的..
	//2.普通的函数,接收者是值类型时,只能接受值(不能是指针)
	//  而方法,可以用指针来调用类型接收者(基于底层的优化), ↓例子
	stu4dog := student{"4狗", 14, "M"}
	//此处的&是语法糖;test方法接收者并非指针,但go底层会优化,使其获取到值进行传递
	(&stu4dog).test("hello")
	//↑ 如果接收者是指针,而传参是类型,也可通过语法糖直接传p而不是(&p),所以接收者类型是关键!
	fmt.Println("4狗的真名:", stu4dog.name)
	//谨记,go的struct是值传递!!!
	//函数也是值传递,哪怕是传递指针,函数会复制指针地址

	//面向对象编程的大概过程(和java类似)
	//确定对象(struct),定义方法,调用

	//===============================================================
	//结构体的工厂模式(goland的alt + insert)
	// stu1 := model.Student{"tom", 11}	// 当没有定义键名时会报警告(可运行):	composite literal uses unkeyed fields
	//↓是传统获取结构体字段的方式
	stu1 := model.Student{
		Name:  "tom",
		Score: 11,
	}
	fmt.Println("stu1: ", stu1)

	//一旦Student是私有便无法访问,因此需要一个工厂模式来存/取值
	stu2 := model.NewStudent("jerry", 10) //get方法
	fmt.Println("stu2: ", *stu2)
	fmt.Println("stu2 name:", stu2.GetName()) //set

	//golang中的封装是简化的，并非像java中强调需要进行封装

	//===============================================================
	//golang的继承 通过匿名结构体实现
	apple := &model.Fruit{Goods: model.Goods{Name: "apple", Price: 5}, Quantity: 2} //通过定义匿名结构体来定义Goods	另外idea推荐创建结构体时指定名字！
	//apple.Goods.ShowGoods()	//此行与下行都是一个意思，可以简写为↓
	apple.ShowGoods()
	apple.ShowFruit()

	apple.Buy(1)      //买了一个
	apple.ShowFruit() //再次显示现在apple有几个
	//补充：如果结构体与匿名结构体都存在相同字段，则会根据就近访问原则，优先取结构体中的字段，需要访问匿名结构体中的字段可以加上匿名结构体名
	//例如：apple.name		apple.goods.name
	//如果结构体中存在2个匿名结构体，且都有相同的字段，则必须指明是引用哪个匿名结构体
	//例如：	apple{A{name:"a"},B{name:"b"}}	此时无法通过apple.name来访问,必须添加匿名结构体名!

	//有名结构体	与匿名结构体相对,无法通过省略结构体名来访问
	carrot := &model.Greens{Goods: model.Goods{Name: "carrot", Price: 6}, Quantity: 2}
	fmt.Println(carrot.Goods.Name)
	//fmt.Println(carrot.Name)	//访问不到的

	//golang中多重继承也就是嵌入多个结构体.一般为了代码简洁性,不推荐经常使用多重继承

}
