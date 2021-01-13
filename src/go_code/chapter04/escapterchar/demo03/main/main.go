package main

import "fmt"

func main() {
	var a bool = true
	var b bool = false
	fmt.Println(a && b)
	fmt.Println(a || b)
	fmt.Println(!a)
	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}
	if age > 30 && age < 40 {
		fmt.Println("ok2")
	}

	var i int = 10
	//短路与
	if i < 9 && test() {
		fmt.Println("ok....")
	}

	//短路或
	if i > 9 || test() {
		fmt.Println("hello....")
	}
}

func test() bool {
	fmt.Println("test...")
	return true
}
