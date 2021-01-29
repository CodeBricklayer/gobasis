package utils

import "fmt"

func SayOk() {
	fmt.Println("Ok~~")
}

//在同一包下，不能有相同的函数名（也不能有相同的全局变量名）
//func Cal()  {
//}
