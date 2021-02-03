package main

import "fmt"

//全局匿名函数定义
var (
	fun1 = func(n1 int, n2 int) int {
		return n1 * n2
	}
)

//匿名函数的演示
func main() {
	//在定义匿名函数时就直接调用，这种方式匿名函数只能调用一次
	res1 := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res1=", res1)

	//将匿名函数func (n1 int, n2 int) int 赋给变量a
	//则a的数据类型就是函数类型，此时，我们可以通过a完成调用
	a := func(n1 int, n2 int) int {
		return n1 - n2
	}

	res2 := a(10, 20)
	fmt.Println("res2=", res2)

	res3 := fun1(4, 9)
	fmt.Println("res3=", res3)

}
