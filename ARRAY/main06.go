package main

import "fmt"

func main(){
	// var languages [5]string
	// languages[0] = "Go"
	// languages[1] = "C#"
	// languages[2] = "Ruby"
	// languages[3] = "JavaScript"
	// languages[4] = "Python"

	// languages := [5]string{"Ruby","JavaScript","Go","C#","Python"}

	languages := [...]string{
		"Ruby",
		"JavaScript",
		"Go",
		"C#",
		"Python",
		"Kotlin",
	}
for index, lang := range languages {
	fmt.Println(index,lang)
}
	// fmt.Println(languages)
	// length := len(languages)
	// fmt.Println(length)
}