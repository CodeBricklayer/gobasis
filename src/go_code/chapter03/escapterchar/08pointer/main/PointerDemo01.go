package main

import "fmt"

//指针类型的演示
func main() {

	var i int = 10
	//&i 是获取i的内存地址
	fmt.Println("i的地址是", &i)

	/*1.ptr是一个指针变量
	2.ptr的类型是*int
	3.ptr本身的值是&i
	*/
	var ptr *int = &i
	fmt.Printf("ptr=%v ptr的类型=%T\n", ptr, ptr)
	fmt.Println("ptr的地址是", &ptr)
	fmt.Printf("ptr指向的值=%v\n", *ptr)

	//修改指针指向的值会修改指针指针对应的内存地址对应的值
	*ptr = 9
	fmt.Println("i=", i)

	var a int = 300
	var b int = 400
	var ptr1 *int = &a
	*ptr1 = 100
	ptr1 = &b
	*ptr1 = 200
	fmt.Printf("a=%d,b=%d,*ptr1=%d", a, b, *ptr1)

}
