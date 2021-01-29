package utils

import "fmt"

var Num1 int = 300

//函数名大写才能被其他包引入使用，类似于public函数
func Cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符号有误")
	}
	return res
}
