package main


import (
	"fmt"
	"golang-basic-project/src/constant"            //自定义的包
	"golang-basic-project/src/part1basic/exercise" //自定义的包
	"strings"
	"time"
	"unsafe"
)

//全局变量,位置任意
var g1 int
var (
	g11 = 10
	g22 = 10
	g33 = 10
	// g33 = 1.1	//错误语句不允许更改变量类型
)

//=================================================================

//init,go框架在运行前会优先执行init函数
//执行顺序: 全局变量 > init > main		全局变量才是最先执行的
var initParam = testFunc()

func testFunc() bool {
	fmt.Print("hello")
	return true
}
func init() {
	fmt.Println("world!")
}

// main函数为入口
func main() {

	//关于变量
	//3种定义变量的方式
	//1.常规赋值
	var i1 int
	i1 = 10

	//2.自动类型推导
	var i2 = 10

	//3.:= 不同于赋值的=，推导赋值(感觉比较常用)
	i3 := 10
	fmt.Println("hello world", "aaaa!", i1, i2, i3)

	//单行赋值多个变量
	//1.同一类型
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	//2.按照顺序赋值
	var n11, n22, n33 = 10, "10", 10.00
	fmt.Println(n11, n22, n33)

	//3.	:=
	n111, n222, n333 := 10, "10", 10.00
	fmt.Println(n111, n222, n333)
	fmt.Println(g1, g11, g22, g33) //最上方的全局变量

	//经典的自增自减
	i4, i5 := 1, 1
	i4++
	i5--
	fmt.Println("i++: ", i4, i5)

	var t1 = 100.0
	fmt.Printf("%T %d\n", t1, unsafe.Sizeof(t1)) // %T 打印变量类型 %d 打印变量占用字节数

	//go没有char,通过pf %c来打印字符
	var c1 byte = 'a'
	//如果超过了assic码表 的中文,可以用int来表示
	var c2 int = '菜'
	fmt.Printf("c1 :  %c\n", c1)
	fmt.Printf("c2 :  %c\n", c2)

	//浮点,会根据分母的小数位数来决定
	fmt.Println(10.0 / 4)
	//哪怕是类型是float也不会打印小数点后的数字
	var f1 float32 = 10 / 4
	fmt.Println(f1)

	//布尔类型
	var b1 = true
	fmt.Printf("%T\n", b1)

	//字符串
	var s1 string = "hello world!"
	fmt.Println(s1)
	//或者整段输出,第一行会输出tab
	s2 := `
	abc
def
ghi
	`
	fmt.Println(s2)

	//多行拼接,+号必须接前一行
	var s3 string = "hello" +
		" world!"
	fmt.Println(s3)

	//go没有显示自动转换类型	转换语法T(v)
	var v1 int32 = 10
	var v2 int64 = int64(v1)
	fmt.Println(v2)
	fmt.Println("")

	//=================================================================
	//Pointer 指针	初始值为nil 注意：一般性为nil的传递类型都需要make！
	var i100 int = 10
	fmt.Println("int adress:", &i100)          //&var 变量的地址
	var ptr *int = &i100                       //*int 指针类型
	fmt.Println("ptr value:", ptr)             //变量的值
	fmt.Println("ptr address:", &ptr)          //变量地址
	fmt.Println("ptr address to value:", *ptr) //变量地址所指向的值
	*ptr = 20                                  //直接修改地址指向的值
	fmt.Println("int value:", i100)            //测试打印结果
	//所有基本类型前加* 便是地址
	fmt.Println("")

	//变量等首字母小写表示私有,大写表示公有
	//需要额外导入包,也需要定义Gopath,左下角 设置→go....
	fmt.Println(constant.PublicString)

	//===============================================================
	//逻辑运算符
	isTrue := true
	isOne := 1
	if isTrue && isOne == 1 {
		fmt.Println("true!")
	}

	//===============================================================
	//2个练习
	fmt.Print("不使用中间值,交换a:10,b:20两数  ")
	exercise.ExchangeAB()
	fmt.Println("")

	fmt.Print("求2/3个数中的最大值  ")
	exercise.GetMaxFromTwo()
	exercise.GetMaxFromThree()
	fmt.Println("")

	//scan的使用示范,vs中有点问题暂时注释
	// var userInput int
	// fmt.Scanln(&userInput)
	// fmt.Println(userInput + 1)

	//===============================================================
	//go中允许在if中直接定义变量
	if num := 10; num > 9 {
		fmt.Println("num > 9")

		//go 也有elif
	} else if num > 8 {
		fmt.Println("num > 8")
	}

	//===============================================================
	//switch
	v11 := 15
	switch v11 {

	//case可以有多个表达式
	case 10, 15:
		fmt.Println("v11 = 10 or 15")
		//switch穿透,匹配成功后回继续执行下一个case(本case也会执行)
		fallthrough
	case 20, 25:
		fmt.Println("v11 = 20 or 20")
	default:
		fmt.Println("v11 status unknow")
	}

	//case 也可以进行简单的逻辑判断
	v22 := 3
	switch {
	case v22 < 2:
		fmt.Println("v22 < 2")
	case v22 > 2:
		fmt.Println("v22 > 2")
	default:
		fmt.Println("v22 = 2")
	}

	//===============================================================
	//for 循环
	//index1的作用域只在for{}函数之间
	for index1 := 0; index1 < 10; index1++ {
		fmt.Print(index1, " ")
	}
	fmt.Println("")

	//java中没有的写法
	index1 := 0 //初始化外置
	//↓会导致无限循环,一般配合break
	for {
		fmt.Print(index1, " ")
		index1++
		if index1 > 9 {
			break
		}
	}
	fmt.Println("")

	//传统的字符串遍历方式,如果出现中文会乱码
	str1 := "abcdefg你大爷"
	for index := 0; index < len(str1); index++ {
		fmt.Printf("tradition for:index=%d,val=%c;\n", index, str1[index])
	}

	//go特有的字符串遍历方式,可以读取汉字!,类似于foreach
	for index, val := range str1 {
		fmt.Printf("typical for:index=%d,val=%c;\n", index, val)
	}

	//九九乘法表的练习
	exercise.MultTable()

	//golang中存在goto，一般不建议使用，没有示范

	//===============================================================
	//函数
	//测试函数，详细进testFunc中查看
	fmt.Println(exercise.TestFunc(10, 20))

	//===============================================================
	//关于包package,以src为根目录,书写时省略src/
	//假设一个package的相对路径是/src/p1/p2,那么他的import路径便是"p1/p2"
	//文件包名与文件夹名通常一致,采用小写
	//package main 只有一个

	//===============================================================
	//关于函数
	exercise.RecursiveFunc(10)                   //递归
	fmt.Println(exercise.Fibonacci(10))          //斐波那契 for 实现
	fmt.Println(exercise.RecursiveFibonacci(10)) //斐波那契 递归实现

	fmt.Println(exercise.CallParamFunc(exercise.SumFunc, 1, 2)) //当函数作为参数传递的例子

	//type 自定义数据类型 相当于取别名
	type myInt int //自定义了一个int类型的myInt
	var i myInt = 1
	//var j int = 1		//此2行会报错，即使myInt和int都是同一类型，但编译器不识
	//i = j
	fmt.Println(i)
	fmt.Println(exercise.CallParamFuncByMyFunc(exercise.SumFunc, 1, 2)) //此函数的第一参数为自定义函数类型

	//直接定义返回值名的时候
	fmt.Println(exercise.FuncHasReturnName(1, 2))

	//多个返回结果，可以使用_忽略结果,如果两个都忽略,你为什么要去调用函数呢
	a, _ := exercise.TwoReturnResult()
	fmt.Println(a)

	//可变参数	sum所有参数
	fmt.Println(exercise.CallVariableParam(1, 2, 3, -10))

	//匿名函数
	//1.直接赋值,只调用一次
	lv1 := func(la1 int, lb1 int) int {
		return la1 + lb1
	}(1, 2)
	fmt.Println(lv1)

	//2.赋值给变量,通过变量来调用匿名函数,主要用于闭包
	lv2 := func(la2 int, lb2 int) int {
		return la2 + lb2
	}
	fmt.Println(lv2(1, 2))
	//3.也可以写成全局函数,略

	//闭包	一个函数与其相关的引用环境,组成一个整体(实体)
	f := exercise.ClosureFunc()
	fmt.Println(f(1), " ", f(2), " ", f(3))

	//defer 关键词 延迟加载。函数执行完后可【快速释放资源】，多用于连接、文件打开等，可直接跟一句defer语句
	fmt.Println(exercise.CallDeferFunc(1, 2))

	//值类型：基本数据，int，float，数组，结构体等
	//引用类型：指针、切片、map集合、管道chan、接口interface
	//传递地址其实是引用传递，效率高，因为占用数据量少

	//===============================================================
	//字符串相关函数介绍
	//以rune进行字符串遍历,可用于处理assic以外的字符（不会乱码）
	str2 := "你你好好"
	r1 := []rune(str2)
	for index := 0; index < len(r1); index++ {
		fmt.Printf("%c  ", r1[index])
	}
	fmt.Println("")

	//EqualFold 字符串的比较
	fmt.Println(strings.EqualFold("abc", "ABC")) //不区分大小写，true
	fmt.Println("abc" == "ABC")                  //区分大小写，false

	//Contains 子串（第二个参数）是否包含
	fmt.Println(strings.Contains("abcdefg", "a")) //true

	//Count 统计子串出现次数
	fmt.Println(strings.Count("abcccc", "c")) //4

	//字串出第一次出现的位置
	fmt.Println(strings.Index("abcdefg", "efg")) //4

	//Split 分割
	fmt.Println(strings.Split("a,b,c,d,e", ",")[0])

	//ToUpper ToLower 大小写转换
	fmt.Println(strings.ToUpper("abcdefg"))
	fmt.Println(strings.ToLower("ABCDEFG"))

	//Trim 去两边的特定字符，TrimLeft/TrimRight 去左/右的特定字符 TrimSpace 去空格
	fmt.Println(strings.Trim("?abcd?", "?"))

	//HasPrefix/HasSuffix 是否以子串为开头/结尾(区分大小写)
	fmt.Println(strings.HasPrefix("abcdefg", "A"))
	fmt.Println(strings.HasSuffix("abcdefg", "a"))

	//===============================================================
	//date相关函数
	now := time.Now()
	//个人觉得下方为推荐写法
	fmt.Printf("%d/%d/%d %d:%d:%d \n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	//或者使用↓这个时间	之后可以按照此时间的格式来打印当前时间	比较奇葩
	now.Format("2006/01/02 15:04:05")
	fmt.Println(now.Format("2006/01/02")) //打印当前时间

	//获取时间的常量	time.Millisecond
	start := time.Now().UnixNano() / 1e6 //Unix 秒，UnixNano 纳秒，go没有毫秒，需要手动转换
	time.Sleep(1000 * time.Millisecond)  //休眠1000毫秒
	end := time.Now().UnixNano() / 1e6
	fmt.Printf("总计等待 %v ms\n", (end - start))

	//unix时间
	fmt.Println(now.Unix())

	//===============================================================
	//一些内置函数
	//new （一般）分配值类型的内存，内存，内存！
	num1 := new(int) //与num1 := 10相比，做了更多的事
	*num1 = 10       //获取num1指向的值，并赋值
	fmt.Printf("num1 type=%T,value=%v,address=%v,point value=%v\n", num1, num1, &num1, *num1)

	//make一般分配引用类型的内存（跳过了）

	//===============================================================
	//异常处理
	//通过defer + recover实现异常捕获
	exercise.CallDRException()

	//自定义异常的实现
	exercise.CallPanicFunc("ok.jpg") //成功的情况

	result2 := exercise.CallPanicFunc("not_ok.jpg") //失败的情况，可以打印result来输出error日志
	if result2 != nil {
		fmt.Println(result2)
	}

	//===============================================================
	//数组	值类型，因为地址直接指向数据空间	新开辟的数组默认值会赋予0/""/false
	//数组的地址即第一个下标的地址，次个的地址便是当前地址单个数组的长度
	//例如int的地址是8，第一个在0xc04204a0a0，第二个便是0xc04204a0a8
	//数组声明后长度会固定无法变化
	//数组可以存放任意类型元素，但同一数组内类型必须相同
	//补充：var arr []int 不是数组，是切片
	var arrayInt [3]int
	for index := 1; index < len(arrayInt); index++ {
		arrayInt[index] = index
	}
	fmt.Println(arrayInt)

	//对数组初始化的三种方式
	var arrayInt1 [3]int = [3]int{1, 2, 3}
	var arrayInt2 = [3]int{1, 2, 3}
	//有点特殊：下标1的赋值1，2赋值2，3没有定义下标但会默认下标为3，下标0没有定义默认为0
	var arrayInt3 = [...]int{1: 1, 2: 2, 3}
	arrayStr1 := [...]string{1: "a", 2: "b", "c"} //让编译器自动推导为str数组，0标为""
	fmt.Println(arrayInt1, arrayInt2, arrayInt3, arrayStr1)

	//for-range遍历数组
	for index, v := range arrayStr1 {
		fmt.Println(index, " ", v)
	}

	//可以通过引用传递来该表数组元素
	fmt.Println(arrayStr1)
	arrStrAdd := &arrayStr1 //获取arrStr的指针地址
	(*arrStrAdd)[0] = "0"   //指针的0标赋值"0"
	//arrayStr1[0] = "0"	//其实这种办法也可以。。。
	fmt.Println(arrayStr1)

	//假设一个func 的参数是([]int),但传入的参数是([...]int{1,2,3}),会出现编译报错,因为不是统一类型
	//相同的,如果原参数为[4],传入[3]也会报错!

	//练习,通过for-range获取数组中的max
	arrInt4 := [5]int{1, 2, 3, 4, 5}
	exercise.GetMaxFromArr(arrInt4)

	//练习,长度5的数组内生成随机数,然后反转
	exercise.GetRandomNumAndReverse()

	//===============================================================
	//切片slice		引用类型 类似与数组,但是可变,也即动态数组	切片的底层其实是结构体struct	类似于JavaBean..
	//语法	var s []int
	//1.第一种创建切片的方法,从数组中切出
	arrInt5 := [7]int{1, 2, 3, 4, 5, 6, 7}                                         //先定义一个数组
	slice1 := arrInt5[2:5]                                                         //切出[1~3)的切片,即第二,第三个元素
	fmt.Println("slice1: ", slice1, "容量capacity :", cap(slice1))                   //关于此处的cap:切原来数组的一刀后切出来的长度等于数组长度 - 从哪里切,即7 - 2 =5
	fmt.Println("slice1 address: ", &slice1[0], "arrint5[1] address", &arrInt5[2]) //切出去的第一个地址=原数组该位置的地址,因此切片才是[引用类型]

	//2.通过make创建切片	该种类的切片(数组)是由make维护,只能通过切片访问各元素(无法通过数组访问),这也是1和2创建数组的不同点
	//var slice2 []int	//此种方法创建的切片是空的,无法使用!只要需要让其引用到一个数组,或者通过make开辟空间
	var slice2 []int = make([]int, 5, 10) //make内的参数含义:(类型,长度,容量)
	fmt.Println("slice2 :", slice2)
	fmt.Println("slice2 len:", len(slice2))
	fmt.Println("slice2 cap:", cap(slice2))

	//3.直接创建,类似于make
	slice3 := []int{5, 4, 3, 2, 1}
	fmt.Println("slice3: ", slice3)

	//遍历打印slice
	fmt.Print("slice3 by for-range: ")
	for _, s := range slice3 { //var关键词不能用..
		fmt.Print(s, " ")
	}
	fmt.Println("")

	//slice的简写
	//[1:] 从1切到尾 [:5]从头切到5 [:] 全切
	slice4 := arrInt5[:]
	fmt.Println("slice4 : ", slice4)

	//切片后还可以再次切片(允许套娃)

	//关于append 本质是数组扩容
	//go底层会创建新的数组,再将原来的元素拷贝至新的数组,再使slice引用新得切片
	//append之后并不会覆盖原来的切片
	slice5 := append(slice4, 10, 9, 8)
	// slice5 := append(slice4, slice4...) //也可以直接追加切片,  ... 是固定写法!
	fmt.Println("slice5 : ", slice5)

	//copy切片	(拷贝至，源文件)	即使源文件元素 > 拷贝后的切片也不会报错,只会拷贝 拷贝文件大小的元素
	// copy(slice5, slice1)

	//注意一点,slice是引用类型,当slice[0] = 10之后,原来所切的数组的对应元素也会发生改变!

	//string的切片
	str3 := "abcdefghijklmn"
	slice6 := str3[:5]
	fmt.Println("slice6 : ", slice6)
	//补充:和java一样 string是不可变类型	修改string的方法:assic字符直接用byte,汉字等可以用rune
	arrStr := []byte(str3) //转为数组
	arrStr[0] = 'z'
	runeStr := []rune(str3) //转为rune
	runeStr[0] = '阿'

	newStr := string(arrStr) //再转回string,下同
	newStrHZ := string(runeStr)
	fmt.Println("new string : ", newStr)
	fmt.Println("new string with hanzi : ", newStrHZ)

	//求n个斐波那契数列,并存入切片,并返回切片
	fmt.Println("Fibonacci slice :", exercise.CallFibonacciSavedBySlice(1))

	//轴x,y   半径r	//半径r的圆最多可以放多少个边长a的正方形,结果不太对
	// count := 0
	// r := 5
	// for y := 0; y <= r; y++ {
	// 	maxX := math.Sqrt(float64(r - y + 1))
	// 	for x := 0; x <= r; x++ {
	// 		if float64(x) <= maxX {
	// 			count++
	// 		}
	// 	}
	// }
	// fmt.Println(count * 4)

	//冒泡排序
	slice7 := []int{7, 6, 5, 4, 3, 2, 1}
	fmt.Println("bubble sotr: ", exercise.BubbleSort(slice7))

	//===============================================================
	//golang中的二维数组
	var twoDimensionArray [3][3]int = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("two dimension arr: ", twoDimensionArray)

	//二维数组的遍历
	for index1 := 0; index1 < len(twoDimensionArray); index1++ {
		for index2 := 0; index2 < len(twoDimensionArray[i]); index2++ {
			fmt.Print(twoDimensionArray[index1][index2], " ")
		}
	}
	fmt.Println("")

	//for-range的方式(更简洁)
	for _, v1 := range twoDimensionArray {
		for _, v2 := range v1 {
			fmt.Print(v2, " ")
		}
	}
	fmt.Println("")

	//===============================================================
	//map 映射/集合	引用类型
	//语法:var mapname map[keyType]valueType
	//不能用==的不能作key	有序无序根据key的类型
	var map1 map[int]string //此时的map 为 nil	;另 map 关键词不能用做name
	//使用map前需要make分配空间,初始空间大小可自定义,也可默认
	map1 = make(map[int]string, 10)
	map1[2] = "second element"
	map1[1] = "first element"
	fmt.Println(map1)

	//比较轻松的写法
	map2 := make(map[int]string)
	map2[10] = "test"
	fmt.Println(map2)

	//直接定义map的写法
	map3 := map[int]string{
		11: "test", //每一行结束都必须加 ,
		12: "test", //最后一行也是
	}
	fmt.Println(map3)

	//map的元素删除	即使key=nil也不会报错
	delete(map3, 11)

	//查找key
	//根据key查找,会返回2个结果,value的值,是否存在bool
	val, flag := map3[12]
	if flag {
		fmt.Println("this value is exist:", val)
	}

	//map的遍历
	for key, value := range map1 {
		fmt.Println(len(map3)) //len也可以打印map的长度
		fmt.Println(key, ":", value)
		//关于顺序：此处打印的结果是先2再1，因为存储的时候也是先2再1，然后直接打印map1则是先1再2
	}

	//map slice map切片		与传统切片类似，但是元素是map，他的本质还是切片，通过slice[0]来获取其中的map元素
	var mapSlice1 []map[string]string        //定义一个map切片
	mapSlice1 = make([]map[string]string, 2) //通过make，创建2个map slice空间

	mapSlice1[0] = make(map[string]string, 2) //新创建的map slice依然是nil！ 必须在此位置再次make创建空间
	mapSlice1[0]["key1"] = "abc"              //为0位置的map填入map键值对
	mapSlice1[0]["key2"] = "def"
	fmt.Println("map slice 1 : ", mapSlice1) //此时1位置依然是nil

	newMap := map[string]string{ //新建一个map
		"KEY1": "ABC",
		"KEY2": "DEF",
	}
	mapSlice1 = append(mapSlice1, newMap) //通过append来添加新的map元素
	fmt.Println("new map slice 1 : ", mapSlice1)
	//打印后的结果，包含3个map集合的map切片：[map[key1:abc key2:def] 		map[] 		map[KEY1:ABC KEY2:DEF]]

	//map的排序	注意！！！可能是版本问题，新版本的fp已经是顺序打印，但for-range还是乱序
	//通过for-range打印的map必然会是乱序，下面是证明
	map4 := make(map[int]int, 10)
	map4[3] = 3
	map4[2] = 2
	map4[5] = 5
	map4[1] = 1
	map4[4] = 4
	map4[7] = 7
	map4[6] = 6
	fmt.Println(map4)
	for _, value := range map4 {
		fmt.Print(value, "  ")
	}
	fmt.Println("")
	//多次执行的结果都不相同

	//想要map有序需要切进map切片中，然后通过sort函数进行排序
	// var sortSlice []int
	// for key := range map4 {
	// 	sortSlice = append(sortSlice, key)
	// }
	// sort.Ints(sortSlice)
	// fmt.Println(sortSlice)
	//虽然这里for-range还是乱序。。
	// for _, value := range map4 {
	// fmt.Print(value,"  ")
	// }

	//map的一些注意点
	//map是引用传递 别的func如果修改了map会导致原map的值也发生变换
	//（注意，go函数传参，传递方式都是值传递，在传递map的时候传递的值其实是地址）
	//map容量到达上限后会自动扩容，并不同于切片需要append
	//map的键值对中，值经常使用结构体struct

}
