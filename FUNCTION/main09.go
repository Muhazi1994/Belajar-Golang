package main

import "fmt"

func main() {
	// sentence := printMyResult("saya suka dan")
	// fmt.Println(sentence)
	//===========================//

	// result := add(22,23,)
	// fmt.Println(result)
	//===========================//

	luas, keliling := calculate(10, 2)
	fmt.Println(luas)
	fmt.Println(keliling)
}
// 	func printMyResult(sentence string) string  {
// 		newSentence := sentence + " ingin belajar golang"
// 		return newSentence
// }
//===========================//

// func add (number int, numberTwo int) int {
// 	// result := number + numberTwo
// 	// return result
// 	return number + numberTwo
// }
//===========================//

func calculate(panjang int, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas,keliling
}