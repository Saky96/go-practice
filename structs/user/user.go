package user

import (
	"errors"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Birthdate string
	CreatedAt time.Time
}

func (u User) OutputUserDataStructMethod() {
	println(u.FirstName, u.LastName, u.Birthdate)
}

func (u *User) ClearUserName() { // Passing pointer in this case so that the same value is edited not the copy of the struct object
	u.FirstName = "" // shortcut syntax do not need to dreference the pointer, automatically done by go.
	u.LastName = ""
}

func NewCreateUser(firstName, lastName, birthdate string) User { //Constructor Function in go
	return User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now(),
	}
}

func NewCreateUserPointer(firstName, lastName, birthdate string) *User { //Constructor Function using pointer in go, that created only one copy for argumants and return types
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now(),
	}
}

func NewCreateUserValidation(firstName, lastName, birthdate string) (*User, error) { //Constructor used for validation
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstName, lastName and birthDate is required!")
	}
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now(),
	}, nil

}
