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

func (puc *ProjectUseCase) GetAllProject(userID uint) ([]_entities.Project, error) {
	projects, err := puc.projectRepository.GetAllProject(userID)
	return projects, err
}

func (puc *ProjectUseCase) GetProjectbyId(id uint, idToken uint) (_entities.Project, int, error) {
	project, rows, err := puc.projectRepository.GetProjectbyId(id, idToken)
	return project, rows, err
}
