package task

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/teooliver/kanban/internal/repository/task"
)

type taskService interface {
	ListAllTasks(ctx context.Context) ([]task.Task, error)
}

type Handler struct {
	service taskService
}

func New(service taskService) Handler {
	return Handler{
		service: service,
	}
}

type ListTaskResponse struct {
	Tasks []task.Task `json:"tasks"`
}

func (h Handler) ListTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := h.service.ListAllTasks(ctx)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Something went wrong")))
	}
	taskResponse := ListTaskResponse{
		Tasks: tasks,
	}

	jsonTasks, err := json.Marshal(taskResponse)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write([]byte(jsonTasks))
}