package project

import (
	_entities "project2/todo-list-app/entities"
)

type ProjectRepositoryInterface interface {
	CreatProject(newproject _entities.Project) (_entities.Project, error)
	GetAllProject(userID uint) ([]_entities.Project, error)
}
