package user

import (
	_entities "project2/todo-list-app/entities"
)

type UserRepositoryInterface interface {
	CreatUser(user _entities.User) (_entities.User, error)
	GetUser(id int) (_entities.User, int, error)
	UpdateUser(userUpdate _entities.User) (_entities.User, int, error)
	DeleteUser(id int) (int, error)
}
