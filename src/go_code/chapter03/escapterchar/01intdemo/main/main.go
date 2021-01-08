package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/**
	1、定义int类型
	*/
	var num int = 10

	fmt.Printf("num=%v 类型：%T", num, num)
	fmt.Println()
	/**
	2、int8的范围演示(-128 ~ 127)
	*/
	var a int8 = 98
	fmt.Printf("a=%v 类型：%T", a, a)
	fmt.Println()

	/**
	3、uint8的范围演示（0-255）
	*/
	var b uint8 = 98
	fmt.Printf("b=%v 类型：%T", b, b)
	fmt.Println()
	/**
	4、int8 int16 ...占用的存储空间大小
	*/
	fmt.Println("int8占用", unsafe.Sizeof(a), "字节")
	fmt.Println("uint8占用", unsafe.Sizeof(b), "字节")
	/**
	5、int不同长度直接的转换
	*/
	var a1 int32 = 10
	var a2 int64 = 21
	fmt.Println(int64(a1) + a2)
	fmt.Println(a1 + int32(a2))
	var n1 int16 = 130
	fmt.Println(int8(n1))
	fmt.Println(int64(n1))
	var n2 int16 = 110
	fmt.Println(int8(n2))
	fmt.Println(int64(n2))
	/**
	6、int数字字面量语法 %d表示10进制输出 %b表示二进制输出 %o表示八进制输出 %x表示十六进制输出
	*/
	number := 12
	fmt.Printf("number=%v 类型：%T", number, number)
	fmt.Println()
	fmt.Println("number占用", unsafe.Sizeof(number), "字节")
	fmt.Printf("number=%d\n", number)
	fmt.Printf("number=%b\n", number)
	fmt.Printf("number=%o\n", number)
	fmt.Printf("number=%x\n", number)
}
