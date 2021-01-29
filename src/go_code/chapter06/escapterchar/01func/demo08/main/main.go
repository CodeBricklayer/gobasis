package main

import "fmt"

//在go中，函数也是一种数据类型
//可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数调用
func getSum(n1 int, n2 int) int {
	return n1 + n2
}

//函数既然是一种数据类型，因此在go中，函数可以作为形参，并且调用
func myFun(funVar func(int, int) int, num1 int, num2 int) int {
	return funVar(num1, num2)
}

func main() {
	a := getSum
	fmt.Printf("a的类型%T,getSum的数据类型是%T\n", a, getSum)

	res := a(10, 20)
	fmt.Println("res=", res)

	res2 := myFun(getSum, 50, 60)
	fmt.Println("res2=", res2)
}
