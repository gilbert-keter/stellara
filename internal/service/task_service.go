package service

import (
	"github.com/gilbert-keter/stellara/internal/model"
	"github.com/gilbert-keter/stellara/internal/repository"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{Repo: repo}
}

func (s *TaskService) CreateTask(task *model.Task) (model.Task, error) {
	return s.Repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]model.Task, error) {
	return s.Repo.GetAllTasks()
}

func (s *TaskService) GetTaskByID(id int64) (model.Task, error) {
	return s.Repo.GetTaskByID(id)
}

func (s *TaskService) UpdateTask(task *model.Task) (model.Task, error) {
	return s.Repo.UpdateTask(task)
}
func (s *TaskService) DeleteTask(id int64) error {
	return s.Repo.DeleteTask(id)
}
