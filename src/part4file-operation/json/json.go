package exerciseJson

import (
	"encoding/json"
	"fmt"
)

//用于演示json序列化

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//普通的K:V类型json
func CallNormalJson() {
	student1 := Student{
		Name: "狗蛋",
		Age:  11,
	}

	marshal, err := json.Marshal(&student1) //需要传入指针类型
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(marshal)	//marshal为byte数组
	fmt.Println(string(marshal)) //转为字符串

}

//map类型json
func CallMapJson() {
	m := make(map[string]interface{}) //创建一个map
	m["name"] = "狗蛋二号"
	m["age"] = 12
	//需要注意的是塞入的map没有顺序	map的key会根据时间随机生成一个哈希种子以此来排序

	marshal, err := json.Marshal(&m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))

}

//切片/数组类型json
func CallSliceJson() {
	s := make([]string, 5, 5) //创建一个切片
	s[0] = "狗蛋三号"
	s[1] = "狗蛋四号"
	marshal, err := json.Marshal(&s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))

}

//Unmarshal 反序列化
//反序列化至struct
func CallUnmarshalToStruct(jsonStr string) { //参数关键词不能传json，会覆盖json包。。。
	var s Student                              //创建结构体对象
	err := json.Unmarshal([]byte(jsonStr), &s) //将指针传入
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Printf("name:%v,age:%v\r", s.Name, s.Age) //也可以详细打印
	//测试后发现去掉tag也可以进行反序列化,可能反序列化不区分大小写
}

//反序列化至map
func CallUnmarshalToMap(jsonStr string) {
	var m map[string]interface{}               //反序列化map时,内置unmarshal函数会帮你make初始化map类型
	err := json.Unmarshal([]byte(jsonStr), &m) //此处传入的指针会进行类型断言
	if err != nil {
		fmt.Println(err)
	}
	//展示结果
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
}

//反序列化至切片
func CallUnmarshalToSlice(jsonStr string) {
	var s []string
	err := json.Unmarshal([]byte(jsonStr), &s)
	if err != nil {
		fmt.Println(err)
	}
	//展示结果
	fmt.Println(s)
	for k, v := range s {
		fmt.Println(k, v)
	}
}
