package main

import "fmt"

//This person struct type has name and age fields.
type person struct {
	name string
	age  int
}

//newPerson constructs a new person struct with the given name.You can safely return a pointer to local variable as a local variable will survive the scope of the function.
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	//This syntax creates a new struct.You can name the fields when initializing a struct.Omitted fields will be zero-valued.
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	//An & prefix yields a pointer to the struct.
	fmt.Println(&person{name: "Ann", age: 40})

	//It’s idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}

	//Access struct fields with a dot.
	fmt.Println(s.name)

	//You can also use dots with struct pointers - the pointers are automatically dereferenced.Structs are mutable.
	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age)
}
