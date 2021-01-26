package main

import (
	"fmt"
	"math"
)

func main() {
	var age int
	fmt.Println("请输入你的年龄")
	fmt.Scanf("%d", &age)
	if age > 18 {
		fmt.Println("你年龄大于18，要对自己的行为负责！")
	} else {
		fmt.Println("你的年龄不大，这次就放过你了")
	}

	var score int
	fmt.Println("请输入你的成绩")
	fmt.Scanf("%d", &score)
	if score == 100 {
		fmt.Println("奖励一辆BMW")
	} else if score > 80 && score <= 99 {
		fmt.Println("奖励一台iphone7plus")
	} else if score >= 60 && score <= 80 {
		fmt.Println("奖励一个ipad")
	} else {
		fmt.Println("什么奖励都没有")
	}

	var a float64 = 3.0
	var b float64 = 100.0
	var c float64 = 6.0
	m := b*b - 4*a*c
	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		x2 := (-b - math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v x=%v", x1, x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v x2=%v", x1, x1)
	} else {
		fmt.Println("无解")
	}
}
