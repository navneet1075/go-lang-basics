package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {

	buf, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
	}
	//fmt.Println(buf)

	d := [][][][]uint8{}
	json.Unmarshal([]byte(buf), &d)
	fmt.Println("Length of Array:", len(buf))
	fmt.Println(d)

	// var biArray [][][][]int64
	// json.Unmarshal([]byte(`[[[[111,222,333],[444,555,666],
	//                      [777,888,999]]]]`), &biArray)
	// fmt.Println("Length of Array:", len(biArray))

	// fmt.Println("\nBi-dimensional Array\n")
	// for index, element := range biArray {
	// 	fmt.Println(index, "=>", element)
	// }

	//fmt.Printf("%v", d)
	// _ = json.Unmarshal([]byte(buf), &arr)
	// log.Printf("Unmarshaled: %v", arr)
	// var biArray [][][]int64
	// json.Unmarshal([]byte(`[[111,222,333],[444,555,666],
	//                      [777,888,999]]`), &biArray)
	// fmt.Println("Length of Array:", len(biArray))

	// fmt.Println("\nBi-dimensional Array\n")
	// for index, element := range biArray {
	// 	fmt.Println(index, "=>", element)
	// }

}
