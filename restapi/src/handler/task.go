package handler

import (
	"encoding/json"
	"net/http"
	"path"
	"restapi/src/common/dto"
	"restapi/src/domain/service"

	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

type ITaskHandler interface {
	GetTasks(http.ResponseWriter, *http.Request, httprouter.Params)
	CreateTask(http.ResponseWriter, *http.Request, httprouter.Params)
	UpdateTask(http.ResponseWriter, *http.Request, httprouter.Params)
	DeleteTask(http.ResponseWriter, *http.Request, httprouter.Params)
}

type TaskHandler struct {
	service.ITaskService
}

func NewTaskHandler(srv service.ITaskService) ITaskHandler {
	return &TaskHandler{srv}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tasks, err := h.ITaskService.GetTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	var task dto.Task
	if err := json.Unmarshal(body, &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tasks, err := h.ITaskService.CreateTask(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)

	var task dto.Task
	if err := json.Unmarshal(body, &task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tasks, err := h.ITaskService.UpdateTask(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	taskId := path.Base(r.URL.Path)

	if err := h.ITaskService.DeleteTask(taskId); err == gorm.ErrRecordNotFound {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
