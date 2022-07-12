package main

import "fmt"

func main() {
	var gamingConsole []string
	// gamingConsole := []string{"PlayStation4","Nintendo Switch","Xbox One"}

	gamingConsole = append(gamingConsole, "PlayStation4",)
	gamingConsole = append(gamingConsole, "Nintendo Switch")
	gamingConsole = append(gamingConsole, "Xbox One")
	
	for _, console := range gamingConsole {
		fmt.Println(console)
	}
	
}