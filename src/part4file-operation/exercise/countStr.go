package exercise

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//统计一个文件有多少个字母 数字 空格

type counts struct {
	strCounts   int
	numCounts   int
	spaCounts   int
	otherCounts int
}

func CountStr(filepath string) {

	counts := counts{}

	//打开文件
	open, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer open.Close()

	//创建reader,读取文件
	reader := bufio.NewReader(open)

	for {
		//一行行读取
		readString, err := reader.ReadString('\n')
		if err == io.EOF { //打印到末尾则终止循环
			break
		}

		//读取每个字符串
		for _, i2 := range readString {
			switch {
			case i2 >= 'a' && i2 <= 'z': //此处是穿透是因为需要大小写问题
				fallthrough
			case i2 >= 'A' && i2 <= 'Z':
				counts.strCounts++
			case i2 == ' ' || i2 == '\t': //\t也是空格
				counts.spaCounts++
			case i2 >= '0' && i2 <= '9':
				counts.numCounts++
			default:
				counts.otherCounts++
			}
			//补充:此处for-range获取的i2已经是兼容utf-8汉字的,没必要再转为rune
		}
	}
	//打印结果
	fmt.Println("文件名:", filepath)
	fmt.Printf("字符串:%v,数字:%v,空格%v,其他:%v", counts.strCounts, counts.numCounts, counts.spaCounts, counts.otherCounts)

}
