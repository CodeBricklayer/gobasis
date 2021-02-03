package utils

import "fmt"

var Age int
var Name string

//Age 和 Name 全局变量，我们需要在main.go使用
//但是我们需要初始化Age 和 Name

func init() {
	fmt.Println("utils 包的 inti()...")
	Age = 25
	Name = "Tom"
}
