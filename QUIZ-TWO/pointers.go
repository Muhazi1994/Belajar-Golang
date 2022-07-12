package main

import "fmt"
type Gamer struct {

	Name string
	Games []string
}

func (gamer *Gamer) AddGame(game string) {
	gamer.Games = append(gamer.Games,game) 
}

func main() {
	gamer := Gamer{Name:"Muhazi"}
	fmt.Println(gamer.Name)

	gamer.AddGame("BorderLands")
	gamer.AddGame("Bioshock")
	gamer.AddGame("Just Dance")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
	}