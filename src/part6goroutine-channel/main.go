package main

import (
	"fmt"
	_ "fmt"
	"golang-basic-project/src/part6goroutine-channel/part6func"
	"math/rand"
	"time"
)

//多线程 goroutine和channel
func main() {

	//协程goroutine 轻量级线程
	//go存在主线程,而主线程是由多个协程构成(由编译优化)

	//go中协程的特色:
	//1.有独立栈空间
	//2.共享程序的堆空间
	//3.由用户调度
	//4.轻量级的线程

	//简单协程演示
	//插入协程 如果主线程(main)退出,即使协程没有执行完毕也会立刻退出
	go part6func.PrintHelloWorldWithTwoSecondInterval()
	//使用主线程,每个1s打印一行hello world
	part6func.PrintHelloWorldWithOneSecondInterval()

	//主线程属于物理线程,直接作用在cpu上,重量级
	//而协程属于轻量级的,逻辑态,消耗资源相对较少
	//golang中可以轻松开启上万的协程;其他语言则基于线程,对资源消耗较大
	//因此golang天生存在并发优势

	//goroutine的调度模式 MPG模式	M:主线程	P:上下文context	G:协程
	//M上存在P,P下开启G,G可以以队列形式挂在(也有可能是链表)
	//MPG模式运行的状态:
	//主线程M0运行中,当遇到阻塞(G0协程发生阻塞),会创建M1线程(也可能从已有的线程中取出),将其他等待的线程G迁移至M1中执行
	//也就是当遇到程序阻塞时,go会创建新的内核级的线程,将等待队列中的协程挂载至新的内核级线程中执行
	//好处即,既可以等待G0的阻塞,也可以继续执行剩余的G1,G2
	//当M0不阻塞时,会将原来的协程G迁移回M0中继续执行(同时唤醒G0)

	//设置cpu核数	1.8以后默认调用多核
	part6func.ShowCPUInfo()

	//关于多线程调用同一对象出现的线程安全问题	通过go build -race 会打印程序竞争的日志
	//多协程调用同一对象会出现concurrent data writes错误:多个协程访问同一对象
	//解决方案1:对象加锁	(低端程序员..)
	for i := 0; i < 20; i++ {
		go part6func.CallGoroutineByLock(i)
	}

	time.Sleep(time.Second * 10)

	for i, i2 := range part6func.LockedGlobalMap {
		fmt.Println(i, "  ", i2)
	}

	//方案2:使用channel	(线程安全)	channel是引用类型	FIFO先进先出机制
	//类似于队列机制,go会去维护channel以保证线程安全
	//channel存在类型,只能存放同一类型的数据;	或者丢入空接口,存放任意类型,但取出时还需要断言,操作繁琐,高手向
	//语法: var dataName chan dataType
	//channel是引用类型,必须make初始化后才可使用!

	part6func.TestBasicChannel()       //基本的chan使用
	part6func.TestAnyTypeDataChannel() //传入任意类型的数据至chan
	part6func.TestTraverseChannel()    //chan的遍历

	part6func.ChannelExercise() //练习，通过chan和goroutine实现多协程数据读写

	part6func.ShowPrime() //练习，通过协程求1~n之间的素数

	//管道默认为读写，也可设置只读只写
	var writeOnlyChan chan<- int //只写
	var readOnlyChan <-chan int  //只读  注意，只读的chan不可以close！只写没问题
	//创建一个只读只写的chan意义不大，一般作为参数时，规定chan参数只读/写
	//readFunc(ch chan <- int)
	//writeFunc(ch <- chan int)
	//ch还是那个chan，只不过读写的属性被限制了
	fmt.Println(writeOnlyChan, readOnlyChan)

	//select，用于监听channel有关的io操作
	//实际开发中，一旦遇到需要遍历读取没有办法关闭的管道时，使用select是一个解决办法
	chan1 := make(chan int, 10)
label1:
	for ii := 0; ii < 10; ii++ {
		i := rand.Intn(100)
		//go对于select模块中的case，会随机执行一条，如果失败则从上之下，从左至右的顺序执行
		//如果所有case都无法执行则执行default，如果default也没有/无法执行则会进行阻塞，直到有一个io可以执行
		select {
		case chan1 <- i:
			fmt.Println("put into chan1: ", i) //一直向chan1输入数据
		case i2 := <-chan1:
			fmt.Println("get from chan1", i2) //向chan1取出数据
		default:
			fmt.Println("nothing to do")
			break label1 //个人建议使用break+label来停止循环，因为单break只会停止select（本案例中的for循环并非无限）
		}
	}

	//goroutine中也可以使用defer+recover，模板语句[deferr]已添加至goland

}
