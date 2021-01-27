package main

import "fmt"

func main() {
	//空心金字塔
	var totalLevel int
	fmt.Println("请输入层数")
	fmt.Scanln(&totalLevel)
	for i := 1; i <= totalLevel; i++ {
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i*2-1; j++ {
			if j == 1 || j == 2*i-1 || i == totalLevel {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}
