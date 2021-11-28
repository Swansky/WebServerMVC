package repositories

import "container/list"

type CRUD interface {
	Read() (*list.List, error)
	Create(entity interface{}) (interface{}, error)
	Delete(entity interface{}) error
	Update(entity interface{}) error
}
