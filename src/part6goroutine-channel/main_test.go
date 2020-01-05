package main

import (
	"golang-basic-project/src/part6goroutine-channel/part6func"
	"testing"
)

func TestForTestPrime(t *testing.T) {
	part6func.ShowPrime() //使用test测试时间

}

func TestForTestNormalPrime(t *testing.T) { //默认还是java快。。
	part6func.ShowPrimeByNormal()

}
