package main

import "fmt"

func main() {
	student := []map[string]string {
		{"name": "Muhazi", "score": "A"},
		{"name": "Aliya", "score": "B"},
		{"name": "Nura", "score": "E"},
	}

	for _, students := range student {
		fmt.Println(students["name"], "-", students["score"])
	}
}