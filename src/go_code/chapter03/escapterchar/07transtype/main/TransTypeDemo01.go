package main

import (
	"fmt"
	"strconv"
)

//数据类型转化的演示
func main() {
	var i int32 = 100

	var n1 float32 = float32(i)
	var n2 int8 = int8(i)
	var n3 int64 = int64(i)
	fmt.Printf("i=%v n1=%v n2=%v n3=%v\n", i, n1, n2, n3)

	//被转换的是变量存储的数据（即值），变量本身的数据类型并没有变化
	fmt.Printf("i type is %T\n", i)

	//超出数据范围，按照溢出处理
	var num1 int64 = 999999
	var num2 int8 = int8(num1)
	fmt.Println("num2=", num2)

	var num3 int = 99
	var num4 float64 = 23.456
	var b bool = true
	var mychar byte = 'h'
	var str string
	str = fmt.Sprintf("%d", num3)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = fmt.Sprintf("%f", num4)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q\n", str, str)
	str = fmt.Sprintf("%c", mychar)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str=%q\n", str, str)

	//‘f'格式  10表示小数位保留10位 64表示这个小数是float64
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.FormatBool(b)
	fmt.Printf("str type %T str=%q\n", str, str)

	str = strconv.Itoa(num3)
	fmt.Printf("str type %T str=%q\n", str, str)

	var str1 string = "true"
	var b1 bool
	b1, _ = strconv.ParseBool(str1)
	fmt.Printf("b1 type %T b1=%t\n", b1, b1)
	var str2 string = "123456"
	var n4 int64
	n4, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("n4 type %T n4=%d\n", n4, n4)

	var str3 string = "123.456"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("f1 type %T f1=%f\n", f1, f1)

	//string转换基本数据类型失败时，会将其转换成默认值
	var str4 string = "hello"
	var n5 int64 = 10
	n5, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("n5 type %T n5=%d\n", n5, n5)
}
