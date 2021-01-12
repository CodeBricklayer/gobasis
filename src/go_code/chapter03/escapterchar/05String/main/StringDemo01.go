package main

import "fmt"

//字符类型的演示
func main() {
	//string的基本使用
	var address string = "北京长城 110 hello wirld！"
	fmt.Println(address)

	//字符串一旦赋值了，字符串就不能修改了
	//var str = "hello"
	//str[0] = 'a'

	//字符串的两种表示形式（1）双引号，会识别转义字符（2）反引号，以字符串的原生形式输出
	str2 := "abc\nabc"
	fmt.Println(str2)

	str3 := `package main

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
`
	fmt.Println(str3)

	var str = "hello " + "word"
	str += " haha!"
	fmt.Println(str)

	//字符串换行写时 +号必须放在行尾
	var str4 = "hello " + "word" +
		"hello " + "word" +
		"hello " + "word" +
		"hello " + "word"
	fmt.Println(str4)
}
