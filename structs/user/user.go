package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u User) OutputUserDataStructMethod() {
	println(u.firstName, u.lastName, u.birthDate)
}

func (u *User) ClearUserName() { // Passing pointer in this case so that the same value is edited not the copy of the struct object
	u.firstName = "" // shortcut syntax do not need to dreference the pointer, automatically done by go.
	u.lastName = ""
}

func New(firstName, lastName, birthdate string) User { //Constructor Function in go with name 'New'
	return User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}
}

func NewCreateUserPointer(firstName, lastName, birthdate string) *User { //Constructor Function using pointer in go, that created only one copy for argumants and return types
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}
}

func NewCreateUserValidation(firstName, lastName, birthdate string) (*User, error) { //Constructor used for validation
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstName, lastName and birthDate is required!")
	}
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthdate,
		createdAt: time.Now(),
	}, nil

}

func outputUserData(u User) {
	println(u.firstName, u.lastName, u.birthDate)
}

func outputUserDataUsingPointer(u *User) {
	usr := *u
	println(usr.firstName, usr.lastName, usr.birthDate)
}

func GetUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
