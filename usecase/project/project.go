package project

import (
	_entities "project2/todo-list-app/entities"
	_projectRepository "project2/todo-list-app/repository/project"
)

type ProjectUseCase struct {
	projectRepository _projectRepository.ProjectRepositoryInterface
}

func NewProjectUseCase(projectRepo _projectRepository.ProjectRepositoryInterface) ProjectUseCaseInterface {
	return &ProjectUseCase{
		projectRepository: projectRepo,
	}
}

func (puc *ProjectUseCase) CreatProject(newproject _entities.Project) (_entities.Project, error) {
	newproject, err := puc.projectRepository.CreatProject(newproject)
	return newproject, err
}
