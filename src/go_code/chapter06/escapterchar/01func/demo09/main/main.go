package main

import (
	"fmt"
	"go_code/chapter06/escapterchar/01func/demo09/utils"
)

//运行顺序 引入包的全局变量定义 -> 引入包的init函数 -> 全局变量定义 -> init函数 ->main函数
var age = test()

func init() {
	fmt.Println("init()...")
}

func test() int {
	fmt.Println("test()...")
	return 90
}
func main() {
	fmt.Println("main()....age=", age)
	fmt.Printf("Age=%v,Name=%v\n", utils.Age, utils.Name)
}
