package main

import "fmt"

func main() {
	//汉字遍历是按照字节码遍历，中文占据3个字节，会乱码
	var str string = "hello,world!北京"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}

	//采用切片方式可解决中文乱码的问题
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("%c\n", str2[i])
	}

	//for-range按照字符遍历，中文不会乱码
	for index, val := range str {
		fmt.Printf("index=%d,val=%c\n", index, val)
	}

	var max int = 100
	var sum int = 0
	var count = 0
	for i := 1; i <= max; i++ {
		if i%9 == 0 {
			sum += i
			count++
		}
	}
	fmt.Printf("1-100之间所有是9的倍数的整数的个数%d,总和为%d\n", count, sum)

	var n int = 6
	for i := 0; i <= n; i++ {
		fmt.Printf("%v + %v = %v \n", i, n-i, n)
	}
}
