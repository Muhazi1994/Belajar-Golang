// package main

// import "fmt"

// func main() {
// 	scores := []int{10, 5, 8, 9, 7}
// 	total := sum(scores)
// 	fmt.Println(total)

// }
// func sum(numbers []int) int {
// 	var total int
// 	for _, number := range numbers {
// 		total = total + number
// 	}
// 	return total
// }

package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := calculate(10, 2, "a")
	if err != nil {
		fmt.Println("Terjadi kesalahan Input")
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}

func calculate(number, numberTwo int, operation string) (int, error) {
	var result int
	var errorResult error
	switch operation {
	case "+":
		result = number + numberTwo
	case "-":
		result = number - numberTwo
	case "*":
		result = number * numberTwo
	case "/":
		result = number / numberTwo
	default:
		errorResult = errors.New("Unknown operation")
	}
	return result, errorResult
}
