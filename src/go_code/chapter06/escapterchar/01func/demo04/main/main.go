package main

import "fmt"

func main() {
	n1 := 10
	test(n1)
	fmt.Println("main() n1=", n1)
	sum := getSum(10, 20)
	fmt.Println("main() sum=", sum)
	res1, res2 := getSumAndSub(2, 1)
	//忽略赋值，用占位符忽略
	_, _ = getSumAndSub(2, 1)
	fmt.Printf("sum=%v sub=%v\n", res1, res2)
}

func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("test() n1=", n1)
}

func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Println("getSum() sum=", sum)
	return sum
}

func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}
