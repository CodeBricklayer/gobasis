package main

import "fmt"

func main() {
	//for循环的第1种写法
	for i := 1; i <= 10; i++ {
		fmt.Println("你好啊", i)
	}

	//for循环的第2种写法
	j := 1
	for j <= 10 {
		fmt.Println("你好啊~", j)
		j++
	}

	//for循环的第3种写法
	for { //等价于 for ; ; {
		fmt.Println("你好啊!!!!!")
		break
	}

}
