package main

import "fmt"

func main() {
	//跳出指定标签，不指定标签，则跳出当前最近的循环
label2:
	for i := 0; i < 4; i++ {
	label1:
		for j := 0; j < 10; j++ {
			if j == 5 {
				break label1
			}
			if j == 2 {
				break label2
			}
			fmt.Println("j=", j)
		}
	}
}
