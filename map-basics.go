package main

import (
	"fmt"
)

func main() {
	//var myMap map[string]string = make(map[string]string)
	//myMap["hello"] = "hhh"

	myMap2 := map[string]string{"name": "navneet", "name2": "arief", "name3": "venkat"}
	for key, value := range myMap2 {
		fmt.Println()
		fmt.Println(key)
		fmt.Println(value)
	}

	_, ok := myMap2["name4"]
	if ok {
		fmt.Println("key present")

	} else {
		fmt.Println("key not present")
	}
	
	//length of map
	
	fmt.Println(len(myMap2))
	
	//delete an element from mymap2
	
	delete(myMap2, "name") // if the key is not present then no error 
	
	fmt.Println(len(myMap2))
	
	// no duplicate keys in a map
	
	// only compare with nil , if really want to compare 2 maps , then use DeepEqual in reflect package.
	
	
	
}
