package main

import "fmt"

func main() {
	//100以内的数求和，求出当和第一次大于20的当前数
	var sum int = 0
	var max int = 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum > 20 {
			max = i
			break
		}
	}
	fmt.Printf("第一次大于20的数为%d", max)

	//输入用户名密码三次
	var name string
	var pwd string
	var loginChance = 3
	for i := 1; i <= 3; i++ {
		fmt.Println("请输入用户名")
		fmt.Scanln(&name)
		fmt.Println("请输入密码")
		fmt.Scanln(&pwd)
		if name != "张三" && pwd == "888" {
			fmt.Println("恭喜你登陆成功！")
			break
		} else {
			loginChance--
			if loginChance == 0 {
				break
			}
			fmt.Printf("还有%v次登陆机会机会，请珍惜！\n", loginChance)
		}
	}
	if loginChance == 0 {
		fmt.Println("机会用光了，没有登陆成功！")
	}
}
