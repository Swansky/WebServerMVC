package repositories

import (
	"awesomeProject1/models"
	"container/list"
	"crypto/sha256"
)

type UserRepository struct {
}

func GetUserRepository() UserRepository {
	return UserRepository{}
}

func (u UserRepository) Read() (*list.List, error) {
	users := list.New()
	users.PushBack(models.NewUser("swansky", sha256.Sum256([]byte("test"))))
	return users, nil
}

func (u UserRepository) Create(entity interface{}) (interface{}, error) {
	panic("implement me")
}

func (u UserRepository) Delete(entity interface{}) error {
	panic("implement me")
}

func (u UserRepository) Update(entity interface{}) error {
	panic("implement me")
}
