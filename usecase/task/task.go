package task

import (
	"errors"
	_entities "project2/todo-list-app/entities"
	_taskRepository "project2/todo-list-app/repository/task"
)

type TaskUseCase struct {
	taskRepository _taskRepository.TaskRepositoryInterface
}

func NewTaskUseCase(taskRepo _taskRepository.TaskRepositoryInterface) TaskUseCaseInterface {
	return &TaskUseCase{
		taskRepository: taskRepo,
	}
}

func (tuc *TaskUseCase) CreatTask(newTask _entities.Task) (_entities.Task, error) {
	newTask, err := tuc.taskRepository.CreatTask(newTask)
	return newTask, err
}

func (tuc *TaskUseCase) GetAllTask(userID uint) ([]_entities.Task, error) {
	tasks, err := tuc.taskRepository.GetAllTask(userID)
	return tasks, err
}

func (tuc *TaskUseCase) GetTaskById(id uint, idToken uint) (_entities.Task, int, error) {
	task, rows, err := tuc.taskRepository.GetTaskById(id, idToken)
	return task, rows, err
}

func (tuc *TaskUseCase) UpdateTask(updateTask _entities.Task, id uint, idToken uint) (_entities.Task, int, error) {
	taskFind, rows, err := tuc.taskRepository.GetTaskById(id, idToken)
	if err != nil {
		return taskFind, 0, err
	}
	if rows == 0 {
		return taskFind, 0, nil
	}
	if updateTask.TaskName != "" {
		taskFind.TaskName = updateTask.TaskName
	}
	if updateTask.TaskStatus != "" {
		taskFind.TaskStatus = updateTask.TaskStatus
	}
	task, rows, err := tuc.taskRepository.UpdateTask(taskFind)
	return task, rows, err
}

func (tuc *TaskUseCase) DeleteTask(id uint, idToken uint) (int, error) {
	taskFind, rows, err := tuc.taskRepository.GetTaskById(id, idToken)
	if err != nil {
		return 0, err
	}
	if rows == 0 {
		return 0, nil
	}
	if taskFind.UserID != idToken {
		return 1, errors.New("unauthorized")
	}

	rowsDelete, errDelete := tuc.taskRepository.DeleteTask(id)
	return rowsDelete, errDelete

}
