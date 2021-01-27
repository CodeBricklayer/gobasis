package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count int = 0

	for {
		//生成随机数需要设置一个种子
		//时间戳
		rand.Seed(time.Now().UnixNano())
		i := rand.Intn(100) + 1
		count++
		if i == 99 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("生成 99 一共使用了 ", count)
}
