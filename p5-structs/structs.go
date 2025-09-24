package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserdata("Enter your first name: ")
	userLastName := getUserdata("Enter your last name: ")
	userBirhDate := getUserdata("Enter your birth date (MM/DD/YYY): ")

	var appUser *user.User
	appUser, err := user.New(userFirstName, userLastName, userBirhDate)
	appUser.DisplayUserData()
	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("m@,.com", "qwerty12345")
	admin.User.DisplayUserData()
	admin.DisplayUserData()

}

func getUserdata(promtText string) string {
	fmt.Print(promtText)
	var value string
	fmt.Scanln(&value)
	return value
}
