package main

import "fmt"

func main() {
	var c1 byte = 'a'
	var c2 byte = '0'
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
}
