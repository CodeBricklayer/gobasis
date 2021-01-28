package main

import "fmt"

func main() {
	//continue演示
label1:
	for i := 0; i < 4; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				continue label1
			}
			fmt.Println("i=", i, "j=", j)
		}
	}

	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("奇数是", i)
	}

	//键盘输入不确定的整数，并判断读入的正数和负数的个数，输入为0时结束程序
	var positiveCount int
	var negativeCount int
	var num int
	for {
		fmt.Println("请输入一个整数")
		fmt.Scanln(&num)
		if num == 0 {
			break
		}
		if num > 0 {
			positiveCount++
			continue
		}
		negativeCount++
	}
	fmt.Printf("正数的个数：%v 负数的个数：%v", positiveCount, negativeCount)
}
