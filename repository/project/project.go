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
