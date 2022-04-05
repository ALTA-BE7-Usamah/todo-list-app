package project

import (
	"errors"
	_entities "project2/todo-list-app/entities"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	database *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		database: db,
	}
}

func (pr *ProjectRepository) CreatProject(newproject _entities.Project) (_entities.Project, error) {
	tx := pr.database.Save(&newproject)
	if tx.Error != nil {
		return newproject, tx.Error
	}
	return newproject, nil
}

func (pr *ProjectRepository) GetAllProject(userID uint) ([]_entities.Project, error) {
	var projects []_entities.Project
	tx := pr.database.Where("user_id = ?", userID).Find(&projects)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return projects, nil
}

func (pr *ProjectRepository) GetProjectbyId(id uint, idToken uint) (_entities.Project, int, error) {
	var project _entities.Project
	tx := pr.database.Preload("Task").Find(&project, id)
	if tx.Error != nil {
		return project, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return project, 0, nil
	}
	if project.UserID != idToken {
		return project, 0, errors.New("id not recognise")
	}
	return project, int(tx.RowsAffected), nil
}

func (pr *ProjectRepository) AddTaskProject(addTask _entities.Task, project _entities.Project) (_entities.Project, int, error) {
	tx := pr.database.Save(&addTask)
	if tx.Error != nil {
		return project, 0, tx.Error
	}
	return project, int(tx.RowsAffected), nil
}

func (tr *ProjectRepository) DeleteProject(id uint) (int, error) {
	var project _entities.Project
	tx := tr.database.Delete(&project, id)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, tx.Error
	}
	return int(tx.RowsAffected), nil
}
