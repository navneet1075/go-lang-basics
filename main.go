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

// func (per secretAgent) speak() {
// 	fmt.Println("my name is", per.fname, per.lname, per.licenseToKill)
// }

func (per secretAgent) speak() {
	fmt.Println("my name is", per.fname, per.lname, per.licenseToKill)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	p1 := person{
		"navneet",
		"pathak",
	}
	saySomething(p1)

	sa1 := secretAgent{
		person{"james",

			"bond"},
		true,
	}
	saySomething(sa1)
	//sa1.person.speak()
}
