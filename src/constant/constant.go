package constant

// PublicString 常量字符串
//const 定义常量
//const与变量var的区别:
//1.常量定义时必须赋值,定义后无法修改
//2.只能修饰bool,数值(int,float类)和字符串
const PublicString string = "HELLO WORLD"
const privateString string = "hello world"	//golang的常量值也可以定义为私有(废话)

//多个常量
const (
	MyInt = 1		//常规赋值方式
	MyFloat = 2.0
	MyIOTA = iota	//此时的iota是[2]位置,因此值为2
)
const (
	MyIOTA0 = iota	//专业写法,iota;赋予初始变量0 iota: 微
	MyIOTA1			//位置[1],MyIOTA1 = 1
	MyIOTA2			//2
	MyIOTA3 = iota  //3
	MyIOTA4,MyIOTA5 = iota,iota //位置都是[4],因此4,4

)


//	fmt.Println(constant.MyIOTA,constant.MyIOTA0,constant.MyIOTA1,constant.MyIOTA2,constant.MyIOTA3,constant.MyIOTA4,constant.MyIOTA5)