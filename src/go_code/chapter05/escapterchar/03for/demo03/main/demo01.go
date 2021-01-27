package main

import "fmt"

func main() {
	//while循环的实现
	var index int = 0
	for {
		if index > 10 {
			break
		}
		fmt.Println("hello,world ", index)
		index++
	}

	//do-while循环的实现
	var i int = 0
	for {
		fmt.Println("hello,ok ", i)
		i++
		if i > 10 {
			break
		}
	}
}
