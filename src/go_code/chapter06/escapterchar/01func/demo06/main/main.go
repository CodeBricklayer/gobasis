package main

import "fmt"

func main() {
	fmt.Println(fbn(7))
	fmt.Println(f(5))
	fmt.Println(peach(0))
}

//斐波那契数列
func fbn(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

//已知f(1)=3;f(n)=2*f(n-1)+1，用递归思想编写代码
func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
	}
}

//猴子吃桃子问题
//有一堆桃子，猴子第一天吃了其中的一半，并再多吃了一个！以后每天猴子都吃其中的一半，然后再多吃一个。当到第10天，想再吃时，发现只有一个桃子了
func peach(n int) int {
	if n > 10 || n < 0 {
		fmt.Println("输入的天数不对")
	}
	if n == 10 {
		return 1
	} else {
		return (peach(n+1) + 1) * 2
	}
}
