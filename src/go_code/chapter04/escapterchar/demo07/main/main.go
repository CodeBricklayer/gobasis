package main

import "fmt"

func main() {
	var i int = 5
	//使用二进制输出
	fmt.Printf("%b\n", i)

	//八进制以0开头
	var j int = 011
	fmt.Println("j=", j)

	//十六进制以0x开头
	var k int = 0x11
	fmt.Println("k=", k)
}
