package task

import (
	"errors"
	_entities "project2/todo-list-app/entities"

	"gorm.io/gorm"
)

type TaskRepository struct {
	database *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		database: db,
	}
}

func (tr *TaskRepository) CreatTask(newTask _entities.Task) (_entities.Task, error) {
	tx := tr.database.Save(&newTask)
	if tx.Error != nil {
		return newTask, tx.Error
	}
	return newTask, nil
}

func (tr *TaskRepository) GetAllTask(userID uint) ([]_entities.Task, error) {
	var tasks []_entities.Task
	tx := tr.database.Where("user_id = ?", userID).Find(&tasks)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tasks, nil
}

func (tr *TaskRepository) GetTaskById(id uint, idToken uint) (_entities.Task, int, error) {
	var task _entities.Task
	tx := tr.database.Find(&task, id)
	if tx.Error != nil {
		return task, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return task, 0, nil
	}
	if task.UserID != idToken {
		return task, 0, errors.New("id not recognise")
	}
	return task, int(tx.RowsAffected), nil
}

func (tr *TaskRepository) UpdateTask(updateTask _entities.Task) (_entities.Task, int, error) {
	tx := tr.database.Save(&updateTask)
	if tx.Error != nil {
		return updateTask, 0, tx.Error
	}
	return updateTask, int(tx.RowsAffected), nil
}

func (tr *TaskRepository) DeleteTask(id uint) (int, error) {
	var task _entities.Task
	tx := tr.database.Delete(&task, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}

func (tr *TaskRepository) CompletedTask(task _entities.Task) (_entities.Task, int, error) {
	tx := tr.database.Save(&task)
	if tx.Error != nil {
		return task, 0, tx.Error
	}
	return task, int(tx.RowsAffected), nil
}
