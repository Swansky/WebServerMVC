package models

type User struct {
	username string
	password [32]byte
}

func NewUser(username string, password [32]byte) *User {
	user := new(User)
	user.username = username
	user.password = password
	return user
}

func (u User) GetUserName() string {
	return u.username
}

func (u User) GetPassword() [32]byte {
	return u.password
}
