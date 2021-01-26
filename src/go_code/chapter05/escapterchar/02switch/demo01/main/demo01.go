package main

import (
	"fmt"
)

func main() {
	var key byte
	fmt.Println("请输入一个字符 a,b,c")
	fmt.Scanf("%c", &key)
	switch key {
	case 'a':
		fmt.Println("周一，猴子穿新衣")
	case 'b':
		fmt.Println("周二，猴子当小二")
	case 'c':
		fmt.Println("周三，猴子爬雪山")
	default:
		fmt.Println("输入有误。。。")
	}

	var age int = 10
	//类似if-else分支
	switch {
	case age == 10:
		fmt.Println("age == 10")
	case age == 20:
		fmt.Println("age == 20")
	default:
		fmt.Println("没有匹配到")
	}

	//switch的穿透 fallthrough
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		//默认只能穿透一层
		fallthrough
	case 20:
		fmt.Println("ok2")
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配到")
	}

	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型：%T", i)
	case int:
		fmt.Println("x 是 int 型")
	case float64:
		fmt.Println("x 是 float64 型")
	case func(int) float64:
		fmt.Println("x 是 func(int) 型")
	case bool, string:
		fmt.Println("x 是 bool 或 string 型")
	default:
		fmt.Println("未知类型")

	}
}
