package main

import "fmt"

//演示字符类型的使用
func main() {
	var c1 byte = 'a'
	var c2 byte = '0'

	//直接输出byte时为字符的ASCII码值
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)

	/**
	对字符进行格式化%c
	*/
	fmt.Printf("c1=%c,c2=%c\n", c1, c2)

	/**
	超过ASCII码值255，用int接收
	*/
	var c3 int = '北'
	fmt.Printf("c3=%c,c3对应的码值=%d\n", c3, c3)

	var c4 int = 22269
	fmt.Printf("c4=%c\n", c4)
	//字符类型时可以进行运算的，相当于一个整数，运算时是按照码值进行计算的
	var n1 = 10 + 'a'
	fmt.Println("n1=", n1)

}
