package main

import (
	"fmt"
	"unsafe"
)

//布尔类型的演示
func main() {
	var b = false
	fmt.Println("b=", b)
	//1、bool类型占用存储空间是1个字节
	fmt.Println(unsafe.Sizeof(b))
}
