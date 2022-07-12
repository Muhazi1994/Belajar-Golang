package main

import "fmt"

func main() {
	// numberA := 5
	// numberB := &numberA // alamat memori number A

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB) // deriferensing

	// *numberB = 10

	// fmt.Println(*numberB)
	// fmt.Println(numberA)

// =======================================

	// var numberA int = 5
	// var numberB *int = &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// numberA = 20

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)
// ======================================
	number :=5
	fmt.Println("alamat memory :", &number)
	fmt.Println("Nilai awal :", number)

	change(&number, 100)

	fmt.Println("Nilai akhir :", number)
	fmt.Println("alamat memory :", &number)
}

func change(old *int, new int) {
	*old = new
	fmt.Println("alamat memory :", old)
	fmt.Println("Di Dalam Function :", *old)

}