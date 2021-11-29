package repositories

import (
	"awesomeProject1/bdd"
	"awesomeProject1/models"
	"container/list"
	"crypto/md5"
	"encoding/hex"
	uuid2 "github.com/nu7hatch/gouuid"
)

type UserRepository struct {
	databaseManager *bdd.DatabaseManager
}

var instance *UserRepository

func NewUserRepository(databaseManager *bdd.DatabaseManager) *UserRepository {
	userRepository := new(UserRepository)
	userRepository.databaseManager = databaseManager
	instance = userRepository
	return userRepository
}

func (u *UserRepository) Read() (*list.List, error) {
	users := list.New()
	connection := u.databaseManager.CreateConnection()
	defer connection.Close()
	query, err := connection.Query("SELECT id,username,password FROM author")
	defer query.Close()
	if err != nil {
		return nil, err
	}
	for query.Next() {
		var user models.User
		err := query.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users.PushBack(&user)
	}
	return users, nil
}

//Create models.User enter with password not hash and without id. Return User with password hash and id
func (u *UserRepository) Create(user *models.User) (*models.User, error) {
	connection := u.databaseManager.CreateConnection()
	defer connection.Close()
	v4, err := uuid2.NewV4()

	if err != nil {
		return nil, err
	}
	hash := md5.Sum([]byte(user.Password))

	user.Password = hex.EncodeToString(hash[:])
	exec, err := connection.Exec("INSERT INTO author(id,username,password) VALUES(?,?,?)", v4.String(), user.Username, user.Password)
	if err != nil {
		return nil, err
	}
	println(exec)

	user.Id = v4.String()
	return user, nil
}

func (u *UserRepository) Update(user *models.User) error {
	panic("implement me")
}

func (u *UserRepository) Delete(user *models.User) error {
	connection := u.databaseManager.CreateConnection()
	defer connection.Close()

	query, err := connection.Query("DELETE FROM author WHERE id=?", user.Id)
	defer query.Close()
	if err != nil {
		return err
	}
	return nil
}

func GetUserRepository() *UserRepository {
	return instance
}
