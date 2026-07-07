package handler

import (
	"net/http"

	"github.com/your-org/kea-ui/internal/web"
)

func Index() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		err := web.Render(w, "index.html", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}

}
