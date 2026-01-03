package main

import (
	"fmt"
	"maps"
)

func main() {
	//maps -> hash table, object, dictionary
	// make(map[key_type]val_type)
	

	//IMP : if key does not exist in map then it returns 0 value

	// m := make(map[string]string)

	// m["hani"] = "rawat"
	// m["krishna"] = "rawat"

	// fmt.Println(m["hani"])

	//using map for practice
	// slice_map := make(map[string]int)
	// slice := []string{"hani", "arpu", "chitu", "vedant"}

	// for i := range(len(slice)) {
	// 	slice_map[slice[i]] = i
	// }
	// fmt.Println(slice_map)
	// fmt.Println(len(slice_map)) //size of map
	// delete(slice_map, "chitu") //deletion of element
	// clear(slice_map) //clearing the map
	// fmt.Println(slice_map)

	// m := map[string]int {
	// 	"price" : 150,
	// 	"size" : 7,
	// 	"phone" : 3,
	// }
	// fmt.Println(m)

	// //check if elem in map or not
	// val, ok := m["price"]

	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }
	// fmt.Println(val)

	m1 := map[string]int {
		"price" : 150,
		"size" : 7,
		"phone" : 3,
	}
	m2 := map[string]int {
		"price" : 150,
		"size" : 7,
		"phone" : 3,
	}

	fmt.Println(maps.Equal(m1, m2))
	
}