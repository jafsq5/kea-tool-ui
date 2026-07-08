package handler

import (
	"net/http"

	"github.com/jafsq5/kea-tool-ui/internal/hosts"
	"github.com/jafsq5/kea-tool-ui/internal/web"
)

type Handler struct {
	service *hosts.Service
}

func New(service *hosts.Service) *Handler {
	return &Handler{
		service: service,
	}
}

type IndexPage struct {
	Hosts []hosts.Host
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {

	list, err := h.service.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	page := IndexPage{
		Hosts: list,
	}

	err = web.Render(w, "index.html", page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
