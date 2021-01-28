package main

import "fmt"

func main() {
	fmt.Println("OK1")
	goto label1
	fmt.Println("OK2")
	fmt.Println("OK3")
	fmt.Println("OK4")
label1:
	fmt.Println("OK5")
	fmt.Println("OK6")
	fmt.Println("OK7")
	fmt.Println("OK8")
	fmt.Println("OK9")
}
