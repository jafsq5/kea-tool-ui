package handler

import (
	"net/http"

	"github.com/jafsq5/kea-tool-ui/internal/service"
	"github.com/jafsq5/kea-tool-ui/internal/web"
)

type IndexPage struct {
	Reservations any
}

func Index() http.HandlerFunc {

	svc := service.NewReservationService()

	return func(w http.ResponseWriter, r *http.Request) {

		page := IndexPage{
			Reservations: svc.List(),
		}

		if err := web.Render(w, "index.html", page); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
