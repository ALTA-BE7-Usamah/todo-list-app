package project

import (
	_entities "project2/todo-list-app/entities"
)

type ProjectUseCaseInterface interface {
	CreatProject(newproject _entities.Project) (_entities.Project, error)
}
