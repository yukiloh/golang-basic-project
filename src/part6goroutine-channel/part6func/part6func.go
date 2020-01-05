package part6func

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//每个1s打印一行hello world
func PrintHelloWorldWithOneSecondInterval() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hello world", i)
		time.Sleep(time.Second) //代表1秒
	}
}

//每2s打印
func PrintHelloWorldWithTwoSecondInterval() {
	for i := 1; i <= 5; i++ {
		fmt.Println("hello world", i)
		time.Sleep(time.Second * 2)
	}
}

func ShowCPUInfo() {
	cpu := runtime.NumCPU() //查看cpu逻辑核数
	runtime.GOMAXPROCS(cpu) //设置同时使用最大核数,1.8后默认设置为最大
	fmt.Println(cpu)
}

//定义全局变量map
var (
	LockedGlobalMap = make(map[int]int, 10)
	lock            sync.Mutex //mutex:互斥锁
)

//计算1~200的阶乘,使用锁对象
func CallGoroutineByLock(n int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	lock.Lock() //用之前加锁
	LockedGlobalMap[n] = result
	lock.Unlock() //用完解锁

}

//基本chan的使用
func TestBasicChannel() {
	intChan := make(chan int, 3)   //创建一个chan
	fmt.Println(intChan, &intChan) //channel本身也是引用类型,打印结果0xc000082000 0xc000006028

	intChan <- 1 //向chan写入数据
	num1 := 2
	intChan <- num1 //变量也可

	// chan与slice不同,不会自动扩容,设定容量为3只可塞入3个数据
	num2 := <-intChan //取数据时,取一个少一个,不影响容量;全部取出后无法再进行取值,会报错 deadlock
	//<-intChan				//也可以只进行取值,直接丢弃该值,不会报错
	fmt.Println("num2:", num2)
	fmt.Println("cap:", cap(intChan), "len:", len(intChan))
}

type s struct {
	key   string
	value int
}

//需要传入任意类型的时候
func TestAnyTypeDataChannel() {
	var allTypeChan chan interface{} //任何数据类型默认实现空接口
	allTypeChan = make(chan interface{}, 10)
	i := 1            //定义一个int
	j := "string"     //定义一个string
	k := s{"key", 22} //定义一个struct

	//把数据塞进去
	allTypeChan <- i
	allTypeChan <- j
	allTypeChan <- k

	//取出数据
	i1 := <-allTypeChan
	j1 := <-allTypeChan
	k1 := <-allTypeChan

	//因为存入的是空接口类型,此时根本无法取所传入的结构体字段!
	//k1.key	//编译报错
	//因此需要使用断言
	k2 := k1.(s)
	fmt.Println(k2.key, k2.value)

	fmt.Println(i1, j1, k1)

	//管道可以关闭,关闭后无法再写入但可以读取; 实际类似于在尾部添加了一个标识EOF,无法继续添加
	close(allTypeChan) //内建函数close
}

//遍历管道
func TestTraverseChannel() {
	//先创建一个管道,添加数据
	allTypeChan := make(chan int, 20)
	for i := 0; i < 20; i++ {
		allTypeChan <- i
	}

	//应该避免使用常规的for去遍历,因为协程情况下无法确定管道的真实长度
	//通常应使用for-range遍历
	//注意,因为管道关闭后会给尾部添加一个标识,for-r遍历会读出这个标识从而终止遍历
	//如果没有close,for-r会认为还会有数据写入,会进行等待,造成死锁
	close(allTypeChan) //先关闭管道
	for i2 := range allTypeChan {
		fmt.Println("value: ", i2)
	}

}

//练习:使用2个协程,分别来读写同一管道
func ChannelExercise() {

	var wg sync.WaitGroup //定义waitGroup来维持主线程
	wg.Add(2)             //计步器，为0时主线程结束，-1 panic

	intChan := make(chan int, 50) //用于读写的管道

	//有waitGroup了之后这个就省了
	//flagChan := make(chan bool, 1) //用于维持主进程的管道

	total := 500 //总共多少个数据

	go writeData(total, intChan, &wg) //此处需要传入wg的地址
	go readData(intChan /*flagChan,*/, &wg)

	wg.Wait()
	fmt.Println("done!")

	//有waitGroup了之后这个就省了
	////维持线程的办法
	//letThreadFin(flagChan)

}

//写入的协程
func writeData(total int, intChan chan int, wg *sync.WaitGroup) {
	//放入total个数据
	for i := 0; i < total; i++ {
		intChan <- i
		fmt.Println("put data : ", i)
	}
	wg.Done()      //done后wg的计步器-1
	close(intChan) //记得关闭
}

//读取
//补充：如果对于管道只有写入没有读取，超过容量后会发生死锁
//但有读有写时不会，底层会对管道优化，无论读写哪个过快都不会发生死锁
func readData(intChan chan int /* flagChan chan bool,*/, wg *sync.WaitGroup) {
	for {
		v, ok := <-intChan //读取管道,!ok表示读到结束标识,用于结束循环
		if !ok {
			break
		}
		fmt.Println("get data: ", v)
	}

	wg.Done()

	//有waitGroup了之后这个就省了
	////完成遍历后代表任务结束,修改flagChan告诉主进程可以结束线程
	//flagChan <- true
	//close(flagChan) //关闭添加结束标识符
}

//有waitGroup了之后这个就没用了
//通过循环读取flagChan来维持主线程
func letThreadFin(flagChan chan bool) {
	for {
		_, flag := <-flagChan
		if !flag {
			break
		}
	}
}

//练习:
//启动一个协程,创建数据n(int)
//启动8个协程,读取数据n,计算阶加,存入结果集chan
//最后使8个协程遍历resChan,打印结果

//练习：
//通过协程求1~n之间的素数，打印结果。1条存放素数协程，4条取+处理，1条打印结果
func ShowPrime() {
	var wg sync.WaitGroup //定义waitGroup
	wg.Add(6)             //定义协程计步器，1条存放，4条取出，1条打印结果

	intChan := make(chan int, 1000)   //用于存放prime的管道
	primeChan := make(chan int, 8000) //用于取出prime的管道
	flag := 0                         //用于协程完成后+1
	maxGoroutine := 4                 //协程线数
	n := 80000                        //n的值

	//putPrime协程
	go putPrime(intChan, &wg, n)

	//getPrime协程
	for i := 0; i < maxGoroutine; i++ {
		go getPrime(intChan, primeChan, &wg, &flag, maxGoroutine)
	}

	//展示结果
	go showResult(primeChan, &wg)
	wg.Wait()
	fmt.Println("all goroutine done!")
}

func putPrime(intChan chan int, wg *sync.WaitGroup, total int) {
	for i := 1; i <= total; i++ {
		intChan <- i
	}
	wg.Done()
	close(intChan)
}

func getPrime(intChan chan int,
	primeChan chan int,
	wg *sync.WaitGroup,
	flag *int,
	maxGoroutine int) {
	for {
		v, ok := <-intChan
		if !ok {
			*flag++                    //当intChan取完后flag++
			if *flag >= maxGoroutine { //当flag大等于最高协程数后关闭管道
				close(primeChan)
			}
			break
		}
		if isPrime(v) {
			primeChan <- v
		}

	}
	wg.Done()
}

func isPrime(v int) bool {
	if v == 1 || v == 2 || v == 3 {
		return true
	}
	for i := 2; i < v; i++ {
		if v%i == 0 {
			return false
		}
	}
	return true
}

func showResult(primeChan chan int, wg *sync.WaitGroup) {
	for {
		v, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println("prime: ", v)
	}
	wg.Done()
}

func ShowPrimeByNormal() {
	//异常捕获 快捷输入deferr
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	n := 80000 //n的值

	for i := 0; i < n; i++ {
		if isPrime(i) {
			fmt.Println("prime: ", i)
		}
	}

	fmt.Println("all goroutine done!")
}
