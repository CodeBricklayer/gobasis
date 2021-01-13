package main

import "fmt"

func main() {
	a := 9
	b := 2
	fmt.Printf("交换前a=%v b=%v\n", a, b)
	t := a
	a = b
	b = t
	fmt.Printf("交换后a=%v b=%v\n", a, b)

	a += 7
	fmt.Println("a=", a)

	var i int = 9
	var x int = 10
	i = i + x
	x = i - x
	i = i - x
	fmt.Println(i, x)

	i = i ^ x
	x = i ^ x
	i = i ^ x
	fmt.Println(i, x)
}
