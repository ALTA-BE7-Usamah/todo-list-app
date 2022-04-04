package task

import (
	_entities "project2/todo-list-app/entities"
)

type TaskRepositoryInterface interface {
	CreatTask(newTask _entities.Task) (_entities.Task, error)
	GetAllTask(userID uint) ([]_entities.Task, error)
	GetTaskById(id uint, idToken uint) (_entities.Task, int, error)
	UpdateTask(updateTask _entities.Task) (_entities.Task, int, error)
	DeleteTask(id uint) (int, error)
	CompletedTask(task _entities.Task) (_entities.Task, int, error)
}
