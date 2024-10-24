package project

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ggicci/httpin"
	"github.com/teooliver/kanban/internal/repository/project"
	"github.com/teooliver/kanban/pkg/postgresutils"
)

type projecService interface {
	ListAllProjects(ctx context.Context, params *postgresutils.PageRequest) (postgresutils.Page[project.Project], error)
}

type Handler struct {
	service projecService
}

func New(service projecService) Handler {
	return Handler{
		service: service,
	}
}

type ListProjectResponse struct {
	Projects postgresutils.Page[project.Project] `json:"projects"`
}

func (h Handler) ListProjects(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	input := r.Context().Value(httpin.Input).(*postgresutils.PageRequest)

	projects, err := h.service.ListAllProjects(ctx, input)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("TASK HANDLER => Something went wrong: %v\n", err)))
		return
	}
	projectResponse := ListProjectResponse{
		Projects: projects,
	}

	jsonProjects, err := json.Marshal(projectResponse.Projects)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("TASK HANDLER MARSHAL => Something went wrong: %v\n", err)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonProjects))
}
