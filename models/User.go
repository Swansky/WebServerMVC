package models

import (
	"container/list"
)

type User struct {
	Id       string
	Username string
	Password string
}

type UserRepository interface {
	Read() (*list.List, error)
	Create(user *User) (*User, error)
	Update(user *User) error
	Delete(user *User) error
}

func NewUser(username string, password string) *User {
	user := new(User)
	user.Username = username
	user.Password = password
	return user
}

func (u User) String() string {
	return u.Id + ", " + u.Username + " " + u.Password
}
