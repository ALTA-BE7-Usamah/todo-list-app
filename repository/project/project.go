package project

import (
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
