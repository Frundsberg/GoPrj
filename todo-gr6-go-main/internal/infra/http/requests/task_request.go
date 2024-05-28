package requests

import (
	"time"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type TaskRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Deadline    int64  `json:"deadline" validate:"required"`
}

func (r TaskRequest) ToDomainModel() (interface{}, error) {
	return domain.Task{
		Title:       r.Title,
		Description: r.Description,
		Deadline:    time.Unix(r.Deadline, 0),
	}, nil
}

type UpdateTaskRequest struct {
	Id          uint64 `json:"id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Deadline    int64  `json:"deadline" validate:"required"`
	Status      string `json:"status" validate:"required"`
}

func (r UpdateTaskRequest) ToDomainModel() (interface{}, error) {
	return domain.Task{
		Id:          r.Id,
		Title:       r.Title,
		Description: r.Description,
		Deadline:    time.Unix(r.Deadline, 0),
		Status:      domain.TaskStatus(r.Status),
	}, nil
}
