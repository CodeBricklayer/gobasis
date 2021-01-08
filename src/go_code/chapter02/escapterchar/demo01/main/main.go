package main

import "fmt"

func main() {
	fmt.Println("tom\tjack")
	fmt.Println("tom\rjack")
	fmt.Println("天龙八部雪山飞狐\r张飞")
	fmt.Println("tom\\jack")
	fmt.Println("tom\"jack")
	fmt.Print("a", "b", "c")
	fmt.Println("a", "b", "c")
	//go语言中变量定义以后必须要使用
	var a = "aaa"
	fmt.Printf("%v", a)
	var i int = 10
	var j int = 20
	k := 5
	fmt.Println()
	fmt.Printf("i=%v,j=%v,k=%v", i, j, k)
	fmt.Printf("i=%v i的类型是%T", i, i)
}
