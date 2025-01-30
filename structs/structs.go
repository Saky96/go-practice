package main

import (
	"example/structs/user"
	"fmt"
	"time"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.NewCreateUserValidation(userFirstName, userLastName, userBirthdate)

	if err != nil {
		return
	}

	// ... do something awesome with that gathered data!
	// var userData user
	userData := user.User{
		FirstName: userFirstName,
		LastName:  userLastName,
		Birthdate: userBirthdate,
		CreatedAt: time.Now(),
	}
	userData.OutputUserDataStructMethod()
	// outputUserData(userData) //normally using struct data in a function
	// outputUserDataUsingPointer(&userData) // pointer to structs are passed
	// userData.clearUserName()
	// userData.outputUserDataStructMethod()
	appUser.OutputUserDataStructMethod()

	userData2 := user.NewCreateUser(
		"Shaurya",
		"Saigal",
		"17/02/2004",
	)

	userData2.OutputUserDataStructMethod()

	userData3 := user.NewCreateUserPointer(
		"avi",
		"Saigal",
		"22/11/2005",
	)

	userData3.OutputUserDataStructMethod()

}

func outputUserData(u user.User) {
	println(u.FirstName, u.LastName, u.Birthdate)
}

func outputUserDataUsingPointer(u *user.User) {
	usr := *u
	println(usr.FirstName, usr.LastName, usr.Birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
