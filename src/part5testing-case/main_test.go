package main_test //测试单元包名可以不是main

//需要导入testing包
import "testing" //也可以通过 _ 来忽视已经导入的包..

//单元测试
//goland可以输入test快速创建测试单元		ctrl+shift+T 快速生成单元测试文件
//测试单元必须以TestXxxx的格式命名,并传入*testing.T指针
func TestCallTest(t *testing.T) {
	result := testFunc(10)     //结果值
	expectation := 10          //期望值
	if result != expectation { //如果与期望不符
		t.Fatalf("expectation: %v,value: %v", expectation, result) //调用Fatal函数打印错误信息
	}
	t.Logf("TestCallTest success") //如果无报错则打印log日志
	//一般常用函数:Fatal致命 Log日志
}

//也可以编写多个测试单元
func TestCallTest2(t *testing.T) {
	if 1 != 1 { //如果与期望不符
		t.Fatalf("expectation: %v,value: %v", 1, 0) //调用Fatal函数打印错误信息
	}
	t.Logf("TestCallTest2 success") //如果无报错则打印log日志
}

//用于单元测试的函数
func testFunc(i int) (result int) {
	result = i
	return
}

/*

附终端打印的结果:
=== RUN   TestCallTest
--- PASS: TestCallTest (0.00s)
    main_test.go:14: TestCallTest success
=== RUN   TestCallTest2
--- PASS: TestCallTest2 (0.00s)
    main_test.go:23: TestCallTest2 success
PASS

PASS代表成功
FAIL代表失败
*/
