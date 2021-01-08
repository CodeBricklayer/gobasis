package main

import "fmt"

func main() {
	/**
	常量,值不可以改变
	*/
	const pi = 3.141592653
	fmt.Println(pi)
	const (
		A = "a"
		B = "b"
	)
	fmt.Println(A, B)

	/**
	同时声明多个变量时，如果省略了值则表示和上面一行的值相同
	*/
	const (
		n1 = 100
		n2
		n3 = 200
		n4
	)
	fmt.Println(n1, n2, n3, n4)

	/**
	iota是计数器
	*/
	const a = iota

	/**
	可跳过某个值，用_
	*/
	const (
		b = iota
		_
		c
		d
		e
	)
	fmt.Println(b, c, d, e)

	/**
	可在计数器中插队
	*/
	const (
		num1 = iota
		num2 = 100
		num3 = iota
		num4
	)
	fmt.Println(num1, num2, num3, num4)

	/**
	多个iota定义在一行
	*/
	const (
		number1, number2 = iota + 1, iota + 2
		number3, number4
		number5, number6
	)
	fmt.Println(number1, number2, number3, number4, number5, number6)
}
