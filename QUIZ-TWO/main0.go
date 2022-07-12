package main

import "fmt"

func main() {
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	// var total int

	// for _, scores := range scores {
	// 	total = total + scores
	// }
	// length := len(scores)
	// average := float64(total) / float64(length)
	// fmt.Println(average)



	var goodScore []int
	for _, scores := range scores {
		if scores >= 90 {
			goodScore = append(goodScore, scores)
		}
		
	}
	fmt.Println(goodScore)
		
}
