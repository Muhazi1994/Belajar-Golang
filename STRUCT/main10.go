package main

import (
	"fmt"
	"struct/management"
)


// func displayUser (user User) string {
// 	return fmt.Sprintf("Name : %s %s, email : %s", user.FirstName, user.LastName, user.Email)
// }

func main() {

	// user := User{}
	// user.ID = 1
	// user.FirstName = "Muhazi"
	// user.LastName = "Ramadhan"
	// user.Email = "muhaziramadhan261@gmail.com"
	// user.IsActive = true

	// user2 := User{}
	// user2.ID = 2
	// user2.FirstName = "Nurma"
	// user2.LastName = "Ningsih"
	// user2.Email = "nurmaningsih@gmail.com"
	// user2.IsActive = true

	// user3 := management.User{
	// 	ID:        3,
	// 	FirstName: "Asma Nurul",
	// 	LastName:  "Awaliyah",
	// 	Email:     "Asma@gmail.com",
	// 	IsActive:  true,
	// }
	// fmt.Println(user3.display())

	user := management.User{1, "Ramadhan", "Muhazi", "Ramadhan@gmail.com", true}
	result := user.Display()
	fmt.Println(result)

	user2 := management.User{2, "Link", "Awaking", "Ramadhan@gmail.com", true}
	fmt.Println(user2.Display())
	// displayUser1 := displayUser(user)
	// displayUser2 := displayUser(user2)
	// displayUser3 := displayUser(user3)
	// displayUser4 := displayUser(user4)

	// fmt.Println(displayUser1)
	// fmt.Println(displayUser2)
	// fmt.Println(displayUser3)
	// fmt.Println(displayUser4)

	users := []management.User{user,user2}
	group := management.Group{"Gamer", user, users, true}

	// displayGroup(group)
	group.DisplayGroup()
}
// func displayGroup(group Group) {
// 	fmt.Printf("Name : %s", group.Name)
// 	fmt.Println("")
// 	fmt.Printf("Member count : %d", len(group.Users))
// 	fmt.Println("")
// 	fmt.Println("User name :")

// 	for _, user := range group.Users {
// 		fmt.Println(user.FirstName)
// 	}

// }


