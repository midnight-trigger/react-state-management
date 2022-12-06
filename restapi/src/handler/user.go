package handler

import (
	"encoding/json"
	"net/http"
	"restapi/src/domain/service"

	"github.com/julienschmidt/httprouter"
)

type ITaskHandler interface {
	GetTasks(http.ResponseWriter, *http.Request, httprouter.Params)
	FindByID(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	// Update(c echo.Context) error
	// Delete(c echo.Context) error
}

type TaskHandler struct {
	service.ITaskService
}

func NewTaskHandler(srv service.ITaskService) ITaskHandler {
	return &TaskHandler{srv}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	tasks, err := h.ITaskService.GetTasks()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if err := json.NewEncoder(w).Encode(tasks); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (h *TaskHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	// tasks, err := h.ITaskService.GetTasks()
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }

	// if err := json.NewEncoder(w).Encode(tasks); err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	// tasks, err := h.ITaskService.GetTasks()
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }

	// if err := json.NewEncoder(w).Encode(tasks); err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }
}
