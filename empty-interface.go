package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type Dog struct {
	Age interface{}
}

func main() {

	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	dog := Dog{}
	dog.Age = "3"

	fmt.Printf("%T %v", dog.Age, dog.Age)
	fmt.Println()
	dog.Age = 3
	fmt.Printf("%T %v", dog.Age, dog.Age)

	x := []int{1, 2, 3, 4}

	s1 := make([]interface{}, len(x))
	for i, v := range x {
		s1[i] = v
	}

	fmt.Printf("%v", s1)

}

// copy slice of any type to slice of type interface{}
