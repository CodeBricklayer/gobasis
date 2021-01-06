package gobyexample

import "fmt"

//This function intSeq returns another function, which we define anonymously in the body of intSeq. The returned function closes over the variable i to form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	//We call intSeq, assigning the result (a function) to nextInt. This function value captures its own i value, which will be updated each time we call nextInt.
	nexInt := intSeq()
	//See the effect of the closure by calling nextInt a few times.
	fmt.Println(nexInt())
	fmt.Println(nexInt())
	fmt.Println(nexInt())

	//To confirm that the state is unique to that particular function, create and test a new one.
	newInts := intSeq()
	fmt.Println(newInts())
}
