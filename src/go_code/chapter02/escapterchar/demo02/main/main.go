package main

import "fmt"

func getUserInfo() (string, int) {
	return "张三", 10
}
func main() {
	fmt.Println("声明变量")
	/**
	变量名称规则：
	1、字母、数字、下划线组成
	2、首个字符不能为数字
	3、不能为关键字和保留字
	4、同一作用域内不支持重复声明
	5、短变量声明法，只能用于声明局部变量，不能用于全局变量的声明
	6、变量名区分大小写
	*/
	var username string
	fmt.Println(username)

	var a1 = "张三"
	fmt.Println(a1)
	var m_a = "李四"
	fmt.Println(m_a)

	age := 20
	fmt.Println(age)

	var a2, a3 string
	a2 = "aaa"
	a3 = "bbbb"
	fmt.Println(a2, a3)

	var (
		b1 string
		b2 int
		b3 string
	)
	b1 = "test"
	b2 = 20
	b3 = "ok"
	fmt.Println(b1, b2, b3)

	var (
		c1 = "张三"
		c2 = 20
		c3 = "男"
	)
	fmt.Println(c1, c2, c3)

	d1 := "张三"
	fmt.Println(d1)

	fmt.Printf("类型：%T", d1)
	i, j, k := 12, 13, "C"
	fmt.Println(i, j, k)

	/**
	忽略某个值可用匿名变量，匿名变量用一个_表示
	*/
	s, _ := getUserInfo()
	var _, i2 = getUserInfo()
	fmt.Println(s, i2)
}
