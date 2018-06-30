package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (per person) speak() {
	fmt.Println("my name is", per.fname, per.lname)
}

func (per secretAgent) speak() {
	fmt.Println("my name is", per.fname, per.lname, per.licenseToKill)
}

func main() {

	p1 := person{
		"navneet",
		"pathak",
	}
	p1.speak()

	sa1 := secretAgent{
		person{"james",

			"bond"},
		true,
	}
	sa1.speak()
}
