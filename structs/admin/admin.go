package admin

import (
	"example/structs/user"
)

type Admin struct {
	email    string
	password string
	user.User
}

func NewAdmin(email string, password string, user user.User) *Admin {
	return &Admin{
		email:    email,
		password: password,
		User:     user,
	}
}

func (a *Admin) Email() string {
	return a.email
}

//func (a *Admin) Password() string {
//	return a.password
//}

func (a *Admin) getUser() user.User {
	return a.User
}
