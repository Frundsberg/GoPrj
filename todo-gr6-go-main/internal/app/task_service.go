package app

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database" // Імпортуйте відповідний пакет
)

type TaskService interface {
	Save(t domain.Task) (domain.Task, error)
	FindById(id uint64) (domain.Task, error)
	Update(t domain.Task) (domain.Task, error)
	Delete(id uint64) error
	FindByUserId(userId uint64) ([]domain.Task, error)
}

type taskService struct {
	repo database.TaskRepository // Використовуйте правильний шлях до репозиторію
}

func NewTaskService(repo database.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) Save(t domain.Task) (domain.Task, error) {
	return s.repo.Save(t)
}

func (s *taskService) FindById(id uint64) (domain.Task, error) {
	return s.repo.FindById(id)
}

func (s *taskService) Update(t domain.Task) (domain.Task, error) {
	return s.repo.Update(t)
}

func (s *taskService) Delete(id uint64) error {
	return s.repo.Delete(id)
}

func (s *taskService) FindByUserId(userId uint64) ([]domain.Task, error) {
	return s.repo.FindByUserId(userId)
}
