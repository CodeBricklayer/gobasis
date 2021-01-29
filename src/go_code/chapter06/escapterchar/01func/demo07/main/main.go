package main

import "fmt"

//n1就是 *int类型
func test(n1 *int) {
	*n1 = *n1 + 10
	fmt.Printf("test() n1的地址=%v\n", &n1)
	fmt.Println("test() n1=", *n1)
}

//传址演示
func main() {
	num := 20
	fmt.Printf("main() num的地址=%v\n", &num)
	test(&num)
	fmt.Println("main() num=", num)

}
