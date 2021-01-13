package main

import "fmt"

func main() {
	fmt.Println(10 / 4)
	fmt.Println(10.0 / 4)
	fmt.Println(10 / 4.0)

	//a % b = a - a / b * b
	fmt.Println(10 % 3)
	fmt.Println(-10 % 3)
	fmt.Println(10 % -3)
	fmt.Println(-10 % -3)

	var i int = 10
	i++
	fmt.Println("i=", i)
	i--
	fmt.Println("i=", i)

	//假如还有97天放假，问：xx个星期零xx天
	fmt.Println(97/7, "个星期", 97%7, "天")

	//定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式为：5/9*(华氏温度-100)，请求出华氏温度对应的摄氏温度
	var huashi float32 = 134.2
	var sheshi float32 = 5.0 / 9 * (huashi - 100)
	fmt.Printf("华氏温度%f对应的摄氏温度为%f", huashi, sheshi)
}
