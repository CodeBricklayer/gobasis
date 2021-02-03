package main

import "fmt"

//累加器
func addUpper() func(int) int {
	//n初始化了一次
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		str += "a"
		fmt.Println("str=", str)
		return n
	}
}

//闭包的演示
func main() {
	f := addUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}
