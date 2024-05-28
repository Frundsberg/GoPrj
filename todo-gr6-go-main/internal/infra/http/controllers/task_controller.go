package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
	"github.com/go-chi/chi/v5"
)

type TaskController struct {
	taskService app.TaskService
}

func NewTaskController(ts app.TaskService) TaskController {
	return TaskController{
		taskService: ts,
	}
}

func (c TaskController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)

		task, err := requests.Bind(r, requests.TaskRequest{}, domain.Task{})
		if err != nil {
			log.Printf("TaskController: %s", err)
			BadRequest(w, err)
			return
		}

		task.UserId = user.Id
		task.Status = domain.New
		task, err = c.taskService.Save(task)
		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}

		var tDto resources.TaskDto
		tDto = tDto.DomainToDto(task)
		Created(w, tDto)
	}
}

func (c TaskController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)

		taskID, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			log.Printf("TaskController: %s", err)
			BadRequest(w, err)
			return
		}

		task, err := requests.Bind(r, requests.TaskRequest{}, domain.Task{})
		if err != nil {
			log.Printf("TaskController: %s", err)
			BadRequest(w, err)
			return
		}

		task.UserId = user.Id
		task.Id = taskID

		task, err = c.taskService.Update(task)
		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}

		var tDto resources.TaskDto
		tDto = tDto.DomainToDto(task)
		Success(w, tDto)
	}
}

func (c TaskController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		taskID, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			log.Printf("TaskController: %s", err)
			BadRequest(w, err)
			return
		}

		err = c.taskService.Delete(taskID)
		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}

		noContent(w)
	}
}

func (c TaskController) FindById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		taskID, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			log.Printf("TaskController: %s", err)
			BadRequest(w, err)
			return
		}

		task, err := c.taskService.FindById(taskID)
		if err != nil {
			log.Printf("TaskController: %s", err)
			NotFound(w, err)
			return
		}

		var tDto resources.TaskDto
		tDto = tDto.DomainToDto(task)
		Success(w, tDto)
	}
}

func (c TaskController) FindByUserId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)

		tasks, err := c.taskService.FindByUserId(user.Id)
		if err != nil {
			log.Printf("TaskController: %s", err)
			InternalServerError(w, err)
			return
		}

		var tasksDto []resources.TaskDto
		for _, task := range tasks {
			var tDto resources.TaskDto
			tDto = tDto.DomainToDto(task)
			tasksDto = append(tasksDto, tDto)
		}
		Success(w, tasksDto)
	}
}
