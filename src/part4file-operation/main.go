package main

import (
	"bufio"
	"fmt"
	"golang-basic-project/src/part4file-operation/exercise"
	exerciseJson "golang-basic-project/src/part4file-operation/json"
	"io"
	"io/ioutil"
	"os"
)

//文件操作
func main() {
	//此处的file可以称为 file对象 file指针 file文件句柄 		goland不会自动添加\\,但golang在win下也可使用/路径
	filepath := "C:/Users/Ash/Documents/Code/GO/golang-first-exercise-project/src/part4file-operation/test/fortest.txt"
	file, err1 := os.Open(filepath)
	if err1 != nil {
		fmt.Println("file open 	err: ", err1)
	}
	fmt.Println("file: ", file) //打印的结果就是一个指针,指针,指针!

	//===============
	//文件的读取
	//err2 := file.Close() //记得关闭
	//if err2 != nil {	//捕获错误
	//	fmt.Println("file close err: ",err2)
	//}
	defer file.Close() //也可以使用defer关键词,在函数执行完毕前执行	os.Open需要close,缓冲流的file大多数已封装,自动会close

	reader := bufio.NewReader(file) //golang默认缓冲区4096
	for {
		str, err3 := reader.ReadString('\n') //delim分隔符,当读到换行符后进行分割
		fmt.Println("string : ", str)
		if err3 == io.EOF { //eof:文件末尾	当读至文末时中断	此处如果先判断是否文末,会导致最后一行不打印
			break
		}
	}

	//一次性读取整个文件,适用于文件不大的情况	此函数不需要close,close已经封装至函数中
	readFile, err3 := ioutil.ReadFile(filepath)
	if err3 != nil {
		fmt.Println("read file err : ", err3)
	}
	//fmt.Println("read file : ",readFile)			//此时输出的结果时byte
	fmt.Println("read file :\n" + string(readFile)) //需要string进行转换

	//===============
	//文件的写入
	filepath2 := "C:/Users/Ash/Documents/Code/GO/golang-first-exercise-project/src/part4file-operation/test/file-write-test.txt"
	//参数:	文件路径	,打开方式	,Linux下的权限(对win无效)
	openFile, err4 := os.OpenFile(filepath2, os.O_WRONLY|os.O_CREATE, 0666) //使用只写,此处|代表组合?
	if err4 != nil {
		fmt.Println("open file err: ", err4)
	}
	defer openFile.Close() //也需要延迟关闭
	/*
	 * 关于文件打开方式
	 * O_RDONLY		只读
	 * O_WRONLY		只写
	 * O_RDWR		读写
	 * O_APPEND		添加至末尾
	 * O_CREAT		如不存在则创建
	 * O_EXCL		与↑结合使用,文件必须不存在
	 * O_SYNC		同步IO流
	 * O_TRUNC		如可能,打开时清空文件
	 */

	//对文件进行操作
	str1 := "hello Gorden\n"
	//获取写入指针
	writer := bufio.NewWriter(openFile)
	for i := 0; i < 3; i++ { //遍历输入3遍
		_, _ = writer.WriteString(str1) //通过写入指针进行写入操作,注意,此处写入的是缓存!
		// goland可以A+Enter来忽视参数返回
	}
	//需要刷新缓存,才是真正的写入文件(和java的buffer类似)
	_ = writer.Flush()

	//读写+追加的操作
	openFile2, err5 := os.OpenFile(filepath2, os.O_RDWR|os.O_APPEND, 0666) //使用读写+追加的方式打开
	if err5 != nil {
		fmt.Println("open file err: ", err5)
	}
	defer openFile2.Close()
	str2 := "append word\n"
	writer2 := bufio.NewWriter(openFile)
	for i := 0; i < 3; i++ {
		_, _ = writer2.WriteString(str2)
	}
	_ = writer2.Flush()

	//查看写入的结果
	readFile2, err6 := ioutil.ReadFile(filepath2)
	if err6 != nil {
		fmt.Println("read file err6 : ", err3)
	}
	fmt.Println("read file :\n" + string(readFile2))

	//练习,复制A.txt至B.txt
	copyTxt()

	//练习,复制图片
	filepathSrc := "C:/Users/Ash/Documents/Code/GO/golang-first-exercise-project/src/part4file-operation/test/src.gif"
	filepathDst := "C:/Users/Ash/Documents/Code/GO/golang-first-exercise-project/src/part4file-operation/test/dst.gif"
	img, err7 := copyImg(filepathDst, filepathSrc)
	if err7 != nil {
		fmt.Println(err7)
	}
	fmt.Println(img)

	//练习,统计一个txt内有多少字母,数字,空格
	exercise.CountStr("C:/Users/Ash/Documents/Code/GO/golang-first-exercise-project/src/part4file-operation/test/fortest.txt")

	//补充:go程序通过命令行执行时,可以通过os.Args接收命令行的参数,通过for-range遍历读取,因此这里不想打包成exe所以略过
	//flag包下也有类似的函数,略过

	//Json部分
	//序列化
	exerciseJson.CallNormalJson() //普通的字符串序列化
	exerciseJson.CallMapJson()    //map
	exerciseJson.CallSliceJson()  //切片

	//反序列化
	jsonStruct := "{\"age\":12,\"name\":\"狗蛋二号\"}"
	jsonMap := "{\"age\":12,\"name\":\"狗蛋二号\"}"
	jsonSlice := "[\"狗蛋三号\",\"狗蛋四号\",\"\",\"\",\"\"]"
	exerciseJson.CallUnmarshalToStruct(jsonStruct) //结构体
	exerciseJson.CallUnmarshalToMap(jsonMap)       //map
	exerciseJson.CallUnmarshalToSlice(jsonSlice)   //切片

}

//复制图片	dst:目标文件路径,src:源文件路径
func copyImg(dst string, src string) (written int64, err error) {

	//打开src文件
	srcFile, err := os.Open(src)
	defer srcFile.Close() //反正os包下的都记得关闭一下..
	if err != nil {
		fmt.Println(err)
		return written, err
	}
	//读取src的内容至newReader
	reader := bufio.NewReader(srcFile)

	//以读写|创建的方式,打开dst文件
	dstFile, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE, 0666)
	defer dstFile.Close()
	if err != nil {
		fmt.Println(err)
		return written, err
	}
	//获取writer
	writer := bufio.NewWriter(dstFile)
	defer writer.Flush() //这里的flush比较重要,如果不写会导致大文件的数据丢失一小部分,小文件直接空文件

	//使用官方库提供的Copy函数完成文件的拷贝
	io.Copy(writer, reader)
	return 666, err

}

//将copy-from.txt复制至copy-to.txt
func copyTxt() {
	filepath1 := "C:/Users/Ash/Documents/Code/GO/golang-first-exercise-project/src/part4file-operation/test/copy-from.txt"
	filepath2 := "C:/Users/Ash/Documents/Code/GO/golang-first-exercise-project/src/part4file-operation/test/copy-to.txt"

	//读a文件
	file1, err1 := ioutil.ReadFile(filepath1)
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	//写入b文件
	err2 := ioutil.WriteFile(filepath2, file1, 0666)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

}
