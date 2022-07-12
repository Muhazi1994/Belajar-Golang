package main

import "fmt"

func main() {
	number := 5

	switch number {
	case 1:
		fmt.Println("satu")
	case 2:
		fmt.Println("Dua")
	case 3:
		fmt.Println("Tiga")
	default:
		fmt.Println("DEFAULT")
	}

	// if number == 2 {
	// 			fmt.Println("Dua")
	// } else if number == 1 {
	// } else if number == 3
}
