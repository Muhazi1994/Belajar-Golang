package main

import "fmt"

func main() {
	// var myMap map[string]int
	// myMap = map[string]int{}

	// myMap["Ruby"] = 9
	// myMap["JavaScript"] = 8
	// myMap["Go"] = 10

	myMap := map[string]string{
		"Ruby": "Is Beutifull",
		"Go":   "is Super fast",
		"javaScript": "Hype",
	}
	// delete(myMap, "Go")

	// for key, value := range myMap{
	// 	fmt.Println(key,value)
	// }
	
	// fmt.Println(myMap)

	value, isAvailable := myMap["javaScript"]
	fmt.Println(isAvailable)
	fmt.Println(value)
}
