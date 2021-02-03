package main

import "fmt"

type myFunType func(int, int) int

//在go中，函数也是一种数据类型
//可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数调用
func getSum(n1 int, n2 int) int {
	return n1 + n2
}

//函数既然是一种数据类型，因此在go中，函数可以作为形参，并且调用
func myFun(funVar myFunType, num1 int, num2 int) int {
	return funVar(num1, num2)
}

//返回结果指定名称
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 + n2
	return
}

//可变参数，可变参数需要放到形参列表最后
func sum(n1 int, args ...int) (sum int) {
	sum = n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return
}

//交换两个参数的值
func swap(n1 *int, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
}

func main() {
	a := getSum
	fmt.Printf("a的类型%T,getSum的数据类型是%T\n", a, getSum)

	res := a(10, 20)
	fmt.Println("res=", res)

	res2 := myFun(getSum, 50, 60)
	fmt.Println("res2=", res2)
	type myInt int
	var num1 myInt
	var num2 int
	num1 = 40
	num2 = int(num1)
	fmt.Println("num1=", num1, "num2=", num2)
	res3 := myFun(getSum, 500, 500)
	fmt.Println("res3=", res3)

	b, c := getSumAndSub(1, 2)
	fmt.Printf("b=%v,c=%v\n", b, c)

	res4 := sum(10, 10, 20, -1)
	fmt.Println("res4=", res4)
	n1 := 10
	n2 := 20

	//传入地址
	swap(&n1, &n2)
	fmt.Printf("n1=%v,n2=%v\n", n1, n2)

}
