package repository

import (
	"github.com/gilbert-keter/stellara/internal/model"
	"gorm.io/gorm"
)

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) CreateTask(task *model.Task) (model.Task, error) {
	if err := r.DB.Create(task).Error; err != nil {
		return model.Task{}, err
	}
	return *task, nil
}

func (r *TaskRepository) GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	if err := r.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
func (r *TaskRepository) GetTaskByID(id int64) (model.Task, error) {
	var task model.Task
	if err := r.DB.First(&task, id).Error; err != nil {
		return model.Task{}, err
	}
	return task, nil
}

func (r *TaskRepository) UpdateTask(task *model.Task) (model.Task, error) {
	if err := r.DB.Save(task).Error; err != nil {
		return model.Task{}, err
	}
	return *task, nil
}
func (r *TaskRepository) DeleteTask(id int64) error {
	if err := r.DB.Delete(&model.Task{}, id).Error; err != nil {
		return err
	}
	return nil
}
