package main

import "fmt"

func main() {
	//1的补码 0000 0001 >> 2
	//0000 0000(补码) -> 0000 0000(反码) -> 0000 0000(原码)
	var a int = 1 >> 2
	//-1的补码 1111 1111 >> 2
	//1111 1111(补码) -> 1111 1110(反码) ->1000 0001(原码)
	var b int = -1 >> 2
	var c int = 1 << 2
	var d int = -1 << 2

	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)

	fmt.Println(2 & 3)
	fmt.Println(2 | 3)
	fmt.Println(2 ^ 3)

	// -2的补码1111 1110 异或 2的补码 0000 0010
	//1111 1100(补码) -> 1111 1011(反码) -> 1000 0100(原码)
	fmt.Println(-2 ^ 2)

}
