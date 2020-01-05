package exercise

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// TestFunc 测试函数
// 语法 func 函数名(参数1,参数2...) (返回值类型1,2...) {...}
func TestFunc(a int, b int) (bool, int) {
	result := a * b
	if result > 0 {
		return true, result
	}
	return false, 0
}

//RecursiveFunc 递归测试
func RecursiveFunc(num int) int {
	num--
	fmt.Println(num)
	if num > 0 {
		RecursiveFunc(num)
	}
	return num
}

//Fibonacci 公式 F(n)=F(n-1)+F(n-2)
func Fibonacci(n int) int {
	num1 := 1
	num2 := 1
	result := 0
	for n-2 > 0 {
		result = num1 + num2
		num1 = num2
		num2 = result
		n--
	}
	return result
}

//RecursiveFibonacci 使用递归
func RecursiveFibonacci(n int) int {
	if n > 2 {
		return RecursiveFibonacci(n-1) + RecursiveFibonacci(n-2)
	}
	return 1
}

//SumFunc let 2 nums sum
func SumFunc(num1 int, num2 int) int {
	return num1 + num2
}

//CallParamFunc 调用参数函数
func CallParamFunc(func1 func(int, int) int, num1 int, num2 int) int {
	return func1(num1, num2)
}

//自定义函数 muFunc
type myFunc func(int, int) int

//CallParamFuncByMyFunc 传入自定义类型的参数
func CallParamFuncByMyFunc(func1 myFunc, num1 int, num2 int) int {
	return func1(num1, num2)
}

//FuncHasReturnName 直接定义返回值名的时候,func的body中定义参数后可以省略return的内容
func FuncHasReturnName(a int, b int) (sum int) {
	sum = a + b
	return //连返回值都不用写了。。
}

//TwoReturnResult 测试多返回结果
func TwoReturnResult() (r1 int, r2 int) {
	r1 = 1
	r2 = 2
	return r1, r2
}

//CallVariableParam 可变参数 args ... 必须三个点
// func CallVariableParam(args ...int) {		// 0~n个
func CallVariableParam(n1 int, args ...int) (sum int) { //1~n个
	sum = n1
	for index := 0; index < len(args); index++ {
		sum += args[index]
	}
	return
}

//ClosureFunc 闭包自增,返回匿名函数
//n只会初始化一次,第二次调用后会在原有的n值上再进行n += a
//一半用于存储1.存储变量；2.封装私有变量，外部无法引用
//概念：返回的函数与n形成了闭包
func ClosureFunc() func(int) int {
	n := 100
	return func(a int) int {
		n += a
		return n
	}
}

//CallDeferFunc 延迟加载
//延迟加载会进行压栈，后进先出，所以会先打印第二个数字
func CallDeferFunc(num1 int, num2 int) (result int) {
	defer fmt.Println("defer num1 :", num1)
	defer fmt.Println("defer num2 :", num2)
	//defer语句入栈后，会将相关值先拷贝入栈，一般用于文件加载，↓案例
	//file = openfile(filename)
	//defer file.close()	//优先拷贝入栈，延迟关闭程序执行完毕后自动释放，多用于数据库连接，文件打开等
	num1++
	num2++
	result = num1 + num2
	return
}

//CallDRException 测试异常捕获，通过defer + recover实现
func CallDRException() {
	defer func() { //直接调用延迟匿名来加载异常捕获
		err := recover() //recover 内置函数，捕获异常
		if err != nil {  //说明捕获到异常
			fmt.Println(err)
		}
	}()
	num11 := 10
	num22 := 0
	res1 := num11 / num22
	fmt.Println("res=", res1)
}

//CallPanicFunc 测试panic，自定义异常
func CallPanicFunc(filename string) (err error) {
	if filename == "ok.jpg" {
		fmt.Println("filename is ok.jpg")
		return nil
	}
	//如果报错则新起异常
	return errors.New("filename error")
}

//GetRandomNumAndReverse 定义一个长度5的数组,生成5个[0~100)的随机数,反转结果
func GetRandomNumAndReverse() {
	arrInt := [5]int{}
	rand.Seed(time.Now().UnixNano()) //go的rand不改seed则结果每次都一样..
	for index := 0; index < 5; index++ {
		arrInt[index] = rand.Intn(100) //[0~100)
	}
	fmt.Println("arrInt origin:", arrInt)
	temp := 0
	for index := 0; index < len(arrInt)/2; index++ {
		temp = arrInt[index]
		arrInt[index] = arrInt[4-index]
		arrInt[4-index] = temp
	}
	fmt.Println("arrInt reverse:", arrInt)

}

//CallFibonacciSavedBySlice 通过切片储存斐波那契,并返回结果
func CallFibonacciSavedBySlice(n int) []uint64 {
	fbSlice := make([]uint64, n)
	fbSlice[0] = 1
	if n == 1 {
		return fbSlice
	}
	fbSlice[1] = 1
	if n == 2 {
		return fbSlice
	}
	//通过遍历存入切片中,注意index 的起始应该是2!
	for index := 2; index < n; index++ {
		fbSlice[index] = fbSlice[index-1] + fbSlice[index-2]
	}
	return fbSlice
}
