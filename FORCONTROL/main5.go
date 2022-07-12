package main


import "fmt"

func main(){
	// for i := 1; i <= 100; i++ {
	// 	fmt.Println("saya sedang belajar golang", i)
	// }
	
	// i := 1
	// for i <= 100 {
	// 	fmt.Println("saya sedang belajar golang :", i)
	// 	i++
	//

	// title := "Golang the best language"
	// for index, letter := range title {
	// 	fmt.Println("index :", index, "letter :", string (letter))
	// }

	// title := "Golang the best language"
	// for index, letter := range title {
	// 	if index%2 == 0 {
	// 		fmt.Println("index:", index, "lertter:", string(letter))
	// 	}
	// }

	title := "Golang the best language"
	for index, letter := range title {
	letterString := string(letter)

	switch letterString {
	case "a","i","u","e","o":
		fmt.Println("index:", index, "lertter:", string(letter))
	}
}
}