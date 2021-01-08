package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"unsafe"
)

func main() {
	//1、定义float类型
	var a float32 = 3.12
	fmt.Printf("a=%v--%f 类型：%T\n", a, a, a)
	fmt.Println("float32占用", unsafe.Sizeof(a), "字节")

	var b float64 = 3.12
	fmt.Printf("b=%v--%f 类型：%T\n", b, b, b)
	fmt.Println("float64占用", unsafe.Sizeof(b), "字节")

	//2、%f 输出float类型	%.2f输出数据的时候保留2位小数
	var c float32 = 3.1415926535
	fmt.Printf("c=%v--%f--%.2f 类型：%T\n", c, c, c, c)

	d := 3.141592653
	fmt.Printf("d=%f 类型：%T\n", d, d)
	fmt.Println("d占用", unsafe.Sizeof(d), "字节")

	//表示f2等于3.13*10的2次方
	var f2 = 3.14e2
	fmt.Printf("f2=%f 类型：%T\n", f2, f2)
	//表示f2等于3.13/10的2次方
	var f3 = 3.14e-2
	fmt.Printf("f3=%f 类型：%T\n", f3, f3)

	//精度丢失问题
	var f4 float64 = 1129.6
	fmt.Println(f4 * 100)
	m1 := 8.2
	m2 := 3.8
	fmt.Println(m1 - m2)
	//使用第三方包解决
	var num1 float64 = 3.1
	var num2 float64 = 4.2
	//加法
	d1 := decimal.NewFromFloat(num1).Add(decimal.NewFromFloat(num2))
	fmt.Println(d1)
}
