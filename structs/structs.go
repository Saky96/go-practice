package main

import (
	"example/structs/user"
)

func main() {
	userFirstName := user.GetUserData("Please enter your first name: ")
	userLastName := user.GetUserData("Please enter your last name: ")
	userBirthdate := user.GetUserData("Please enter your birthdate (MM/DD/YYYY): ")

	appUser, err := user.NewCreateUserValidation(userFirstName, userLastName, userBirthdate)

	if err != nil {
		return
	}

	// ... do something awesome with that gathered data!
	// var userData user
	// userData := user.User{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthDate: userBirthdate,
	// 	createdAt: time.Now(),
	// }
	// userData.OutputUserDataStructMethod()
	// outputUserData(userData) //normally using struct data in a function
	// outputUserDataUsingPointer(&userData) // pointer to structs are passed
	// userData.clearUserName()
	// userData.outputUserDataStructMethod()
	appUser.OutputUserDataStructMethod()

	userData2 := user.New(
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
