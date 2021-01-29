package main

import (
	"fmt"
	//取包名别名为util
	util "go_code/chapter06/escapterchar/01func/demo03/utils"
)

func main() {
	fmt.Println("utils Num1 = ", util.Num1)
	result := util.Cal(1.2, 3.4, '+')
	fmt.Println("result=", result)
}
